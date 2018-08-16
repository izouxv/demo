package main

import (
	"fmt"
	"net/url"
	"net/http"
	"log"
	"encoding/json"
	"bufio"
	"os"
	"strings"
)

func main() {
	//从标准输入读取数据
	inputReader := bufio.NewReader(os.Stdin)
	fmt.Println("Please input your name：")
	//读取数据直到遇见\n位置
	input, err := inputReader.ReadString('\n')
	if err != nil {
		fmt.Printf("An error occurred:%s\n", err)
		//异常错误后退出
		os.Exit(1)
	} else {
		//用切片操作删除最后的\n
		name := input[:len(input)-1]
		fmt.Printf("Hello,%s!What can i di for you?\n", name)
	}
	for {
		input, err = inputReader.ReadString('\n')
		if err != nil {
			fmt.Printf("An error occurred:%s\n", err)
			continue
		}
		input = input[:len(input)-1]
		//全部转换为小写
		input = strings.TrimSpace(strings.ToLower(input))
		switch input {
		case "":
			continue
		case "nothing", "bye","exit","quit":
			fmt.Println("Bye!")
			//正常退出
			os.Exit(0)
		default:
			fmt.Println(tlAI(input))
		}
	}
}

//从tlAI获取回复
func tlAI(info string) *tlReply {
	tuLingURL := fmt.Sprintf("httpReq://www.tuling123.com/openapi/api?key=%s&info=%s", "70c990cb300e4b9ba83a036586df280a", url.QueryEscape(info))
	resp, err := http.Get(tuLingURL)
	reply := new(tlReply)
	if err != nil {
		log.Println(err)
		return reply
	}
	defer resp.Body.Close()
	decoder := json.NewDecoder(resp.Body) //decode reply from response body
	decoder.Decode(reply)
	return reply
}

type tlReply struct {
	code	int
	Text	string 			`json:"text"`
	List	[]interface{}	`json:"list"`
	Url		string			`json:"url"`
	Function	interface{}	`json:"function"`
}
