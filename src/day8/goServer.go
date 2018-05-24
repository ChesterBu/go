package main

import (
	"net"
	"fmt"
	"strings"
)

func onConn(conn net.Conn) {
	defer conn.Close()
	addr := conn.RemoteAddr().String()
	fmt.Println(addr)
	buf := make([]byte, 1024)
	for{
		n, err := conn.Read(buf)
		if err != nil {
			return
		}
		fmt.Println(string(buf[:n]))
		if string(buf[:n-1]) == "exit"{
			return
		}
		conn.Write([]byte(strings.ToUpper(string(buf[:n]))))
	}


}

func main() {
	listener, err := net.Listen("tcp", ":8080")
	if err != nil {
		return
	}
	defer listener.Close()
	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println(err)
			return
		}
		go onConn(conn)
	}
}
