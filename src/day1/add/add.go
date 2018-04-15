package add

func Add(a int , b int , c chan int)  {
	c <- a+b
}

var Name int

