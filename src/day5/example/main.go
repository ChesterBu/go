package main

import "fmt"



//这样写的化只能在这个文件里用
type student struct {
	name string
	age int32
	score float64
}

//这样写的化外界可以调用，但是不可执行.id的操作，因为id不可见
type Stu struct {
	id int
}
//这样写的化外界可以调用，id也可以改了
type Stu2 struct {
	Id int
}

func (ex student) init()  {
	fmt.Println(ex.score)
}

func test(stu student)  {
	stu.name = "aaaa"
}

func main()  {
	stu :=new(student)
	stu.name = "Gavin"
	stu.age = 18
	stu.score = 65.0
	stu.init()
	p:=new(student)
	p.name = "Gavin"
	p.age = 18
	p.score = 65.0
	fmt.Println(p==stu)  //false
	c:=student{"cc",18,65}
	test(c)
	fmt.Println(c) //{cc 18 65} 所以结构体做函数参数是值传递
	fmt.Printf("name: %p\n" ,&stu.name)
	fmt.Printf("age: %p\n" ,&stu.age)
	fmt.Printf("score: %p\n" ,&stu.score)


}