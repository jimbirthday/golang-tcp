package main

import (
	"fmt"
	"io"
	"net"
	"os"
	"strings"
)

func main() {
	client, err := net.Dial("tcp", "127.0.0.1:8000")
	if err != nil {
		fmt.Println(fmt.Sprintf("listen err :%v", err))
	}
	go func() {
		input := make([]byte, 1024)
		for {
			read, err2 := os.Stdout.Read(input)
			if err2 != nil {
				fmt.Println("input err:", err)
				continue
			}
			fmt.Println("输入的数据:", string(input[:read]))
			client.Write(input[:read])
		}
	}()
	//server := New()
	bytes := make([]byte, 100)
	for {
		conn, err := client.Read(bytes)
		if err != nil {
			if err == io.EOF {
				return
			}
			fmt.Println("read err:", err)
			continue
		}
		fmt.Println("返回的数据:", strings.Replace(string(bytes[:conn]),"\n","n",-1))
	}
}
