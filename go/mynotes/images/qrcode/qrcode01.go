package main

import (
	cre "github.com/skip2/go-qrcode"

	"fmt"
	"os"
	read "github.com/tuotoo/qrcode"
)

//二维码相关
func main() {
	//createFile()
	readFile()
}

func createFile()  {
	err := cre.WriteFile("http://www.wangdy.com/index.html", cre.Medium, 256, "./images/qrcode/qr.png")
	if err != nil {
		fmt.Println("write error")
	}
}

func readFile()  {
	fi, err := os.Open("./images/qrcode/1.jpg")
	if err != nil{
		fmt.Println(err.Error())
		return
	}
	defer fi.Close()
	qrmatrix, err := read.Decode(fi)
	if err != nil{
		fmt.Println(err.Error())
		return
	}
	fmt.Println(qrmatrix.Content)
}