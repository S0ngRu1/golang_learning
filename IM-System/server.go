package main

import (
	"bufio"
	"fmt"
	"io"
	"net"
	"strings"
	"sync"
	"time"
)

type Server struct {
	Ip   string
	Port int
	// 在线用户的列表
	OnlineMap map[string]*User
	mapLock   sync.RWMutex

	// 消息广播的channel
	Message chan Message
}

type Message struct {
	SenderAddr string
	Content    string
	Timestamp  time.Time
}

// 创建一个server的接口
func NewServer(ip string, port int) *Server {
	server := &Server{
		Ip:        ip,
		Port:      port,
		OnlineMap: make(map[string]*User),
		Message:   make(chan Message),
	}
	return server
}

// 监听Message广播消息channel的goroutine 一旦有消息就发送给全部的在线user
func (s *Server) ListenMessage() {
	for msg := range s.Message {
		formatted := fmt.Sprintf("[%s] %s", msg.SenderAddr, msg.Content)

		s.mapLock.Lock()
		for _, cli := range s.OnlineMap {
			if cli.Addr == msg.SenderAddr {
				continue // 跳过自己
			}
			cli.C <- formatted // 或直接发 Message 结构体
		}
		s.mapLock.Unlock()
	}
}

// 广播消息的方法
func (s *Server) BroadCast(senderAddr, content string) {
	msg := Message{
		SenderAddr: senderAddr,
		Content:    content,
		Timestamp:  time.Now(),
	}
	s.Message <- msg
}

// 当前的链接业务
func (s *Server) Handler(conn net.Conn) {
	user := NewUser(conn, s)
	user.Online()
	isLive := make(chan bool)
	// 接收客户端发送的消息
	go func() {
    reader := bufio.NewReader(conn)
    for {
        msg, err := reader.ReadString('\n')
        if err != nil {
            if err != io.EOF {
                fmt.Println("Read error:", err)
            }
            user.Offline()
            return
        }
			// 去除 \r\n 或 \n
			msg = strings.TrimSpace(msg)
			// 用户将消息进行广播
			user.DoMessage(msg)
			// 用户的任意消息，代表当前用户是活跃的
			isLive <- true
		}
	}()

	// 当前handler阻塞
	for {
		select{
		case <- isLive:
			// 当前用户是活跃的，应该重置定时器
		case <- time.After(time.Second * 300):
			// 已经超时
			// 将当前的User 强制的关闭
			user.SendMsg("you are out!")
			// 销毁资源
			close(user.C)
			// 关闭连接
			conn.Close()

			// 退出当前的Handler
			return
		}
		
	}
}

// 启动服务器的接口
func (s *Server) Start() {
	// socket listen
	listener, err := net.Listen("tcp", fmt.Sprintf("%s:%d", s.Ip, s.Port))
	if err != nil {
		fmt.Println("net.Listen err: ", err)
		return
	}
	// 启动监听Message的goroutine
	go s.ListenMessage()

	// clsoe listen socket
	defer listener.Close()

	for {
		// accept
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("listener accept err: ", err)
			continue
		}
		// do handler
		go s.Handler(conn)
	}
}
