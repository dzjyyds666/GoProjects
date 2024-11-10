package main

import (
	"fmt"
	"io"
	"net"
	"sync"
)

type Server struct {
	IP   string
	port int

	//在线用户map
	OnlineMap map[string]*User
	maplock   sync.RWMutex

	//消息广播饿channel
	Message chan string
}

// 创建一个server接口
func NewServer(ip string, port int) *Server {
	server := &Server{
		IP:        ip,
		port:      port,
		OnlineMap: make(map[string]*User),
		Message:   make(chan string),
	}

	return server
}

func (this *Server) ListenMessage() {
	for {
		msg := <-this.Message
		//将msg发送给在线的全部用户
		this.maplock.Lock()
		for _, cli := range this.OnlineMap {
			cli.C <- msg
		}
		this.maplock.Unlock()
	}
}

// 把用户上线的消息送入管道
func (this *Server) broadCast(user *User, msg string) {
	sendMsg := "[" + user.Addr + "]" + user.Name + ":" + msg
	this.Message <- sendMsg
}

func (this *Server) Handler(conn net.Conn) {
	//...当前链接业务
	//fmt.Println("链接建立成功")
	user := NewUser(conn, this)

	user.Online()

	//接收客户端消息（去除‘\n’）
	go func() {
		buf := make([]byte, 4096)
		for {
			n, err := conn.Read(buf)
			if n == 0 {
				user.Offline()
				return
			} else if err != nil && err != io.EOF {
				fmt.Println("conn.Read err:", err)
				return
			} else {
				// 提取用户消息，去除最后的'\n'
				msg := string(buf[:n-1])

				user.Domessage(msg)
			}

		}
	}()

	//当前Handler阻塞
	select {}
}

// 启动服务器的接口
func (this *Server) start() {
	//socket监听
	listener, err := net.Listen("tcp", fmt.Sprintf("%s:%d", this.IP, this.port))
	// 监听错误
	if err != nil {
		fmt.Println("net listen err:", err)
		return
	}

	//关闭sokcet监听
	defer listener.Close()

	//启动监听Message的方法
	go this.ListenMessage()

	for {
		//accpet
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("listener accpect err:", err)
			continue
		}

		//do Handler
		go this.Handler(conn)
	}
}
