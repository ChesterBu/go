package main

import (
	"os"
	"fmt"
	"io"
	"bufio"
)

func writeFile(path string) {
	if f, err := os.Create(path); err == nil {
		defer f.Close()
		for i := 0; i < 10; i++ {
			buf := fmt.Sprintf("i=%d\n", i)
			if n, err := f.WriteString(buf); err == nil {
				fmt.Println(n) //写入长度
			}
		}
	}
}

func readFile(path string) {
	if f, err := os.Open(path); err == nil {
		defer f.Close()

		buf := make([]byte, 2048)
		if n, err := f.Read(buf); err == nil&& err==io.EOF {
			fmt.Println(string(buf[:n])) //读取正好的长度
		}

	}
}

func readFileLine(path string) {
	if f, err := os.Open(path); err == nil {
		defer f.Close()
		r:=bufio.NewReader(f)
		for   {
			buf,err:=r.ReadBytes('\n')
			if err!=nil && err == io.EOF{
				break
			}
			fmt.Println(string(buf))
		}

	}
}

func copyFile(src,dst string)  {
	if sF,err1 := os.Open(src);err1==nil{
		defer sF.Close()
		if dF,err2 := os.Create(dst);err2==nil{
			defer dF.Close()
			buf:=make([]byte,4096)  //建立缓冲区
			for{
				n,err:=sF.Read(buf)   //将文件内容读到buf中
				if err != nil && err == io.EOF{ //结束就停止
					break
				}
				dF.Write(buf[:n]) //写入
			}
		}
	}

}
func main() {
	//os.Stdout.WriteString("are u ok ?")
	//writeFile("./a.txt")
	//readFile("./a.txt")
	//readFileLine("./a.txt")
	copyFile("./a.txt","./b.txt")
}
