package main

import (
	"fmt"
	"io"
	"net"
	"sync"
	"time"
)

type Server struct {
	Ip   string
	Port int

	//在线用户列表
	OnlineMap map[string]*User
	maplock   sync.RWMutex

	Message chan string
}

//创建一个server的接口
func NewServer(ip string, port int) *Server {
	server := &Server{
		Ip:        ip,
		Port:      port,
		OnlineMap: make(map[string]*User),
		Message:   make(chan string),
	}

	return server
}

//监听Message广播消息chan的goroutine,一旦有消息就发送给全部的User
func (this *Server) ListenMessage() {
	for {
		msg := <-this.Message
		//将msg发送给全部的在线User
		this.maplock.Lock()
		for _, cli := range this.OnlineMap {
			cli.C <- msg
		}
		this.maplock.Unlock()
	}
}

//广播消息的方法
func (this *Server) BroadCast(user *User, msg string) {
	sendMsg := "[" + user.Addr + "]" + user.Name + ":" + msg
	this.Message <- sendMsg
}

func (this *Server) Handler(conn net.Conn) {
	//逻辑处理
	fmt.Println("链接建立成功")
	//用户上线，将用户加入到表中
	user := NewUser(conn, this)
	user.Online()

	//监听用户是否活跃的channel
	isLive := make(chan bool)
	//接受客户端发送的消息
	go func() {
		buf := make([]byte, 4096)
		for {
			n, err := conn.Read(buf)
			if n == 0 {
				user.OffOnline()
				return
			}

			if err != nil && err != io.EOF {
				fmt.Println("Conn read err", err)
				return
			}
			//提取用户的消息（去除'\n'）
			msg := string(buf[:n-1])
			//广播MSG
			user.DoMessage(msg)

			//用户的任意消息，代表当前用户是活跃的
			isLive <- true
		}
	}()

	//当前handler阻塞
	for {
		select {
		case <-isLive:
			//当前用户活跃的，重置定时器
			//不做任何事情，为了激活select，更新下面定时器

		case <-time.After(time.Second * 300):
			//已经超时
			//将当前User强制关闭
			user.SendMsg("you are out!\n")
			//自行添加
			user.OffOnline()
			//销毁用户资源
			close(user.C)

			//关闭连接
			conn.Close()

			//退出当前的Handler
			return
		}
	}

}

func (this *Server) Start() {
	//socket listen
	listener, err := net.Listen("tcp", fmt.Sprintf("%s:%d", this.Ip, this.Port))
	if err != nil {
		fmt.Println("net.Listen err:", err)
		return
	}
	//close listen socket
	defer listener.Close()

	//启动监听Message的goroutine
	go this.ListenMessage()

	for {
		//accept
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("Listen accept err:", err)
			continue
		}

		///do handler

		go this.Handler(conn)
	}

}
