package main

import "fmt"

type student struct {
	name string
	age int32
	score float64
}

func main()  {
	var stu student
	stu.name = "Gavin"
	stu.age = 18
	stu.score = 65.0
	fmt.Printf("name: %p\n" ,&stu.name)
	fmt.Printf("age: %p\n" ,&stu.age)
	fmt.Printf("score: %p\n" ,&stu.score)


}