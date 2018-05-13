package main

import "fmt"

type person struct {
	id int
}


//方法是值拷备的
func (p person) printIfo()  {
	p.id = 999
	fmt.Println(p.id)
}
//所以要改变值要传指针
func (p *person)changeInfo()  {
	p.id = 0
}

func main()  {
	p:=person{123}
	p.printIfo()//999
	fmt.Println(p.id)//123
	//(&p).changeInfo()
	p.changeInfo()  //效果同上
	fmt.Println(p.id) //0

	p2 := new(person)
	p2.id= 456
	p2.printIfo()//999
	p2.changeInfo()
	fmt.Println(p2.id) //0

}