package main

import (
	"net"
	"fmt"
)

func main() {
	listener, err := net.Listen("tcp", ":8080")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer listener.Close()

	conn, err := listener.Accept()
	if err != nil {
		fmt.Println(err)
		return
	}
	buf := make([]byte, 1024)
	n, err1 := conn.Read(buf)
	if err1 != nil {
		fmt.Println(err1)
		return
	}
	fmt.Println(string(buf[:n]))
	defer conn.Close()

}
