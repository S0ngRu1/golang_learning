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


	c3 := make(chan int)
	go func(){
		for i := 0; i<5; i++{
			c3 <- i
		}
		// close 可以关闭一个channel
		close(c3)
	}()

	// for {
	// 	// ok 如果为true表示channel没有关闭，如果为false表示channel已经关闭
	// 	if data, ok := <-c3; ok{
	// 		fmt.Println(data)
	// 	} else {
	// 		break
	// 	}
	// }

	



	// 可以使用range来迭代不断操作channel
	for data:= range c3{
		fmt.Println(data)
	}

	fmt.Println("Main Finished.")


	// 使用select

	c4 := make(chan int)
	quit := make(chan int)

	go func ()  {
		for i:=0;i<6; i++{
			fmt.Println(<-c4)
		}

		quit <- 0
	}()

	fibonacci(c4,quit)
	 
}

func fibonacci(c, quit chan int) {
    x, y := 0, 1  // 从 F(0)=0, F(1)=1 开始
    for {
        select {
        case c <- x:       // 先发送当前值 x
            x, y = y, x+y  // 再更新：新 x = 旧 y，新 y = 旧 x + 旧 y
        case <-quit:
            fmt.Println("quit")
            return
        }
    }
}