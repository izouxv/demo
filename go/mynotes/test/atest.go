package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
	"time"
	"unicode"
	"net"
	"golang.org/x/text/message"
	"golang.org/x/text/language"
	"golang.org/x/text/currency"
	"sync/atomic"
	"sync"
	"mynotes/httpReq/method"
	"io"
)

func main() {
	//slices()
	//stringMe()
	//IsChineseChar()
	//test01()
	//test02()
	//test03()
	//test04()
	//test05()
	//test06()
	//test07()
	//test08()
	//test09()
	//test10()
	//test11()
	//test12()
	//test13()
	//test14()
	locationToDecimal()
}

func locationToDecimal()  {
	file,err := os.Open("g.log")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()
	buf := bufio.NewReader(file)
	var points []*method.Point
	time1 := time.Now().Add(-time.Minute*15).Unix()
	for {
		line, err := buf.ReadString('\n')
		line = strings.TrimSpace(line)
		if err != nil {
			if err == io.EOF {
				fmt.Println("File read ok!")
				break
			} else {
				fmt.Println("Read file error!", err)
				return
			}
		}
		lines := strings.Split(line," ")
		lng,lat := method.ToLocation(lines[1],lines[0])
		points = append(points,&method.Point{Longitude:lng,Latitude:lat,LocTime:time1,EntityName:method.EntityName,CoordTypeInput:method.CoordTypeInput})
		time1 += 1
		if len(points) == 100 {
			method.ToParams(points)
			points = nil
			time.Sleep(time.Minute)
		}
	}
	method.ToParams(points)
}

func test14() {
	fmt.Println(time.Now().Unix())
	//_, filePath, _, _ := runtime.Caller(1)
	//fmt.Println("aa",path.Dir(filePath))
}

func test13()  {
	var a int32
	fmt.Println("a : ", a)
	//函数名以Add为前缀，加具体类型名,参数一，是指针类型,参数二，与参数一类型总是相同
	new_a := atomic.AddInt32(&a, 3)
	fmt.Println("new_a : ", new_a)
	new_a = atomic.AddInt32(&a, -2)
	fmt.Println("new_a : ", new_a)

	//CAS(Compare And Swap)比较并交换操作,函数名以CompareAndSwap为前缀，并具体类型名
	var b int32
	fmt.Println("b : ", b)
	//函数会先判断参数一指向的值与参数二是否相等，如果相等，则用参数三替换参数一的值,最后返回是否替换成功
	atomic.CompareAndSwapInt32(&b, 0, 3)
	fmt.Println("b : ", b)

	//Load函数来确保我们正确的读取,函数名以Load为前缀，加具体类型名
	var c int32
	wg := sync.WaitGroup{}
	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			tmp := atomic.LoadInt32(&c)
			if !atomic.CompareAndSwapInt32(&c, tmp, tmp + 1) {
				fmt.Println("c 修改失败")
			}
		}()
	}
	wg.Wait()
	fmt.Println("c : ", c)

	var d int32
	d++
	fmt.Println("d : ", d)
	//存储某个值时，任何CPU都不会都该值进行读或写操作,存储操作总会成功，它不关心旧值是什么，与CAS不同
	atomic.StoreInt32(&d, 666)
	fmt.Println("d : ", d)

	//直接设置新值，返回旧值，与CAS不同，它不关心旧值,函数名以Swap为前缀，加具体类型名
	var e int32
	wg2 := sync.WaitGroup{}
	for i := 0; i < 10; i++ {
		wg2.Add(1)
		go func() {
			defer wg2.Done()
			tmp := atomic.LoadInt32(&e)
			old := atomic.SwapInt32(&e, tmp+1)
			fmt.Println("e old : ", old)
		}()
	}
	wg2.Wait()
	fmt.Println("e : ", e)
}

func test12()  {
	p := message.NewPrinter(language.English)
	p.Printf("%d", currency.Symbol(currency.USD.Amount(0.1)))
	fmt.Println()
	p.Printf("%d", currency.NarrowSymbol(currency.JPY.Amount(1.6)))
	fmt.Println()
	p.Printf("%d", currency.ISO.Kind(currency.Cash)(currency.EUR.Amount(12.255)))
	fmt.Println()
}

func test11()  {
	//fmt.Println(utf8.RuneLen('\u0002'))
}

func test10() {
	//fmt.Println(net.InterfaceAddrs())
	fmt.Println("program initContext...")
	ifaces, _ := net.Interfaces()
	for _, i := range ifaces {
		addrs, _ := i.Addrs()
		for _, addr := range addrs {
			match, _ := regexp.MatchString(`^[0-9]+\.[0-9]+\.[0-9]+\.[0-9]+/[0-9]+$`, addr.String())
			if !match {
				continue
			}
			slit := strings.Split(addr.String(), "/")
			fmt.Println(i.Name, i.Flags, slit)
		}
	}
}

func test09() {
	fmt.Println(3617.1 + 27.1)
}

//表情码测试
func test08() {
	fmt.Println(UnicodeEmojiDecode("哈哈[\u1f604]"))
	//fmt.Println(UnicodeEmojiEnCode("哈哈😄"))

}

//表情解码
func UnicodeEmojiDecode(s string) string {
	//emoji表情的数据表达式
	re := regexp.MustCompile("\\[[\\\\u0-9a-zA-Z]+\\\\]")
	//提取emoji数据表达式
	reg := regexp.MustCompile("\\[\\\\u|]")
	src := re.FindAllString(s, -1)
	for i := 0; i < len(src); i++ {
		e := reg.ReplaceAllString(src[i], "")
		p, err := strconv.ParseInt(e, 16, 32)
		if err == nil {
			s = strings.Replace(s, src[i], string(rune(p)), -1)
		}
	}
	return s
}
//表情转换
func UnicodeEmojiEnCode(s string) string {
	ret := ""
	rs := []rune(s)
	for i := 0; i < len(rs); i++ {
		if len(string(rs[i])) == 4 {
			u := `[\u` + strconv.FormatInt(int64(rs[i]), 16) + `]`
			ret += u
		} else {
			ret += string(rs[i])
		}
	}
	return ret
}

func test07() {
	var i1 interface{} = f1()
	fmt.Println(i1 == nil) // false, why?
	i2 := f2(-1)
	fmt.Println(i2 == nil) // false, why?
	f3(-1)
}
func f1() *bool {
	return nil
}
func f2(length int) interface{} {
	var ss []int = nil
	if length >= 0 {
		ss = make([]int, length)
	}
	return ss
}
func f3(length int) interface{} {
	var ss []int = nil
	if length >= 0 {
		ss = make([]int, length)
	}
	fmt.Println(ss == nil)
	return ss
}

func test06() {
	a0 := [8]int{789, 0, 0, 0, 0, 0, 0, 0}
	a1 := [8]int{0: 789}
	a2 := [...]int{789, 0, 0, 0, 0, 0, 0, 0}
	a3 := [...]int{0: 789, 7: 0}
	var a4 [8]int
	a4[0] = 789
	a5 := *new([8]int)
	a5[0] = 789
	fmt.Println(a0, a1, a2, a3, a4, a5)
}

//chan用法-缓存
func test05() {
	values := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	resultChan := make(chan int, 3)
	go sum(values[:len(values)/2], resultChan)
	go sum(values[len(values)/2:], resultChan)
	go sum(values[len(values)/3:], resultChan)
	sum1, sum2, sum3 := <-resultChan, <-resultChan, <-resultChan
	fmt.Println("Result:", sum1, sum2, sum3)
}
func sum(values []int, resultChan chan int) {
	sum := 0
	for _, value := range values {
		sum += value
	}
	// 将计算结果发送到channel中
	resultChan <- sum
}

//chan用法-缓存
func test04() {
	ch1 := make(chan int, 10)
	for i := 1; i < 10; i++ {
		ch1 <- i
	}
	// 显式关闭
	close(ch1)
	v, ok := <-ch1
	for ok {
		fmt.Println("out:", v, ok)
		v, ok = <-ch1
		if !ok {
			break
		}
	}
	time.Sleep(time.Second * 3)
}

//标准获取控制台字符与输出
func test03() {
	//从标准输入读取数据
	inputReader := bufio.NewReader(os.Stdin)
	fmt.Println("Please input your name：")
	//读取数据直到遇见\n位置
	input, err := inputReader.ReadString('\n')
	if err != nil {
		fmt.Printf("An error occurred:%s\n", err)
		//异常错误后退出
		os.Exit(1)
	} else {
		//用切片操作删除最后的\n
		name := input[:len(input)-1]
		fmt.Printf("Hello,%s!What can i di for you?\n", name)
	}
	for {
		input, err = inputReader.ReadString('\n')
		if err != nil {
			fmt.Printf("An error occurred:%s\n", err)
			continue
		}
		input = input[:len(input)-1]
		//全部转换为小写
		input = strings.ToLower(input)
		switch input {
		case "":
			continue
		case "nothing", "bye":
			fmt.Println("Bye!")
			//正常退出
			os.Exit(0)
		default:
			fmt.Println("Sorry,I didn't catch you.")

		}
	}
}

//字符处理test
func test02() {
	a := "aaa \r\n "
	var bu bytes.Buffer
	bu.WriteString(fmt.Sprintf("%q", a)) //带引号
	fmt.Println(bu.String())
	type point struct {
		x, y int
	}
	//Go 为常规 Go 值的格式化设计提供了多种打印方式。例如，这里打印了 point 结构体的一个实例。
	p := point{1, 2}
	fmt.Printf("%v\n", p) // {1 2}
	//如果值是一个结构体，%+v 的格式化输出内容将包括结构体的字段名。
	fmt.Printf("%+v\n", p) // {x:1 y:2}
	//%#v 形式则输出这个值的 Go 语法表示。例如，值的运行源代码片段。
	fmt.Printf("%#v\n", p) // main.point{x:1, y:2}
	//需要打印值的类型，使用 %T。
	fmt.Printf("%T\n", p) // main.point
	//格式化布尔值是简单的。
	fmt.Printf("%t\n", true)
	//格式化整形数有多种方式，使用 %d进行标准的十进制格式化。
	fmt.Printf("%d\n", 123)
	//这个输出二进制表示形式。
	fmt.Printf("%b\n", 14)
	//这个输出给定整数的对应字符。
	fmt.Printf("%c\n", 33)
	//%x 提供十六进制编码。
	fmt.Printf("%x\n", 456)
	//对于浮点型同样有很多的格式化选项。使用 %f 进行最基本的十进制格式化。
	fmt.Printf("%f\n", 78.9)
	//%e 和 %E 将浮点型格式化为（稍微有一点不同的）科学技科学记数法表示形式。
	fmt.Printf("%e\n", 123400000.0)
	fmt.Printf("%E\n", 123400000.0)
	//使用 %s 进行基本的字符串输出。
	fmt.Printf("%s\n", "\"string\"")
	//像 Go 源代码中那样带有双引号的输出，使用 %q。
	fmt.Printf("%q\n", "\"string\"")
	//和上面的整形数一样，%x 输出使用 base-16 编码的字符串，每个字节使用 2 个字符表示。
	fmt.Printf("%x\n", "hex this")
	//要输出一个指针的值，使用 %p。
	fmt.Printf("%p\n", &p)
	//当输出数字的时候，你将经常想要控制输出结果的宽度和精度，可以使用在 % 后面使用数字来控制输出宽度。默认结果使用右对齐并且通过空格来填充空白部分。
	fmt.Printf("|%6d|%6d|\n", 12, 345)
	//你也可以指定浮点型的输出宽度，同时也可以通过 宽度.精度 的语法来指定输出的精度。
	fmt.Printf("|%6.2f|%6.2f|\n", 1.2, 3.45)
	//要最对齐，使用 - 标志。
	fmt.Printf("|%-6.2f|%-6.2f|\n", 1.2, 3.45)
	//你也许也想控制字符串输出时的宽度，特别是要确保他们在类表格输出时的对齐。这是基本的右对齐宽度表示。
	fmt.Printf("|%6s|%6s|\n", "foo", "b")
	//要左对齐，和数字一样，使用 - 标志。
	fmt.Printf("|%-6s|%-6s|\n", "foo", "b")
	//到目前为止，我们已经看过 Printf了，它通过 os.Stdout输出格式化的字符串。Sprintf 则格式化并返回一个字符串而不带任何输出。
	s := fmt.Sprintf("a %s", "string")
	fmt.Println(s)
	//你可以使用 Fprintf 来格式化并输出到 io.Writers而不是 os.Stdout。
	fmt.Fprintf(os.Stderr, "an %s\n", "error")
}

func test01() {
	fmt.Println("hi")
	time.Sleep(30 * time.Second)
}

//判断包含中文
func IsChineseChar() {
	str := "你好"
	fmt.Println(unicode.Is(unicode.Scripts["Han"], []rune(str)[0]))
	for _, r := range str {
		if unicode.Is(unicode.Scripts["Han"], r) {
			fmt.Println("中文：", r)
			break
		} else {
			fmt.Println("英文：", r)
			break
		}
	}
}

//中文转换unicode码
func HanToUnicode() {
	sText := "中文"
	textQuoted := strconv.QuoteToASCII(sText)
	textUnquoted := textQuoted[1 : len(textQuoted)-1]
	fmt.Println(textUnquoted)
	sUnicodev := strings.Split(textUnquoted, "\\u")
	var context string
	for _, v := range sUnicodev {
		if len(v) < 1 {
			continue
		}
		temp, err := strconv.ParseInt(v, 16, 32)
		if err != nil {
			panic(err)
		}
		context += fmt.Sprintf("%c", temp)
	}
	fmt.Println(context)
}

//大小写
func stringMe() {
	fmt.Println(strings.ToLower("你好nIHaOa啊"))
	fmt.Println(strings.ToUpper("你好nIHaOa啊"))
}

//切片删除元素
func slices() {
	s := []string{"e", "b", "f"}
	ss := []string{"a", "b", "c", "d", "e", "f"}

	for _, v := range s {
		for kk, vv := range ss {
			if v == vv {
				ss = append(ss[:kk], ss[(kk+1):]...)
			}
			fmt.Println(vv)
		}
	}
	fmt.Println(ss)
}
