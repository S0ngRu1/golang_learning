package main

import (
	"flag"
	"fmt"
	"net"
)

type Client struct{
    ServerIp string
    ServerPort int
    Name string
    conn net.Conn
    flag int // 当前client的模式
}

func NewClient(serverIp string, serverPort int) *Client {
    // 创建客户端对象
    client := &Client{
        ServerIp: serverIp,
        ServerPort: serverPort,
        flag: 999,
    }

    // 连接server
    conn, err := net.Dial("tcp", fmt.Sprintf("%s:%d",serverIp,serverPort))
    if err != nil {
        fmt.Println("net.Dial error : ",err)
        return nil
    }
    client.conn = conn
    return client
}

func (c *Client) meau() bool{
    var flag int 
    fmt.Println("1.公聊模式")
    fmt.Println("2.私聊模式")
    fmt.Println("3.更新用户名")
    fmt.Println("0.退出")
    // 从用户键盘接收一个字符
    fmt.Scanln(&flag)
    if flag >= 0 && flag <=3{
        c.flag = flag
        return  true
    }else {
        fmt.Println(">>>>请输入合法范围内的数字<<<<")
        return false
    }
}



var serverIp string
var serverPort int

// .client -ip 127.0.0.1 - port 8888
func init(){

// 在main之前执行
    flag.StringVar(&serverIp, "ip", "127.0.0.1", "Configure the server IP address (default is 127.0.0.1).")
    flag.IntVar(&serverPort, "port", 8888, "Configure the server Port (default is 8888).")

}

func (c *Client) Run(){
    for c.flag !=0 {
        for c.meau() != true{

        }
        // 根据不同的模式处理不同的类型
        switch c.flag{
        case 1:
            // 公聊模式
            fmt.Println("公聊模式选择....")
            break
        case 2:
            // 私聊模式
            fmt.Println("私聊模式选择....")
            break
        case 3:
            // 更新用户名
            fmt.Println("更改用户名 模式选择....")
            break
        }
    }
}


func main() {
    // 命令行解析
    flag.Parse()

    client := NewClient(serverIp,serverPort)
    if client == nil {
        fmt.Println(">>>>>>连接服务器失败...")
        return
    }else{
        fmt.Println(">>>>>>连接服务器成功...")
    }
    // 启动客户端的业务
    client.Run()
}