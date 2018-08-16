package main

import (
	"strconv"
	"mynotes/httpReq/method"
	"github.com/360EntSecGroup-Skylar/excelize"
	"strings"
	"fmt"
)

//爬test
func main() {
	method.GetGoChinaDoc()
	//getJokes()
	//fmt.Println(method.GetBaikeDog("阿富汗猎犬"))
	//test01()
	craFangUser()
}

var LinkUrl = "http://esf.sjz.fang.com"
func craFangUser()  {
	var url string
	uri := "http://esf.sjz.fang.com/integrate"
	for i:=1;i <= 100;i++ {
		if i == 1 {
			url = uri
		} else {
			url = uri + "/i"+ strconv.Itoa(i+30) +"/"
		}
		links := method.GetFangLink(url)
		for _,v := range links {
			go method.GetFangUser(LinkUrl+v)
		}
	}
	method.Wg.Wait()
	xlsx := excelize.NewFile()
	for k,v := range method.Users {
		xlsx.SetCellValue("Sheet1", "A"+strconv.Itoa(k+1), v.Name)
		xlsx.SetCellValue("Sheet1", "B"+strconv.Itoa(k+1), v.Tel)
		var firm []rune
		for _,v := range []rune(v.Firm) {
			if v == 32 || v == 160 || v == 8194 {
				continue
			}
			firm = append(firm,v)
		}
		firms := strings.Split(string(firm),"|")
		xlsx.SetCellValue("Sheet1", "C"+strconv.Itoa(k+1), firms[0])
	}
	err := xlsx.SaveAs("user_info.xlsx")
	if err != nil {
		fmt.Println(err)
	}
}
