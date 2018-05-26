package main

import (
	"os"
	"net"
	"fmt"
	"io"
)

func sendFile(path string,conn net.Conn){
	f,err := os.Open(path)
	if err != nil {
		return
	}
	defer f.Close()
	buf := make([]byte, 1024)
	for  {
		n,err := f.Read(buf)
		if err != nil {
			if err == io.EOF{
				fmt.Println("文件发送完毕")
			}
			return
		}
		conn.Write(buf[:n])
	}
}

func main() {
	list := os.Args
	if len(list) != 2 {
		return
	}
	fileName := list[1]
	info, err := os.Stat(fileName)
	if err != nil {
		return
	}

	conn, err := net.Dial("tcp", ":8080")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer conn.Close()

	//发送文件名
	_, err = conn.Write([]byte(info.Name()))
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
	if "ok"==string(buf[:n]){
		//发送真正文件内容
		sendFile(fileName,conn)
	}

}
