package main

import (
	. "fmt"
)

const(
	a= iota
	c = 1
	b= iota
)

var x = 1

func list(n int) {
	for i := 0; i < n; i++ {
		Printf("i=%d,n-i=%d\n", i, n-i)
	}
	Println(a,c,b)
}

func swap (a,b *int) {
	*a,*b = *b,*a
}

func foo (){
	var a int8 = 100
	var b = int16(a)
	Println(a,b)
	Println(x)
}


func main() {
	//f:=1
	//s:=2
	//swap(&f,&s)
	//Println(f,s)
	//list(10)
	//---------
	// go 静态作用域
	x:=12
	Println(x)
	foo()
	func(){
		Println(x)
	}()
}
