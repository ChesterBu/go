package main

import "fmt"

func showNew() {
	j := new(int) //new返回指针 make返回数据结构
	*j = 100
	fmt.Println(*j) //100

	//new返回指针 make返回数据结构
	s1 := new([]int)
	fmt.Println(s1) //&[]
	s2 := make([]int, 10)
	fmt.Println(s2) //[0 0 0 0 0 0 0 0 0 0]

}

func showAppend() {
	var a []int
	b := append(a, 1, 2, 3)
	fmt.Println(a, b) //[] [1 2 3]
}

func showArr() {
	//数组是值类型
	var a [10]int
	b := a
	b[0] = 100
	fmt.Println(a,b) //[0 0 0 0 0 0 0 0 0 0] [100 0 0 0 0 0 0 0 0 0]

}

func main() {
	//showNew()
	//showAppend()
	showArr()
}
