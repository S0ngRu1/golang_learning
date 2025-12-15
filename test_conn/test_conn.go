package main

import (
    "bufio"
    "fmt"
    "net"
)

func main() {
    // 向ip发起TCP链接，如果链接失败err不为nil
    conn, err := net.Dial("tcp", "127.0.0.1:8888")
    if err != nil {
        panic(err)
    }
    defer conn.Close()

    // 发送数据
    fmt.Fprintln(conn, "Hello from client!")

    // 读取响应
    // 包装 conn 为带缓冲的 Reader，提升 I/O 效率。
    reader := bufio.NewReader(conn)
    // bufio.Reader 支持按行读取（ReadString('\n')）、按字节读等。
    response, _ := reader.ReadString('\n')
    fmt.Print("Server says: ", response)
}