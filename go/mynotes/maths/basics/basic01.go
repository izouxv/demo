package main

import (
	"encoding/json"
	"fmt"
	"runtime"
	"sync"
	"time"
	"reflect"
)

func main() {
	//itest01()
	//itest02()
	//itest03()
	//itest04()
	//itest05()
	//itest06()
	//itest07()
	//itest08()
	//itest09()
	//iTest10()
	//iTest11()
	//iTest12()
	//iTest13()
	//iTest14()
	//iTest15()
	//iTest17()
	//iTest18()
	iTest19()
}

//recover只会捕捉最后一个panic，panic执行顺序与覆盖性
func iTest19() {
	defer func() {
		if err:=recover();err!=nil{
			fmt.Println("++++")
			f:=err.(func()string)
			fmt.Println(err,f(),reflect.TypeOf(err).Kind().String())
		}else {
			fmt.Println("fatal")
		}
	}()

	defer func() {
		panic(func() string {
			return  "defer panic"
		})
	}()
	panic("panic")
}

//go1.9新特性
/**
因为MyUser2完全等价于User，所以具有其所有的方法，并且其中一个新增了方法，另外一个也会有。
但是i1.m2()是不能执行的，因为MyUser1没有定义该方法
 */
func iTest18() {
	var i1 MyUser1
	var i2 MyUser2
	i1.m1()
	i2.m2()
}
type User struct {}
type MyUser1 User
type MyUser2 = User
func (i MyUser1) m1(){
	fmt.Println("MyUser1.m1")
}
func (i User) m2(){
	fmt.Println("User.m2")
}
/**
基于一个类型创建一个新类型，称之为defintion；基于一个类型创建一个别名，称之为alias。
MyInt1为称之为defintion，虽然底层类型为int类型，但是不能直接赋值，需要强转；
MyInt2称之为alias，可以直接赋值。
 */
func iTest17() {
	type MyInt1 int
	type MyInt2 = int
	var i int =9
	var i1 MyInt1
	var i2 MyInt2
	i1 = MyInt1(i)
	i2 = i
	fmt.Println(i1,i2)
}

//goto的使用：不能跳转到函数中或内层代码
func iTest16() {
//	for i:=0;i<10 ;i++  {
//loop:
//		println(i)
//	}
//	goto loop
}

//常量的使用：常量不同于变量的在运行期分配内存，常量通常会被编译器在预处理阶段直接展开，作为指令数据使用
func iTest15()  {
	println(i15c,i15c)
	println(&i15v,i15v)
}
const i15c = 100
var i15v = 100

//iota的使用与位置使用
func iTest14()  {
	fmt.Println(x,y,z,k,p)
}
const (
	x = iota
	y
	z = "zz"
	k
	p = iota
)


//new 与make 的使用区别
//1、make只能用来分配及初始化类型为slice，map，chan的数据；new可以分配任意类型的数据
//2、new分配返回的是指针，即类型*T；make返回引用，即T
//3、new分配的空间被清零，make分配后，会进行初始化
func iTest13() {
	//list := new([]int)
	list := make([]int,0)
	list = append(list, 1)
	fmt.Println(list)
}

//interface内部结构,
// 一种为var的接口，var inter interface{}
// 一种为带方法的接口，比如下例
func iTest12() {
	if live() == nil {
		fmt.Println("AAAAAAA")
	} else {
		fmt.Println("BBBBBBB")
	}
}
type Inter1 interface {
	Show()
}
type Student struct{}
func (stu *Student) Show() {
}
func live() Inter1 {
	var stu *Student
	return stu
}


//结构体的导出与非导出变量（首字母）
func iTest11() {
	js := `{ "age": 1 }`
	var test Test
	json.Unmarshal([]byte(js), &test)
	fmt.Println(test)
}
type Test struct {
	age int32 `json:"age"`
}

//接口与结构体方法
func iTest10() {
	var peo demo01 = &demo02{}
	think := "bitch"
	fmt.Println(peo.Speak(think))
}
type demo01 interface {
	Speak(string) string
}
type demo02 struct{}
func (stu *demo02) Speak(think string) (talk string) {
	if think == "bitch" {
		talk = "You are a good boy"
	} else {
		talk = "hi"
	}
	return
}

//chan缓存池的应用
func iTest09() {
	th := threadSafeSet{
		s: []interface{}{"1", "2", "3", "4"},
	}
	v := <-th.Iter()
	//fmt.Sprintf("%s%v","ch",v)
	fmt.Println(v)
}
type threadSafeSet struct {
	sync.RWMutex
	s []interface{}
}
func (set *threadSafeSet) Iter() <-chan interface{} {
	ch := make(chan interface{}, 2) //chan大小的区别
	go func() {
		set.RLock()
		for k, v := range set.s {
			ch <- k
			println("aaa:", k, v.(string))
		}
		close(ch)
		set.RUnlock()
	}()
	return ch
}

//map线程安全
/**
Set加锁，Get也得加锁,最好使用 sync.RWMutex
*/
func iTest08() {
	runtime.GOMAXPROCS(1)
	u := &UserAges{ages: map[string]int{"aaa": 1, "bbb": 2}}
	u.Add("ccc", 3)
	fmt.Println("get:", u.Get("ccc"))
	fmt.Println(u)
}
type UserAges struct {
	ages map[string]int
	sync.Mutex
}
func (ua *UserAges) Add(name string, age int) {
	ua.Lock()
	defer ua.Unlock()
	ua.ages[name] = age
}
func (ua *UserAges) Get(name string) int {
	//ua.Lock()
	//defer ua.Unlock()
	if age, ok := ua.ages[name]; ok {
		return age
	}
	return -1
}

//make默认值和append操作的区别
func iTest07() {
	s := make([]int, 5)
	s = append(s, 1, 2, 3)
	fmt.Println(s)
}

//函数与defer的执行顺序
func iTest06() {
	a := 1
	b := 2
	defer calc("1", a, calc("10", a, b))
	a = 0
	defer calc("2", a, calc("20", a, b))
	b = 1
}
func calc(index string, a, b int) int {
	ret := a + b
	fmt.Println(index, a, b, ret)
	return ret
}

//select执行的随机性
/**
1.select中只要有一个case能return，则立刻执行。
2.当如果同一时间有多个case均能return则伪随机方式抽取任意一个执行。
3.如果没有一个case能return则可以执行default块。
*/
func iTest05() {
	runtime.GOMAXPROCS(1)
	int_chan := make(chan int, 1)
	string_chan := make(chan string, 1)
	//int_chan <- 1
	//string_chan <- "hello"
	select {
	case value := <-int_chan:
		fmt.Println(value)
	case value := <-string_chan:
		panic(value)
	}
}

//结构体的用法与实现
/**
golang的组合模式，实现OOP继承，必须是匿名字段
*/
func iTest04() {
	t := Teacher{People:People{A:"111"}}
	fmt.Println(t.A)
}
type People struct{
	A	string
}
func (p *People) ShowA() {
	fmt.Println("showA-p")
	p.ShowB()
}
func (p *People) ShowB() {
	fmt.Println("showB-p")
}
type Teacher struct {
	People
}
func (t *Teacher) ShowA() {
	fmt.Println("showA-t")
}
func (t *Teacher) ShowB() {
	fmt.Println("showB-t")
}

//GOMAXPROCS的作用
/**
Golang默认所有的任务都在一个cpu核里，如果想使用多核来跑goroutine的任务，需要配置runtime.GOMAXPROCS。
GOMAXPROCS的数目根据自己任务量分配就可以了，有个前提是不要大于你的cpu核数。
并行比较适合那种cpu密集型计算，如果是IO密集型使用多核的化会增加cpu切换的成本。
*/
func iTest03() {
	runtime.GOMAXPROCS(2)
	wg := sync.WaitGroup{}
	wg.Add(10)
	for i := 0; i < 10; i++ {
		fmt.Println("1-i:", i)
		go func() {
			fmt.Println("1-go-i: ", i)
			wg.Done()
		}()
	}
	for i := 0; i < 10; i++ {
		go func(i int) {
			fmt.Println("2-i: ", i)
			wg.Done()
		}(i)
	}
	wg.Wait()
}

//指针的用法-错误示例
func iTest02() {
	type test struct {
		Name string
	}
	m := make(map[string]*test)
	te := []test{
		{Name: "test01"}, {Name: "test02"}, {Name: "test03"},
	}
	//注意指针的引用于创建指针，使用指针变量要注意指针的引用与golang是否会避免重复分配内存
	for _, stu := range te {
		fmt.Printf("%p\n", &stu)
		fmt.Println(&stu)
		m[stu.Name] = &stu
	}
	for _, mm := range m {
		fmt.Println(mm)
	}
}

//defer与panic异常与recover的执行顺序
/**
defer是后入先出，panic等defer结束后再执行
*/
func iTest01() {
	defer func() { fmt.Println("打印前", time.Now()) }()
	time.Sleep(time.Second)
	defer func() { fmt.Println("打印中", time.Now()) }()
	time.Sleep(time.Second)
	defer func() { fmt.Println("打印后", time.Now()) }()
	defer func() { fmt.Println("打印", recover()) }()
	time.Sleep(time.Second)
	panic("触发异常" + time.Now().String())
}
