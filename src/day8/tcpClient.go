package main

import (
	"net"
	"fmt"
)

func main() {
	conn,err:=net.Dial("tcp",":8080")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer conn.Close()
	conn.Write([]byte("are u ok?"))
}