package main

import (
	"mynotes/httpReq/crawler01/pipeline"
	"mynotes/httpReq/crawler01/spider"
	"log"
	"sync"
)

// 根据返回的json数据来处理
func main() {
	for _, kd := range kds {
		for _, city := range citys {
			wg.Add(1)
			go func(city string, kd string) {
				defer wg.Done()
				initResult, err := spider.InitJobs(city, 1, kd)
				if err != nil {
					log.Fatalln(err)
				}

				initResults = append(initResults, initResult...)
				loopResults = append(loopResults, spider.LoopJobs())
			}(city, kd)
		}
	}

	wg.Wait()

	jobPipeline.Push()

	log.Printf("Init Results: %v", initResults)
	log.Printf("Loop Results: %v", loopResults)
}

var (
	kds = []string{
		"golang",
	}
	citys = []string{
		"北京",
		"石家庄",
		"天津",
		//"上海",
		//"广州",
		//"深圳",
		//"杭州",
		//"成都",
	}

	initResults = []spider.InitResult{}
	loopResults = []spider.LoopResult{}
	jobPipeline = pipeline.NewJobPipeline()

	wg sync.WaitGroup
)
