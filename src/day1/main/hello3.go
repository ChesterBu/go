package main

import "fmt"


var a [3]int

var b = make([]int, 2, 10)


func changeArray(a [3]int){
	a[2] =2
}

func changeArray2(a *[3]int){
	// a[2] =2 下面的语法糖
	(*a)[2] = 2
}

func changeArray3(a []int){
	a[1] = 2
}

func changeArray4(a []int){
	a = append(a,1)
	fmt.Println(a) // [0,2,1]
}

func changeArray5(a *[]int){
	// a = append(*a,1)  // cannot use append(*a, 1) (type []int) as type *[]int in assignment
	fmt.Println(a) // [0,2,1]
}

func main(){
	fmt.Println(a)
	changeArray(a)
	fmt.Println(a) // 0 0 0
	changeArray2(&a) 
	fmt.Println(a) // 0 0 2
	// 切片的分割线
	fmt.Println(b)  // 0 0
	changeArray3(b)
	fmt.Println(b)  // 0 2
	changeArray4(b)
	fmt.Println(b) // 0 2
	// changeArray5(&b)
	// fmt.Println(b)
	c := make([]string, 5, 10)
	fmt.Println(c) // [    ]
	for i := 0; i < 10; i++ {
		c = append(c, fmt.Sprintf("%v", i))
	}
	fmt.Println(c)  // [     0 1 2 3 4 5 6 7 8 9]
}




