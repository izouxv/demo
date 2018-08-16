package main

import (
	"fmt"
	"sync"
	"runtime"
	"strings"
)

func main() {
	//i3Test01()
	//i3Test02()
	//i3Test03()
	//i3Test04()
	//i3Test05()
	//i3Test06()
	//i3Test07()
	i3Test08()
}

func i3Test08() {
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			print(i, ",", j, " ")
			break
		}
		println()
	}
//outer:
//	for i := 0; i < 3; i++ {
//		for j := 0; j < 3; j++ {
//			print(i,",",j)
//			break outer
//		}
//		println()
//	}
}

//指针、实体对nil的应用
func i3Test07() {
	//fmt.Println(InitType() == nil)
}
func InitType() t1 {
	var t t1
	return t
}
type t1 struct {}

func i3Test06() {
	c := &ConfigOne{Daemon:"aaa"}
	fmt.Println(c.String())
}
type ConfigOne struct {
	Daemon string
}
func (c *ConfigOne) String() string {
	return fmt.Sprintf("aprint: %v", c.Daemon)
}

//判断是否包含字符串并返回下标（中文长度为3）
func i3Test05() {
	fmt.Println(Utf8Index("111北京天安门最美丽", "天安门"))
	fmt.Println(strings.Index("111北京天安门最美丽", "天安门"))
	//fmt.Println(Utf8Index("12ws北京天安门最美丽", "天安门"))
}
func Utf8Index(str, substr string) int {
	asciiPos := strings.Index(str, substr)
	if asciiPos == -1 || asciiPos == 0 {
		return asciiPos
	}
	pos := 0
	totalSize := 0
	reader := strings.NewReader(str)
	for _, size, err := reader.ReadRune(); err == nil; _, size, err = reader.ReadRune() {
		fmt.Println("size:",size)
		totalSize += size
		pos++
		// 匹配到
		if totalSize == asciiPos {
			return pos
		}
	}
	return pos
}

//接口类型与实体类型的应用
func i3Test04() {
	var peo People1 = &Stduent{}
	think := "bitch"
	fmt.Println(peo.Speak(think))
}
type People1 interface {
	Speak(think string) string
}
type Stduent struct{}
func (stu *Stduent) Speak(think string) (talk string) {
	if think == "bitch" {
		talk = "You are a good boy"
	} else {
		talk = "hi"
	}
	return
}

//题意理解与字符串解析
//有一个机器人，给一串指令，L左转 R右转，F前一步 B后一步，问最后机器人的坐标，最开始，机器人位于 0 0，方向为正Y。
//可以输入重复指令n ： 比如 R2(LF) == RLFLF，问最后机器人的坐标是多少？
func i3Test03() {
	//a001
	//a002
}

//go执行顺序随机,goroutine并发不等于并行
func i3Test02() {
	runtime.GOMAXPROCS(1)//goroutine数
	wg := sync.WaitGroup{}
	wg.Add(20)
	for i := 0; i < 10; i++ {
		go func() {
			fmt.Println("i:", i)
			defer wg.Done()
		}()
	}
	for i := 0; i < 10; i++ {
		go func(i int) {
			fmt.Println("j: ", i)
			defer wg.Done()
		}(i)
	}
	//time.Sleep(time.Second)
	wg.Wait()
}

func i3Test01() {
	s := []int{1, 2, 3}
	ss := s[1:]
	fmt.Println(ss)
	ss = append(ss, 4)
	for _, v := range ss {
		v += 10
	}
	for i := range ss {
		ss[i] += 10
	}
	fmt.Println(s)
}
