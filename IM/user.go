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
		Name:   userAddr,
		Addr:   userAddr,
		C:      make(chan string),
		conn:   conn,
		server: server,
	}
	go user.ListenMessage()
	return user
}

func (this *User) ListenMessage() {
	for {
		msg := <-this.C
		this.conn.Write([]byte(msg + "\r\n"))
	}
}

// 用户上线的业务
func (this *User) Online() {
	// 用户上线，将用户加入到OnlineMap中
	this.server.mapLock.Lock()
	this.server.OnlineMap[this.Name] = this
	this.server.mapLock.Unlock()

	// 广播当前用户上线消息
	this.server.BroadCast(this, "已上线")
}

// 用户下线的业务
func (this *User) Offline() {

	// 用户上线，将用户加入到OnlineMap中
	this.server.mapLock.Lock()
	delete(this.server.OnlineMap, this.Name)
	this.server.mapLock.Unlock()

	// 广播当前用户上线消息
	this.server.BroadCast(this, "已下线")
}

func (this *User) SendMsg(msg string) {
	msg = msg + "\r\n"
	this.conn.Write([]byte(msg))
}

// 用户处理消息的业务
func (this *User) DoMessage(msg string) {
	if msg == "who" { // 处理查询在线用户，消息格式：who
		this.server.mapLock.Lock()
		for _, user := range this.server.OnlineMap {
			onlineMap := "[" + user.Addr + "]" + user.Name + "：" + "在线..."
			this.SendMsg(onlineMap)
		}
		this.server.mapLock.Unlock()
	} else if len(msg) > 7 && msg[:7] == "rename|" { // 处理重命名，消息格式：rename|新名字
		newName := strings.Split(msg, "|")[1]

		// 判断name是否存在
		_, ok := this.server.OnlineMap[newName]
		if ok {
			this.SendMsg("当前用户名被使用")
		} else {
			this.server.mapLock.Lock()
			delete(this.server.OnlineMap, this.Name)
			this.Name = newName
			this.server.OnlineMap[this.Name] = this
			this.server.mapLock.Unlock()
			this.SendMsg("更新成功")
		}
	} else if len(msg) > 3 && msg[:3] == "to|" { // 处理私聊，消息格式：to|user.Name|消息
		remoteName := strings.Split(msg, "|")[1]
		msg := strings.Split(msg, "|")[2]
		if remoteName == "" {
			this.SendMsg("消息格式不正确")
		}
		if msg == "" {
			this.SendMsg("无消息内容")
		}

		remoteUser, ok := this.server.OnlineMap[remoteName]
		if !ok {
			this.SendMsg("用户不存在")
		}
		remoteUser.SendMsg(this.Name + "对你说：" + msg)

	} else {
		this.server.BroadCast(this, msg)
	}

}
