package main

import (
	"net"
	"fmt"
	"os"
	"io"
)

func receiveFile(fileName string,conn net.Conn)  {
	f,err:=os.Create(fileName)
	if err != nil {
		return
	}
	buf := make([]byte, 1024)
	for  {
		n,err := conn.Read(buf)
		if err != nil {
			if err == io.EOF{
				fmt.Println("文件接受完毕")
			}
			return
		}
		if n==0{
			fmt.Println("文件接受完毕")
		}
		f.Write(buf[:n])
	}

}

func main() {
	listener, err := net.Listen("tcp", ":8080")
	if err != nil {
		return
	}
	defer listener.Close()
	conn, err1 := listener.Accept()
	if err1 != nil {
		fmt.Println(err1)
		return
	}
	buf := make([]byte, 1024)
	n,err2 := conn.Read(buf)
	if err2 != nil {
		fmt.Println(err1)
		return
	}
	//读取到文件名
	filename:=string(buf[:n])
	//回复ok
	conn.Write([]byte("ok"))
	//接受内容
	receiveFile(filename,conn)
}
