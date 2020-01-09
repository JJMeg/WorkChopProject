package main

import (
	"fmt"
	"io"
	"net"
	"os"
)

type Sender struct {
	net.Conn
}

func NewSender(c net.Conn) *Sender {
	return &Sender{
		c,
	}
}

func (sender *Sender) SendFile(filePath string) {
	file, err := os.Open(filePath)
	if err != nil {
		fmt.Println("os open file error: ", err)
		return
	}
	defer file.Close()

	buf := make([]byte, 4096)

	for {
		//	读文件，写入到socket中
		n, err := file.Read(buf)
		if err != nil {
			if err == io.EOF {
				fmt.Println("发送完毕")
			} else {
				fmt.Println("file read error: ", err)
				return
			}
		}

		_, err = sender.Write(buf[:n])
	}
}

func DoSend() {
	path := ""
	//获取文件属性
	fileInfo, err := os.Stat(path)
	if err != nil {
		fmt.Println("os stat error: ", err)
		return
	}

	//	主动发起连接
	conn, err := net.Dial("tcp", "loacalhost:8000")
	if err != nil {
		fmt.Println("net dial error: ", err)
		return
	}
	sender := NewSender(conn)

	//	发送文件名给接收端
	_, err = sender.Write([]byte(fileInfo.Name()))

	//	读取接收端发回的信息
	buf := make([]byte, 4096)
	n, err := sender.Read(buf)
	if err != nil {
		fmt.Println("conn read error: ", err)
		return
	}

	if string(buf[:n]) == "ok" {
		sender.SendFile(path)
	}
}
