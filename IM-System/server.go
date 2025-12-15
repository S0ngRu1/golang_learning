package main

import (
	"net"
	"fmt"
	"bufio"
)

type Server struct {
	Ip   string
	Port int
}

// 创建一个server的接口
func NewServer(ip string, port int) *Server {
	server := &Server{
		Ip:   ip,
		Port: port,
	}
	return server
}

// 当前的链接业务
func (s *Server) Handler(conn net.Conn) {
    defer conn.Close()
    reader := bufio.NewReader(conn)
    msg, _ := reader.ReadString('\n')
    fmt.Fprintf(conn, "Got your message: %s", msg) // 必须写回！
}



// 启动服务器的接口
func (s *Server) Start(){
	// socket listen
	listener, err := net.Listen("tcp",fmt.Sprintf("%s:%d",s.Ip,s.Port))
	if err != nil{
		fmt.Println("net.Listen err: ",err)
		return
	}

	// clsoe listen socket
	defer listener.Close()

	for {
		// accept
		conn, err := listener.Accept()
		if err != nil{
			fmt.Println("listener accept err: ",err)
			continue
		}
		// do handler
		go s.Handler(conn)
	}
}