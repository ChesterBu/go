package main

import (
	"net"
	"fmt"
)

func main() {
	listener,err:=net.Listen("tcp",":8080")
	if err !=nil{
		return
	}
	defer listener.Close()
	conn,err1 := listener.Accept()
	if err1 !=nil{
		return
	}
	defer conn.Close()

	buf := make([]byte,1024)
	n,err2 := conn.Read(buf)
	if err2 !=nil{
		return
	}
	fmt.Printf("%v",string(buf[:n]))
}