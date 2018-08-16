package main

import (
	"fmt"
	"strconv"
	"strings"
	"time"
	"math"
	"sync"
	"github.com/coreos/etcd/clientv3"
	"context"
	"os/exec"
	"golang.org/x/text/encoding/simplifiedchinese"
	"bytes"
)

func main() {
	//fmt.Println(powerInt(10,63))
	//StrReplace()
	//FloatString()
	//FloatToStrLength()
	//timeToStr()
	//testMax()
	//mathSqrt()
	//controChan()
	//testEtcd()
	//execCmd01()
	execCmd02()
}

func execCmd02() {
	cmd := exec.Command("cmd")
	in := bytes.NewBuffer(nil)
	cmd.Stdin = in
	var out bytes.Buffer
	cmd.Stdout = &out
	go func() {
		in.WriteString("cd maths \n")
		in.WriteString("type tmp.txt \n")
	}()
	err := cmd.Start()
	if err != nil {
		fmt.Println("error1:",err)
	}
	fmt.Println("Args:",cmd.Args)
	err = cmd.Wait()
	if err != nil {
		fmt.Println("error2:", err)
	}
	fmt.Println(out.String())

}
func execCmd01()  {
	cmd := exec.Command("ls")
	bs,err := cmd.CombinedOutput()
	if err != nil {
		fmt.Println("error:", err)
		return
	}
	utf8Bs,err := simplifiedchinese.GBK.NewDecoder().Bytes(bs)
	fmt.Println("bs:",string(utf8Bs))

	//cmd = exec.Command("type","maths/tmp.txt")
	//stdout, err := cmd.StdoutPipe()
	//if err != nil {
	//	fmt.Println("err1:",err)
	//	return
	//}
	//err = cmd.Start()
	//if err != nil {
	//	fmt.Println("err3:",err)
	//	return
	//}
	//content, err := ioutil.ReadAll(stdout)
	//if err != nil {
	//	fmt.Println("err2:",err)
	//	return
	//}
	//fmt.Println(string(content))     //输出ls命令查看到的内容
	//cmd.Wait()
}

func testEtcd()  {
	cli,err := clientv3.New(
		clientv3.Config{
			Endpoints:[]string{"127.0.0.1:2379"},
			DialTimeout:5*time.Second,
		})
	if err != nil {
		fmt.Println("connect failed, err:", err)
		return
	}
	fmt.Println("connect succ")
	defer cli.Close()
	//设置1秒超时，访问etcd有超时控制
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	//操作etcd
	putRes, err := cli.Put(ctx, "/logagent/conf/", "888888")
	if err != nil {
		fmt.Println("put failed, err:", err)
		return
	}
	defer cancel()
	fmt.Println("putRes:",putRes)
	rch := cli.Watch(context.Background(), "/logagent/conf/")
	for wresp := range rch {
		for _, ev := range wresp.Events {
			fmt.Printf("%s %q : %q\n", ev.Type, ev.Kv.Key, ev.Kv.Value)
		}
	}

	//取值，设置超时为1秒
	//ctx, cancel = context.WithTimeout(context.Background(), time.Second)
	//resp, err := cli.Get(ctx, "/logagent/conf/")
	//cancel()
	//if err != nil {
	//	fmt.Println("get failed, err:", err)
	//	return
	//}
	//for _, ev := range resp.Kvs {
	//	fmt.Printf("%s : %s\n", ev.Key, ev.Value)
	//}
}

var (
	num = 20
	sy	sync.WaitGroup
	line  = 0
	chans []chan int
	lens int
)
func controChan() {
	chans = []chan int{
		make(chan int),
		make(chan int),
		make(chan int),
		make(chan int) }
	lens = len(chans)
	sy.Add(1)
	go ChanWork()
	chans[0] <- 1
	sy.Wait()
}
func ChanWork() {
	if count := <-chans[line]; count <= num {
		fmt.Println(strconv.Itoa(line)+" -> ",count)
		count++
		line++
		if line >= lens {
			line = 0 //循环，防止索引越界
		}
		go ChanWork()
		chans[line] <- count
	} else {
		fmt.Println("aaaaaaaa")
		sy.Done()
		return
	}
}

func mathSqrt() {
	//fmt.Println(62*62+(-380*-380)+4120*4120)
	fmt.Println(math.Sqrt(17122644))
}

func testMax() {
	fmt.Println(math.MaxUint16)
	fmt.Println(1<<8)
}

func timeToStr() {
	fmt.Println(time.Now().Format("2006-01-02-15-04-05"))
}

//保留小数点后n位
func FloatToStrLength()  {
	fmt.Println(fmt.Sprintf("%."+"2"+"f",1.111111))
}

//字符串替换
func StrReplace()  {
	fmt.Println(strings.Replace("111222333111","111","---",3))
}

//数组截取
func Strs()  {
	aa := []string{"aa","bb","cc"}
	fmt.Println(aa[:2])
}

//小数转str
func FloatString()  {
	fmt.Println(strconv.FormatFloat(float64(123456)/float64(1000),'f',-1,32))
}
//递归幂函数
func powerInt(x uint64, n uint64) uint64 {
	if n == 0 {
		return 1
	}
	fmt.Println(x)
	return x * powerInt(x, n-1)
}
