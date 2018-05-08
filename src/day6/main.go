package main

import "fmt"

func main()  {
var c bool
just(c)

}

func just(a interface{} )  {
	switch a.(type) {
	case bool:
		fmt.Println("It's boolean")
	}
}
