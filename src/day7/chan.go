package main

import (
	"fmt"
	"time"
)

var ch = make(chan int)

func printer(s string) {
	for _, value := range s {
		fmt.Printf("%c", value)
		time.Sleep(time.Second)
	}
}

func print1() {
	printer("hello")
	ch <- 666
}

func print2() {
	<-ch
	printer("world")
}

func main() {
	go print1()
	go print2()

	time.Sleep(20*time.Second) //等待goroutine
}
