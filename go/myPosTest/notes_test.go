package main

import (
	"fmt"
	"testing"
)

func Test_channel(t *testing.T) {
	var done chan int
	if done == nil {
		fmt.Println("channel a is nil, going to define it")
		done = make(chan int)
		fmt.Println("Type of a is %T", done)
	}
	go hello(done)
	fmt.Println(done,<- done)

}

func hello(done chan int) {
	fmt.Println("Hello world goroutine")
	done <- 1
}
