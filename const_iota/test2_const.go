package main

import "fmt"

// const 来定义枚举类型
const (
	// iota 只能够配合const()一起使用，iota只有在const进行累加效果
	// 可以在const()添加一个关键字iota ,每一行的iota都会累加1，第一行的iota的默认值为0
	 BEIJING = iota    // iota = 0    还可以 iota*10
	 SHANGHAI
	 SHENZHEN
	 TAIJIN
)

func main() {
	// 常量（只读属性）
	const length int =10 
	fmt.Println("length = ", length)
	// length = 100  常量不允许修改
	fmt.Println("BEIJING = ", BEIJING)
	fmt.Println("SHANGHAI = ", SHANGHAI)	
	fmt.Println("SHENZHEN = ", SHENZHEN)
}