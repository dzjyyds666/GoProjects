package main

import (
	"fmt"
	"net"
	"strings"
)

type User struct {
	Name string
	Addr string
	C    chan string
	conn net.Conn

	//当前用户所属server
	server *Server
}

// 创建user的接口
func NewUser(conn net.Conn, server *Server) *User {
	userAddr := conn.RemoteAddr().String()
	user := &User{
		Name: userAddr,
		Addr: userAddr,
		C:    make(chan string),
		conn: conn,

		server: server,
	}

	//创建用户后即监听消息
	go user.ListenMessage()

	return user
}

// 用户上线
func (this *User) Online() {
	//用户上线，将用户加入到onlineMap中
	this.server.maplock.Lock()
	this.server.OnlineMap[this.Name] = this
	this.server.maplock.Unlock()

	// 广播当前用户上线消息
	this.server.broadCast(this, "上线")
}

// 用户下线
func (this *User) Offline() {
	//删除用户
	this.server.maplock.Lock()
	delete(this.server.OnlineMap, this.Name)
	this.server.maplock.Unlock()

	//广播当前用户下线消息
	this.server.broadCast(this, "下线")
}

func (this *User) SendMsg(msg string) {
	this.conn.Write([]byte(msg))
}

// 用户处理消息的业务
func (this *User) Domessage(msg string) {
	if msg == "who" {
		//查询当前在线用户
		this.server.maplock.Lock()
		for _, user := range this.server.OnlineMap {
			onlineMsg := "[" + user.Addr + "]" + user.Name + ":" + "在线...\n"
			this.SendMsg(onlineMsg)
		}
		this.server.maplock.Unlock()
	} else if strings.Contains(msg, "rename|") {
		this.SendMsg("修改名字")
		//改名字
		this.server.maplock.Lock()
		for _, user := range this.server.OnlineMap {
			if user.Name == this.Name {
				user.Name = strings.Split(msg, "|")[1]

			}
			break
		}
		fmt.Println(this.server.OnlineMap)
		this.server.maplock.Unlock()
	} else {
		this.server.broadCast(this, msg)
	}

}

// 监听当前User channel的方法，一旦有消息，就直接发送给客户端
func (this *User) ListenMessage() {
	for {
		msg := <-this.C
		this.conn.Write([]byte(msg + "\n"))
	}
}
