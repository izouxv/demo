package main

import (
	"fmt"
	"time"
	"math/rand"
)

var a = []int{4, 2, 8, 5, 1, 9, 6, 11, 0, 4, 2, 9, 7, 3, 12, 43, 23, 12, 0, 97, 4}

func main() {
	quickSort(a, 0, len(a)-1)
	fmt.Println(a)
	//sort02()
	//heapSort()
	//cocktailSort()
	//isPrimeNumber()
}

//素数判断
func isPrimeNumber() {
	for value:=2; value < 30 ;value++ {
		if value <= 3 {
			fmt.Println("PrimeNumber:",value)
		}
		if value%2 == 0 || value%3 == 0 {
			continue
		}
		for j := 5; j*j <= value; j += 6 {
			if value%j == 0 || value%(j+2) == 0 {
				goto breaks
			}
		}
		fmt.Println("PrimeNumber:",value)
breaks:
	}
}

//双向冒泡排序
func cocktailSort() {
	left := 0
	right := len(a) - 1
	flag := true
	for left < right {
		flag = true
		for i :=left; i < right; i++{//将最大元素放到后面
			if a[i] > a[i+1] {
				a[i], a[i+1] = a[i+1], a[i]
				flag = false
			}
		}
		right--
		if flag {//序列访问没有发生交换--表明序列有序
			break
		}
		for i :=right; i > left; i--{// 后半轮,将最小元素放到前面
			if a[i] < a[i-1] {
				a[i], a[i-1] = a[i-1], a[i]
				flag = false
			}
		}
		left++
		if flag {//如果此次序列访问没有发生交换--表明序列有序
			break
		}
	}
	fmt.Println(a)
}

//堆排序
func heapSort() {
	//array := []interface{}{4, 2, 8, 5, 1, 9, 6, 11, 0, 4, 2, 9, 7, 3, 12, 43, 23, 12, 0, 97, 4}
	var array []interface{}
	for a:=0; a< 50;a++ {
		array = append(array,rand.Intn(100))
	}
	fmt.Println(array)
	a := time.Now()
	fmt.Println(a)
	headSoft(array, func(a, b interface{}) bool {
		a1 := a.(int)
		a2 := b.(int)
		return a1 <= a2
	})
	fmt.Println(time.Now().Sub(a))
	fmt.Println(array)
}
func headSoft(array []interface{}, cmp func(a, b interface{}) bool) {
	size := len(array)
	//建堆
	for i := (size / 2) - 1; i >= 0; i-- {
		adjustHeap(array, i, size, cmp)
	}
	//排序
	for i := 0; i < size; i++ {
		sz := size - i - 1
		array[0], array[sz] = array[sz], array[0]
		adjustHeap(array[0:sz], 0, sz, cmp)
	}
}
func adjustHeap(array []interface{}, i, len int, cmp func(a, b interface{}) bool) {
	for ch := i<<1 + 1; ch < len; ch = i<<1 + 1 {
		rg := i<<1 + 2
		if rg < len && cmp(array[ch], array[rg]) {
			ch = rg
		}
		if cmp(array[ch], array[i]) {
			return
		}
		array[ch], array[i] = array[i], array[ch]
		i = ch
	}
}

//快速排序
func quickSort(array []int, begin, end int) {
	if begin < end {
		var i, j int
		i,j = begin + 1,end
		// 将array[begin]作为基准数，从array[begin+1]与基准数比较
		// array[end]是数组的最后一位
		for i<j {
			if array[begin] < array[i] {
				array[i], array[j] = array[j], array[i]
				j--
			} else {
				i++
			}
		}
		/* 跳出while循环后，i = j,数组被分割成两个部分
		-->  array[begin+1] ~ array[i-1] < array[begin]
		-->  array[i+1] ~ array[end] > array[begin]
		 * 这个时候将数组array分成两个部分，再将array[i]与array[begin]进行比较，决定array[i]的位置。
		 * 最后将array[i]与array[begin]交换，进行两个分割部分的排序！以此类推，直到最后i = j不满足条件退出
		 */
		//处理数组元素相同值时取>=
		if array[i] >= array[begin] {
			i = i - 1
		}
		array[begin], array[i] = array[i], array[begin]
		quickSort(array, begin, i)
		quickSort(array, j, end)
	}
}