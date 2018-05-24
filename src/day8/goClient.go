package main

import (
	"net"
	"fmt"
	"os"
)

func main() {
	conn, err := net.Dial("tcp", ":8080")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer conn.Close()

	str := make([]byte, 1024)
	go func() {
		//发送内容
		for {
			n, err := os.Stdin.Read(str)
			if err != nil {
				fmt.Println(err)
				return
			}
			//发送
			conn.Write(str[:n])
		}

	}()
	//收到服务器的内容
	buf := make([]byte, 1024)
	for {
		n, err := conn.Read(buf)
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println(string(buf[:n]))
	}

}
