package main

import (
	"fmt"
	"net"
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
	// defer conn.Close()
	// reader := bufio.NewReader(conn)
	// msg, _ := reader.ReadString('\n')
	// fmt.Fprintf(conn, "Got your message: %s", msg) // 必须写回！
	user := NewUser(conn)

	// 用户上线，将用户加入到OnlineMap中
	s.mapLock.Lock()
	s.OnlineMap[user.Name] = user
	s.mapLock.Unlock()

	// 广播当前用户的上线消息
	s.BroadCast(user.Addr, user.Name+": is online \n")

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
