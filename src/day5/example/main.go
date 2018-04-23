package main

import "fmt"

type student struct {
	name string
	age int32
	score float64
}

func (ex student) init()  {
	fmt.Println(ex.score)
}

func main()  {
	var stu student
	stu.name = "Gavin"
	stu.age = 18
	stu.score = 65.0
	stu.init()
	fmt.Printf("name: %p\n" ,&stu.name)
	fmt.Printf("age: %p\n" ,&stu.age)
	fmt.Printf("score: %p\n" ,&stu.score)


}