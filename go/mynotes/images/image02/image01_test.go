package image02

import (
	"testing"
	"encoding/base64"
	"io/ioutil"
	"fmt"
	"bytes"
	"image"
	"os"
	"image/jpeg"
	"net/url"
)

//原生包读取与转换字符串
func Test_ImgFileBase64(t *testing.T)  {
	//file1, err := os.Open("./imagepkg/01.jpg")
	//fmt.Println(err)
	bs, _ := ioutil.ReadFile("./imagepkg/a.txt")
	fmt.Println("aaa",url.QueryEscape(base64.StdEncoding.EncodeToString(bs)),"aaa")
	//fmt.Println(base64.URLEncoding.EncodeToString(bs))
}

func Test_FileToBase64ToFile(t *testing.T)  {
	//fileBytes, _ := ioutil.ReadFile("./imagepkg/01.jpg")
	//byteLen := fileBytes
	//lens := len(byteLen)/2*3
	//fmt.Println("byteLen:", len(byteLen))
	//dataSource := make([]byte, lens)         //数据缓存
	//base64.StdEncoding.Encode(dataSource,fileBytes)   // 文件转base64
	//fmt.Println(string(dataSource))
	////ddd, _ := base64.StdEncoding.DecodeString(string(dataSource)) //成图片文件并把文件写入到buffer
	////err := ioutil.WriteFile("./output.jpg", ddd, 0666)   //buffer输出到jpg文件中
	////fmt.Println("err:",err)
	//bbb := bytes.NewBuffer(dataSource)
	//fmt.Println(bbb.String())

	//todo Buffer -> ImageBuff
	ff,err := os.Open("./imagepkg/01.jpg")
	fmt.Println(err)
	m, fn, _ := image.Decode(ff)                                       // 图片文件解码
	fmt.Println("fn:",fn)
	rgbImg := m.(*image.YCbCr)
	subImg := rgbImg.SubImage(image.Rect(0, 0, 200, 200)).(*image.YCbCr) //图片裁剪x0 y0 x1 y1

	//todo Img -> File
	f, _ := os.Create("./output.jpg")     //创建文件
	defer f.Close()                   //关闭文件
	jpeg.Encode(f, subImg, nil)       //写入文件
}

func Test_ImgFile(t *testing.T)  {
	//todo Base64 -> buffer
	dataSource := ""
	ddd, _ := base64.StdEncoding.DecodeString(dataSource) //成图片文件并把文件写入到buffer
	bbb := bytes.NewBuffer(ddd)
	fmt.Println(bbb)

	//todo Buffer -> ImageBuff
	m, _, _ := image.Decode(bbb)                                       // 图片文件解码
	rgbImg := m.(*image.YCbCr)
	subImg := rgbImg.SubImage(image.Rect(0, 0, 200, 200)).(*image.YCbCr) //图片裁剪x0 y0 x1 y1

	//todo Img -> File
	f, _ := os.Create("test.jpg")     //创建文件
	defer f.Close()                   //关闭文件
	jpeg.Encode(f, subImg, nil)       //写入文件

	//todo ImgBase64
	emptyBuff := bytes.NewBuffer(nil)                  //开辟一个新的空buff
	jpeg.Encode(emptyBuff, subImg, nil)                //img写入到buff
	dist := make([]byte, 50000)                        //开辟存储空间
	base64.StdEncoding.Encode(dist, emptyBuff.Bytes()) //buff转成base64
	fmt.Println(string(dist))                          //输出图片base64(type = []byte)
	_ = ioutil.WriteFile("./base64pic.txt", dist, 0666) //buffer输出到jpg文件中

	//todo ImgFileBase64
	ff, _ := ioutil.ReadFile("output2.jpg")
	bufstore := make([]byte, 5000000)                     //数据缓存
	base64.StdEncoding.Encode(bufstore, ff)               // 文件转base64
	_ = ioutil.WriteFile("./output2.jpg.txt", dist, 0666) //写入到文件
}


