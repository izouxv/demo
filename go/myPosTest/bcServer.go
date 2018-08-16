package main

import (
	"bufio"
	"encoding/json"
	"io"
	"log"
	"net"
	"os"
	"strconv"
	"sync"
	"time"

	"github.com/davecgh/go-spew/spew"
	"github.com/joho/godotenv"
	"myPosTest/structs"
	"fmt"
)



// bcServer handles incoming concurrent Blocks
var bcServer chan []*structs.Block
var mutex = &sync.Mutex{}

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal(err)
	}

	bcServer = make(chan []*structs.Block)

	// create genesis block
	t := time.Now()
	genesisBlock := &structs.Block{0, t.String(), 0, "", "",0,""}
	spew.Dump(genesisBlock)
	structs.BlockChain = append(structs.BlockChain, genesisBlock)

	httpPort := os.Getenv("PORT")

	// start TCP and serve TCP server
	server, err := net.Listen("tcp", ":"+httpPort)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("HTTP Server Listening on port :", httpPort)
	defer server.Close()

	for {
		conn, err := server.Accept()
		if err != nil {
			log.Fatal(err)
		}
		go handleConn(conn)
	}

}

func handleConn(conn net.Conn) {

	defer conn.Close()

	io.WriteString(conn, "Enter a new BPM:")

	scanner := bufio.NewScanner(conn)

	// take in BPM from stdin and add it to blockchain after conducting necessary validation
	go func() {
		for scanner.Scan() {
			bpm, err := strconv.Atoi(scanner.Text())
			fmt.Println("bpm:",bpm)
			if err != nil {
				log.Printf("%v not a number: %v", scanner.Text(), err)
				continue
			}
			newBlock,_ := structs.GenerateBlock(structs.BlockChain[len(structs.BlockChain)-1], bpm)
			if structs.IsBlockValid(structs.BlockChain[len(structs.BlockChain)-1], newBlock) {
				newBlockChain := append(structs.BlockChain, newBlock)
				structs.ReplaceChain(newBlockChain)
			}

			bcServer <- structs.BlockChain
			io.WriteString(conn, "\nEnter a new BPM:")
		}
	}()

	// simulate receiving broadcast
	go func() {
		for {
			time.Sleep(30 * time.Second)
			mutex.Lock()
			output, err := json.Marshal(structs.BlockChain)
			if err != nil {
				log.Fatal(err)
			}
			mutex.Unlock()
			io.WriteString(conn, string(output))
		}
	}()

	for _ = range bcServer {
		spew.Dump("BlockChain:",structs.BlockChain)
	}

}
