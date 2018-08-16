package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"
	"time"

	"github.com/huichen/gobo"
	"github.com/huichen/gobo/contrib"
)

var (
	access_token = flag.String("access_token", "2.00SQ2TqDFilfAE3c407b0e09azHEyB", "用户的访问令牌")
	weibo        = gobo.Weibo{}
	users_file   = flag.String("users_file", "./data/users.txt", "从该文件读入要下载的微博用户名，每个名字一行")
	output_file  = flag.String("output_file", "./data/weibo_data.txt", "将抓取的微博写入下面的文件")
	num_weibos   = flag.Int("num_weibos", 100, "从每个微博账号中抓取多少条微博")
)

func main() {
	flag.Parse()
	content, err := ioutil.ReadFile(*users_file)
	if err != nil {
		fmt.Println("无法读取users_file err:",err)
		return
	}
	users := strings.Split(string(content), "\n")
	outputFile, err := os.Create(*output_file)
	if err != nil {
		fmt.Println("Create file err:",err)
		return
	}
	defer outputFile.Close()
	for _, user := range users {
		if user == "" {
			continue
		}
		log.Printf("抓取 %s", user)
		statuses, err := contrib.GetStatuses(&weibo, *access_token, user, 0, *num_weibos, 5000) // 超时5秒
		fmt.Println("len:",len(statuses))
		if err != nil {
			fmt.Println(err)
			continue
		}
		for _, status := range statuses {
			fmt.Println("aaa:",status.Id)
			t, _ := time.Parse("Mon Jan 2 15:04:05 -0700 2006", status.Created_At)
			outputFile.WriteString(fmt.Sprintf(
				"%d||||%d||||%d||||%s||||%d||||%d||||%d||||%s||||%s||||%s\n",
				status.Id, uint32(t.Unix()), status.User.Id, status.User.Screen_Name,
				status.Reposts_Count, status.Comments_Count, status.Attitudes_Count,
				status.Thumbnail_Pic, status.Original_Pic, status.Text))
		}
	}
}
