package method

import (
	"net/http"
	"github.com/PuerkitoBio/goquery"
	"log"
	"fmt"
	"strconv"
	"os"
	"net/url"
	"strings"
	"github.com/Tang-RoseChild/mahonia"
	"sync"
	"io/ioutil"
)

//返回response
func getResponse(str string) *http.Response {
	client := &http.Client{}
	request, _ := http.NewRequest("GET", str, nil)
	//request.Header.Add("Content-Type", "text/html; charset=gb2312")
	response, _ := client.Do(request)
	return response
}

var Wg sync.WaitGroup

//爬fang链接
func GetFangLink(path string) []string {
	var info []string
	urlParse,_ := url.Parse(path)
	response := getResponse(urlParse.String())
	resCode := response.StatusCode
	log.Println("code:",resCode)
	if response.StatusCode == 200 {
		dec := mahonia.NewDecoder("GB18030")
		doc, err := goquery.NewDocumentFromReader(dec.NewReader(response.Body))
		if err != nil{
			log.Fatal(err)
		}
		fmt.Println(path,",Code：",resCode)
		doc.Find(".title a").Each(func(i int, s *goquery.Selection) {
			link,_ := s.Attr("href")
			info = append(info,link)
		})
	}

	return info
}

type UserInfo struct {
	Name string
	Firm string
	Tel string
}
var Users []*UserInfo
//爬fang人信息
func GetFangUser(path string) {
	Wg.Add(1)
	defer Wg.Done()
	urlParse,_ := url.Parse(path)
	response := getResponse(urlParse.String())
	resCode := response.StatusCode
	if response.StatusCode == 200 {
		dec := mahonia.NewDecoder("utf-8")
		doc, err := goquery.NewDocumentFromReader(dec.NewReader(response.Body))
		if err != nil{
			log.Fatal(err)
		}
		fmt.Println(path,",Code：",resCode)
		doc.Find(".tjcont-list-c").Each(func(i int, s *goquery.Selection) {
			if i > 2 {
				return
			}
			name := s.Find(".zf_jjname").Find(".font16").Text()
			firm := s.Find(".clearfix").Find(".gray9").Text()
			tel := s.Find(".font22").Find(".red").Text()
			Users = append(Users, &UserInfo{
				Name:strings.Replace(name,"\n","",-1),
				Firm:strings.Replace(firm,"\n","",-1),
				Tel:strings.Replace(tel,"\n","",-1),
			})
		})
	}
}

var aa = "https://baike.baidu.com/item/"
//爬百科
func GetBaikeDog(str string) string {
	var info string
	urlParse,_ := url.Parse(aa+str)
	response := getResponse(urlParse.String())
	resCode := response.StatusCode
	if response.StatusCode == 200 {
		doc, err := goquery.NewDocumentFromReader(response.Body)
		if err != nil{
			log.Fatal(err)
		}
		fmt.Println(str,",Code：",resCode)
		doc.Find(".basicInfo-block").Each(func(i int, s *goquery.Selection){
			info += strings.Replace(s.Text(),"\n","",-1)
		})
	}
	return info
}

//爬糗事
func getJokes(){
	urlParse := "httpReq://www.qiushibaike.com"//"https://www.nvshens.com/girl/18071/album/"
	response := getResponse(urlParse)
	resCode := response.StatusCode
	if response.StatusCode == 200 {
		doc, err := goquery.NewDocumentFromReader(response.Body)
		if err != nil{
			log.Fatal(err)
		}
		fmt.Println(doc.Find(".article").Text())
		//doc.Find(".context").Each(func(i int, s *goquery.Selection){
		//	text := s.Text()
		//	fmt.Println(i,"----",text)
		//})
	}
	fmt.Println("爬取状态：",resCode)
	os.Exit(1)
}

//爬go中文
func GetGoChinaDoc() {
	var page int = 1
	baseUrl := "https://studygolang.com/topics?p="
	count :=getPageCount(baseUrl)
	fmt.Println(count)
	for {
		str := baseUrl + strconv.Itoa(page)
		response := getResponse(str)
		if response.StatusCode == 403 {
			fmt.Println("IP 已被禁止访问")
			os.Exit(1)
		}
		if response.StatusCode == 200 {
			doc, err := goquery.NewDocumentFromReader(response.Body)
			if err != nil {
				log.Fatalf("失败原因", response.StatusCode)
			}
			doc.Find(".topics .topic").Each(func(i int, content *goquery.Selection) {
				link,_ := content.Find(".title a").Attr("href")
				fmt.Println(link)
			})
		}
		if page >= 1{
			fmt.Println("-----------数据拉取完成共"+strconv.Itoa(page)+"页-------------")
			break
		}
		page++
	}
	os.Exit(1)
}
//得到文章总数
func getPageCount(baseUrl string) int {
	response := getResponse(baseUrl)
	dom, err := goquery.NewDocumentFromReader(response.Body)
	if err != nil {
		log.Fatalf("失败原因", response.StatusCode)
	}
	resDom := dom.Find(".text-center .pagination-sm li a")
	//len := resDom.Length()
	count,_ := strconv.Atoi(resDom.Eq(resDom.Length()-2).Text())
	return count
}

func ReqForm(address string, params url.Values) []byte {
	res, _ := http.PostForm(address, params)
	if res.StatusCode == 200 {
		body, _ := ioutil.ReadAll(res.Body)
		return body
	} else {
		body, _ := ioutil.ReadAll(res.Body)
		fmt.Println(string(body))
		return nil
	}
}