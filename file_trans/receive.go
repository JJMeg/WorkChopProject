package main

import (
	"fmt"
	"io"
	"net"
	"os"
)

type Receiver struct {
	net.Conn
}

func NewReceiver(c net.Conn) *Receiver {
	return &Receiver{
		c,
	}
}

func (receiver *Receiver) RecvFile(fileName string) {
	file, err := os.Create(fileName)
	if err != nil {
		fmt.Println("os.Create file error: ", err)
		return
	}
	defer file.Close()

	//	从conn中读取数据，写入到本地文件
	for {
		buf := make([]byte, 4096)
		n, err := receiver.Read(buf)
		// 读多少 写多少
		_, err = file.Write(buf[:n])
		if err != nil {
			if err == io.EOF {
				fmt.Printf("接收文件完成。\n")
			} else {
				fmt.Printf("conn.Read()方法执行出错，错误为:%v\n", err)
			}
			return
		}
	}
}

func DoReceive() {
	listener, err := net.Listen("tcp", "localhost:8000")
	if err != nil {
		fmt.Println("net listen error: ", err)
		return
	}
	defer listener.Close()

	//阻塞监听
	conn, err := listener.Accept()
	receiver := NewReceiver(conn)
	defer receiver.Close()

	//	文件名
	buf := make([]byte, 1024)
	n, err := conn.Read(buf)
	if err != nil {
		fmt.Println("conn read error: ", err)
		return
	}
	fileName := string(buf[:n])
	conn.Write([]byte("ok"))

	receiver.RecvFile(fileName)
}
