package main

import (
	"net"
	"strings"
)

type User struct {
	Name   string
	Addr   string
	C      chan string
	conn   net.Conn
	server *Server
}

// 创建一个用户的API
func NewUser(conn net.Conn, server *Server) *User {
	userAddr := conn.RemoteAddr().String()
	user := &User{
		Name:   userAddr,
		Addr:   userAddr,
		C:      make(chan string),
		conn:   conn,
		server: server,
	}

	// 启动监听当前user channel消息的goroutine
	go user.ListenMessage()
	return user
}

// 用户上线的业务
func (u *User) Online() {
	// 用户上线，将用户加入到OnlineMap中
	u.server.mapLock.Lock()
	u.server.OnlineMap[u.Name] = u
	u.server.mapLock.Unlock()

	// 广播当前用户的上线消息
	u.server.BroadCast(u.Addr, u.Name+": is online \n")
}

// 用户下线的业务
func (u *User) Offline() {
	// 用户下线，将用户从OnlineMap中删除
	u.server.mapLock.Lock()
	delete(u.server.OnlineMap, u.Name)
	u.server.mapLock.Unlock()

	// 广播当前用户的上线消息
	u.server.BroadCast(u.Addr, u.Name+": is offline \n")
}

func (u *User) SendMsg(msg string) {
	u.conn.Write([]byte(msg))
}

// 用户处理消息的业务
func (u *User) DoMessage(msg string) {
	if msg == "who" {
		// 查询当前在线用户都有哪些
		u.server.mapLock.Lock()
		for _, user := range u.server.OnlineMap {
			onlineMsg := "[" + user.Addr + "]" + user.Name + ":" + "is online...\n"
			u.SendMsg(onlineMsg)
		}
		u.server.mapLock.Unlock()

	} else if len(msg) > 7 && msg[:7] == "rename|" {
		// 消息格式: rename|zhangsan
		newName := strings.Split(msg, "|")[1]
		// 判断name是否存在
		_, ok := u.server.OnlineMap[newName]
		if ok {
			u.SendMsg("The current username is in use\n")
		} else {
			u.server.mapLock.Lock()
			delete(u.server.OnlineMap, u.Name)
			u.server.OnlineMap[newName] = u
			u.server.mapLock.Unlock()
			u.Name = newName
			u.SendMsg("You have updated your username:" + u.Name + "\n")
		}

	} else if len(msg) > 4 && msg[:3] == "to|"{
		// 消息格式: to|张三|消息内容
		// 1 获取对方的用户名
		remoteName := strings.Split(msg,"|")[1]
		if remoteName == ""{
			u.SendMsg("消息格式不正确,请使用\"to|张三|消息内容\"格式.\n")
			return
		}

		// 2 根据用户名 得到对方的User对象
		remoteUser,ok := u.server.OnlineMap[remoteName]
		if !ok {
			u.SendMsg("该用户名不存在\n")
			return
		}
		// 3 获取消息内容,通过对方的User对象将消息内容发送过去
		content := strings.Split(msg,"|")[2]
		if content == ""{
			u.SendMsg("无消息内容,请重发\n")
			return
		}
		remoteUser.SendMsg(u.Name + "to o tell you:"+content)

	}else {
		u.server.BroadCast(u.Addr, msg)
	}
}

// 监听当前User  channel 的方法，一旦有消息，就直接发送客户端
func (u *User) ListenMessage() {
	for {
		msg := <-u.C
		u.conn.Write([]byte(msg + "\n"))
	}
}
