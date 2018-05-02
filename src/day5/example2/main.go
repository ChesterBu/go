package main

func main()  {
	for i := 0; i < 3; i++ {
		defer func() {
			println(i)
		}()
	}
	c()
	b()
}

func b()  {
	panic("panic in b")
}

func c()  {
	defer func() {
		if err := recover(); err != nil {
			println("recover")
		}
	}()
	panic("painc in c")

}