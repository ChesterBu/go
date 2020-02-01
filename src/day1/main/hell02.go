package main

import "fmt"

const (
	a = 10-iota
	c = 1
	_
	b = 8-iota
)
func main()  {
	fmt.Println(a,b)
}