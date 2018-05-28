package main

import (
	"net"
	"strings"
	"time"
)

type Client struct {
	C    chan string
	Name string //用户名
	Addr string //地址
}

var onLine map[string]Client

var message = make(chan string)

func writeMessage(cli Client, conn net.Conn) {
	for msg := range cli.C {
		conn.Write([]byte(msg + "\n"))
	}
}
func makeMessage(cli Client, msg string) (buf string) {
	buf = "[" + cli.Addr + "]" + cli.Name + ": " + msg
	return
}
func onConn(conn net.Conn) {
	defer conn.Close()
	cliAddr := conn.RemoteAddr().String()
	cli := Client{make(chan string), cliAddr, cliAddr}
	onLine[cliAddr] = cli
	//新开协程给客户端发送信息
	go writeMessage(cli, conn)
	//广播某个人在线
	message <- makeMessage(cli, "online")

	isQuit := make(chan bool)
	hasData := make(chan bool)
	//接受用户发来的信息
	go func() {
		buf := make([]byte, 2048)
		for {
			n, _ := conn.Read(buf)
			if n == 0 {
				isQuit <- true
				return
			}
			msg := string(buf[:n])
			if len(msg) == 3 && msg == "who" {
				//遍历map给当前用户发送所有成员
				for _, tmp := range onLine {
					msg = tmp.Addr + ":" + tmp.Name + "\n"
					conn.Write([]byte(msg))
				}
			} else if len(msg) >= 8 && msg[:6] == "rename" {
				//rename|mike
				name := strings.Split(msg, "|")[1]
				cli.Name = name
				onLine[cliAddr] = cli
				conn.Write([]byte("done"))

			} else {
				message <- makeMessage(cli, msg)
			}
			hasData <- true

		}

	}()

	for {
		select {
		case <-isQuit:
			delete(onLine, cliAddr)
			message <- makeMessage(cli, "logOut")
			return
		case <-hasData:
		case <-time.After(60*time.Second): //60s后
			delete(onLine, cliAddr)
			message <- makeMessage(cli, "timeOut")
			return
		}
	}
}
func onMessage() {
	onLine = make(map[string]Client)
	for {
		msg := <-message //没有消息前这里会阻塞
		//	给每个人都发送消息
		for _, cli := range onLine {
			cli.C <- msg
		}
	}
}

func main() {
	listener, err := net.Listen("tcp", ":8080")
	if err != nil {
		return
	}
	defer listener.Close()
	//新开一个协程，只要有消息来了给map每个成员都发送消息
	go onMessage()
	//主协程 循环阻塞等待用户接入
	for {
		conn, err := listener.Accept()
		if err != nil {
			continue
		}
		go onConn(conn) //处理用户链接
	}
}
