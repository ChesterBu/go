package main

import (
	"fmt"
	"strings"
	"strconv"
	"time"
)

func tryDefer()  {
	i:=0
	defer fmt.Println(i)  //0
	defer fmt.Println("second")
	i=10
	fmt.Println(i)
}

func add(arg ...int)  {
	fmt.Println(arg)
}

func main() {
	var str = "hello worldh"
	fmt.Println(strings.Index(str, "e"))
	fmt.Println(strings.Trim(str, "h"))
	fmt.Println(strings.Split(str, ""))
	fmt.Println(strings.Fields(str))
	fmt.Println(strconv.Itoa(1000))
	now := time.Now()
	fmt.Println(now.Format("2006/01/02 15:04:05"))
	fmt.Println(&str)
	for i,v:=range str{
		fmt.Printf("index=%d,val=%c \n",i,v)
	}
	add(1,2,3,4)
	tryDefer()

}
