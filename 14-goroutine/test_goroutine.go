package main

import (
	"fmt"
	"sync"
	"time"
)

func newTask() {
	i := 0
	for {
		i++
		fmt.Println("new Goroutine : i = ",i)
		time.Sleep(1 *time.Second)
	}
}
func main() {
	// 创建一个go程，去执行newTask() 流程
	// go newTask()
	// fmt.Println("main goroutine exit")
	
	// i := 0
	// for {
	// 	i++
	// 	fmt.Println("main Goroutine : i = ",i)
	// 	time.Sleep(1 *time.Second)
	// }


	// 用go创建承载一个形参为空，返回值为空的一个函数
	// go func(){
	// 	defer fmt.Println("A.defer")
	// 	func ()  {
	// 		defer fmt.Println("B.defer")
	// 退出当前的go程
	// runtime.Goexit()    //终止当前的goroutine
	// 		fmt.Println("B")
	// 	}()    //调用当前函数
	// 	fmt.Println("A")
	// }()     //调用当前函数

	// go func (a int, b int ) bool{
	// 	fmt.Println("a = ",a , "b = ",b)
	// 	return true
	// }(10,20)
// // 死循环
// for {
// 	time.Sleep(1 *time.Second)
// }


	var wg sync.WaitGroup
	wg.Add(1)
	go func(){
		defer wg.Done()
			defer fmt.Println("A.defer")
			func ()  {
				defer fmt.Println("B.defer")
		// 退出当前的go程
		// runtime.Goexit()    //终止当前的goroutine
				fmt.Println("B")
			}()    //调用当前函数
			fmt.Println("A")
		}()     //调用当前函数
	wg.Wait()  // 等待goroutine完成
	fmt.Println("main exit")
	
}