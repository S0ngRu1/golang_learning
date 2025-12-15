package main

import (
	"fmt"
	"time"
)

func main() {

	// 定义一个无缓存channel
	c := make(chan int)
	go func(){
		defer fmt.Println("goroutine 结束")
		fmt.Println("goroutine 正在运行...")
		c <- 666   // 将666 发送给c
	}()

	num := <- c
	fmt.Println("num = ", num)
	fmt.Println("main goroutine 结束...")	


	// 定义一个有缓存的channel
	c2 := make(chan int ,3) // 带有三个缓存
	fmt.Println("len(c2) = ",len(c2), ", cap(c) = ",cap(c2))
	
	go func(){
		defer fmt.Println("子go 程结束")
		for i := 0;i<5;i++{
			c2 <- i
			fmt.Println("子go程正在运行，发送的元素=",i,"len(c2) = ",len(c2),",cap(c2) = ",cap(c2))
		}
	}()

	time.Sleep(2 *time.Second)
	for i := 0;i < 5 ;i++{
		num := <-c2
		fmt.Println("num = ", num)
	}
	time.Sleep(2 *time.Second)
	fmt.Print("main 结束")
}