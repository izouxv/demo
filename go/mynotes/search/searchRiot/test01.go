package main

import (
	"github.com/go-ego/riot"
	"github.com/go-ego/riot/types"
	"fmt"
)

var (
	// searcher 是线程安全的
	searcher01 = riot.Engine{}
)



func main() {
	// 初始化
	searcher01.Init(types.EngineOpts{
		Using: 3,
		// IDOnly:        true,
		GseDict: "search/searchRiot/dictionary.txt",
		NotUseGse:true,
	})
	defer searcher01.Close()

	text := "此次度量百十个国家中成绩的是XXX"
	text1 := "中国百度宣布拟全资收购91无线业务"
	text2 := "百度是中国最大的搜索引擎"

	// 将文档加入索引，docId 从1开始
	go searcher01.Index(1, types.DocData{Content: text})
	go searcher01.Index(2, types.DocData{Content: text1},false)
	go searcher01.Index(3, types.DocData{Content: text2},true)

	// 等待索引刷新完毕
	searcher01.Flush()

	// 搜索输出格式见 types.SearchResp 结构体
	output := searcher01.Search(types.SearchReq{Text: "百 度 中 国"})
	fmt.Println(output.Timeout)
	fmt.Println(output.NumDocs)
	fmt.Println(output.Docs)
	fmt.Println(output.Tokens)
}
