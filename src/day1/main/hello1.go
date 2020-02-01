package main

import (
	. "fmt"
	"../add"
)


func sum(a int, b int) int {
	return a + b
}

func main() {
	c := sum(1, 2)
	for i := 0; i < 10; i++ {
		Println(i)
	}

	pipe := make(chan int ,1)
	go add.Add(100,300,pipe)
	sum := <- pipe
	Println(sum)
	Println(add.Name)

	Println("hello world", c)
}
