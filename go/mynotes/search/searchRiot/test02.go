package main

import (
	"log"

	"github.com/go-ego/riot"
	"github.com/go-ego/riot/types"
)

var (
	searcher02 = riot.New("zh")
)

func main() {
	data := types.DocData{Content: `I wonder how, I wonder why, I wonder where they are`}
	data1 := types.DocData{Content: "所以, 你, 再见"}
	data2 := types.DocData{Content: "没有理由"}
	searcher02.Index(1, data)
	searcher02.Index(2, data1)
	searcher02.IndexDoc(3, data2)
	searcher02.Flush()

	req := types.SearchReq{Text: "你好"}
	search := searcher02.Search(req)
	log.Println("search...", search)
}