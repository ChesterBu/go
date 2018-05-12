package main

import (
	"fmt"
	"reflect"
)

func main() {
	var c bool
	just(c)
	test(c)

}

func just(a interface{}) {
	switch a.(type) {
	case bool:
		fmt.Println("It's boolean")
	}
}

func test(a interface{}) {
	fmt.Println(reflect.TypeOf(a))
}
