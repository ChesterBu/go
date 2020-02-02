package main

import "fmt"

const (
	a = 10-iota
	c = 1
	_
	b = 8-iota
)

func sum(x,y int)(sum int){
	sum = x + y
	// 不可以省略return 不然会报错
	return
}
func sum2()int{
	x:=5
	defer func(){
		x++
	}()
	// return流程： 1. 返回值赋值  2.defer执行  3.返回返回值
	return x  //5
}

func sum3()(x int){
	defer func(){
		x++
	}()
	// return流程： 1. 返回值赋值  2.defer执行  3.返回返回值
	return 5  // 6
}

func main()  {
	
	s := sum(1,2)
	fmt.Println(s)
	for key,val := range "asccc" {
		// defer stack
		defer fmt.Println(key,val)
	}
	fmt.Println(a,b)
}