package main

import (
	"github.com/joho/godotenv"
	"time"
	"myPosTest/structs"
	"github.com/davecgh/go-spew/spew"
	"os"
	"log"
	"net/http"
	"github.com/gorilla/mux"
	"encoding/json"
	"io"
)

/**
参考源码地址：https://github.com/mycoralhealth/blockchain-tutorial
 */

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal(err)
	}
	go func() {
		t := time.Now()
		genesisBlock := &structs.Block{0, t.String(), 0, "", "",0, ""}
		spew.Dump("main-genesisBlock:",genesisBlock)
		structs.BlockChain = append(structs.BlockChain, genesisBlock)
	}()
	log.Fatal(run())
}

//启动web服务
func run() error {
	muxRouter := makeMuxRouter()
	port := os.Getenv("ADDR")
	log.Println("Listening on port:", port)
	s := &http.Server{
		Addr:           ":" + port,
		Handler:        muxRouter,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	if err := s.ListenAndServe(); err != nil {
		return err
	}
	return nil
}

//定义不同endpoint以及对应的 handler
func makeMuxRouter() http.Handler {
	muxRouter := mux.NewRouter()
	muxRouter.HandleFunc("/", handleGetBlockchain).Methods("GET")
	muxRouter.HandleFunc("/", handleWriteBlock).Methods("POST")
	return muxRouter
}

//GET 请求的 handler
func handleGetBlockchain(w http.ResponseWriter, r *http.Request) {
	bytes, err := json.MarshalIndent(structs.BlockChain, "", "  ")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	io.WriteString(w, string(bytes))
}


//POST 请求的handler
func handleWriteBlock(w http.ResponseWriter, r *http.Request) {
	var m structs.Message
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&m); err != nil {
		respondWithJSON(w, r, http.StatusBadRequest, r.Body)
		return
	}
	defer r.Body.Close()
	newBlock,_ := structs.GenerateBlock(structs.BlockChain[len(structs.BlockChain)-1], m.BPM)
	if structs.IsBlockValid(structs.BlockChain[len(structs.BlockChain)-1],newBlock) {
		structs.ReplaceChain(append(structs.BlockChain, newBlock))
		spew.Dump(structs.BlockChain)
	}
	respondWithJSON(w, r, http.StatusCreated, newBlock)

}

//POST请求后，返回响应
func respondWithJSON(w http.ResponseWriter, r *http.Request, code int, payload interface{}) {
	response, err := json.MarshalIndent(payload, "", "  ")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("HTTP 500: Internal Server Error"))
		return
	}
	w.WriteHeader(code)
	w.Write(response)
}




