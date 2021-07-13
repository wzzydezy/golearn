package main

import (
	"net"
	"strings"
)

type User struct {
	Name string
	Addr string
	C    chan string
	conn net.Conn

	server *Server
}

func NewUser(conn net.Conn, server *Server) *User {
	userAddr := conn.RemoteAddr().String()

	user := &User{
		Name: userAddr,
		Addr: userAddr,
		C:    make(chan string),
		conn: conn,

		server: server,
	}

	go user.ListenMessage()

	return user
}

//上线业务
func (this *User) Online() {
	//用户上线，将用户加入到表中
	this.server.maplock.Lock()
	this.server.OnlineMap[this.Name] = this
	this.server.maplock.Unlock()

	//广播当前用户上线消息
	this.server.BroadCast(this, "已上线")
}

//用户下线业务
func (this *User) OffOnline() {
	//用户下线，将用户从表中删除
	this.server.maplock.Lock()
	delete(this.server.OnlineMap, this.Name)
	this.server.maplock.Unlock()
	//广播当前用户下线消息
	this.server.BroadCast(this, "已下线")
}

func (this *User) SendMsg(msg string) {
	this.conn.Write([]byte(msg))
}

//处理消息业务
func (this *User) DoMessage(msg string) {
	if msg == "who" {
		//查询用户有哪些

		this.server.maplock.Lock()
		for _, user := range this.server.OnlineMap {
			onlineMsg := "[" + user.Addr + "]" + user.Name + ":" + "在线\n"
			this.SendMsg(onlineMsg)
		}
		this.server.maplock.Unlock()

	} else if len(msg) > 7 && msg[:7] == "rename|" {
		//消息格式 rename|张三
		newName := strings.Split(msg, "|")[1]
		//判断name是否存在
		_, ok := this.server.OnlineMap[newName]
		if ok {
			this.SendMsg("name have been used\n")
		} else {
			this.server.maplock.Lock()
			delete(this.server.OnlineMap, this.Name)
			this.server.OnlineMap[newName] = this
			this.server.maplock.Unlock()

			this.Name = newName
			this.SendMsg("your newName are updated:" + this.Name + "\n")
		}

	} else if len(msg) > 4 && msg[:3] == "to|" {
		//消息格式 to|who|msg
		//1 获取用户名
		remoteName := strings.Split(msg, "|")[1]
		if remoteName == "" {
			this.SendMsg("please use to|who|msg\n")
			return
		}

		remoteUser, ok := this.server.OnlineMap[remoteName]
		if !ok {
			this.SendMsg("user no exist\n")
			return
		}

		content := strings.Split(msg, "|")[2]
		if content == "" {
			this.SendMsg("msg no exist\n")
			return
		}
		remoteUser.SendMsg(this.Name + " say to you:" + content + "\n")

		//2 根据用户名 获取User

		//3 获取消息内容，发送消息

	} else {
		this.server.BroadCast(this, msg)
	}
}

//监听channel的方法，有消息就要发送给客户端

func (this *User) ListenMessage() {
	for {
		msg := <-this.C
		//fmt.Println("客户端", msg)
		this.conn.Write([]byte(msg + "\n"))
	}
}
