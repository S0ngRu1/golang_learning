package main

import "fmt"

// 返回单个值
func fool(a string, b int) int {
	fmt.Println("a = ",a)
	fmt.Println("b = ",b)

	c := 100
	return	c
}

// 返回多个返回值，匿名
func fool2(a string, b int) (int, int) {
	fmt.Println("a = ",a)
	fmt.Println("b = ",b)

	return	666, 777
}

// 返回多个返回值，有形参名称的
func fool3(a string, b int) (r1 int, r2 int) {
	fmt.Println("a = ",a)
	fmt.Println("b = ",b)
	// r1 ,r2 属于fool3的形参，默认值为0
	fmt.Println("r1 = ",r1)
	fmt.Println("r2 = ",r2)


	// 给有名称的返回值变量赋值
	r1 = 1000
	r2 = 2000
	return	
}

// 返回多个返回值，有形参名称的
func fool4(a string, b int) (r1, r2 int) {
	fmt.Println("a = ",a)
	fmt.Println("b = ",b)

	// 给有名称的返回值变量赋值
	r1 = 1000
	r2 = 2000
	return	
}


func main() {
	c := fool("abc",555)
	fmt.Println("c = ",c)
	
	ret1, ret2 := fool2("haha", 999)
	fmt.Println("ret1 = ", ret1,", ret2 = ",ret2 )
	
	ret1, ret2 = fool3("fool3", 333)
	fmt.Println("ret1 = ", ret1,", ret2 = ",ret2 )

	ret1, ret2 = fool4("fool4", 444)
	fmt.Println("ret1 = ", ret1,", ret2 = ",ret2 )
}
