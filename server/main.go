package main

import (
	"context"
	"fmt"
	"net"
)

type Server struct {
	que chan net.Conn
	ctx context.Context
}

var Serverice *Server

func New() *Server {
	if Serverice != nil {
		c := make(chan net.Conn, 1)
		s := &Server{
			que: c,
			ctx: context.Background(),
		}
		Serverice = s
	}
	return Serverice
}

//
//func (c *Server) Put(conn net.Conn) {
//	c.que <- conn
//}

func main() {
	listen, err := net.Listen("tcp", "127.0.0.1:8000")
	if err != nil {
		fmt.Println(fmt.Sprintf("listen err :%v", err))
		return
	}
	//server := New()
	for {
		conn, errs := listen.Accept()
		if errs != nil {
			fmt.Println(fmt.Sprintf("Accept err :%v", errs))
		}
		go handleConn(conn)
	}
}

//func ListenQue() {
//	for {
//		time.Sleep(time.Second * 1)
//		select {
//		case conn := <-Serverice.que:
//		default:
//			fmt.Println("Server que is empty")
//		}
//	}
//}

func handleConn(conn net.Conn) {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println(fmt.Sprintf("handleConn recover err :%v", err))
		}
	}()
	defer func() {
		conn.Close()
	}()
	fmt.Println("handleConn ~~~~~")
	if conn == nil {
		fmt.Println("失效的连接")
		return
	}
	bytes := make([]byte, 10)
	for {
		read, err := conn.Read(bytes)
		if err != nil || read == 0 {
			fmt.Println(fmt.Sprintf("Read err :%s", err))
			break
		}
		fmt.Println("接受的数据", string(bytes[:read]))
		conn.Write(bytes)
	}
	fmt.Println("handleConn end ~~~~~")
}
