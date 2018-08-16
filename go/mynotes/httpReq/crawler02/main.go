package main

import (
	"github.com/gocolly/colly"
	"fmt"
	"github.com/360EntSecGroup-Skylar/excelize"
	"time"
	"sync"
	"strconv"
	"runtime"
)

var wg sync.WaitGroup

// colly框架
func main() {
	runtime.GOMAXPROCS(8)
	c := colly.NewCollector()
	go c.OnHTML(".title a", func(e *colly.HTMLElement) {
		wg.Add(1)
		e.Request.Visit(e.Attr("href"))
		wg.Done()
	})
	xlsx1 := excelize.NewFile()
	err := xlsx1.SaveAs("./links.xlsx")
	if err != nil {
		fmt.Println(err)
	}
	xlsx2, err := excelize.OpenFile("./links.xlsx")
	if err != nil {
		fmt.Println(err)
		time.Sleep(time.Second*5)
	}
	//var links []string
	a := 1
	go c.OnRequest(func(req *colly.Request) {
		wg.Add(1)
		fmt.Println("visiting:",req.URL.String())
		xlsx2.SetCellValue("Sheet1", "A"+strconv.Itoa(a), req.URL)
		err = xlsx2.Save()
		if err != nil {
			fmt.Println(err)
		}
		//links = append(links,req.URL.String())
		a++
		wg.Done()
	})
	c.Visit("http://esf.sjz.fang.com/integrate/")
	wg.Wait()
	//fmt.Println(links)

}
