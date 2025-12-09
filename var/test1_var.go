package main
// 四种变量的声明方式
import "fmt"

// 声明全局变量，方法一、二、三是可以的
var gA int =100
var gB =200

// 用方法四来声明全局变量
// := 只能够用在  函数体来声明
// gc :=200 会编译出错

func main() {
	// 方法一：声明一个变量,默认的值为0
	var a int   
	fmt.Println("a = ",a)
	fmt.Printf("Type of a = %T\n", a)      //打印变量的类型 %T

	// 方法二：声明一个变量,初始值设置为100
	var b int = 100
	fmt.Println("b = ",b)
	fmt.Printf("Type of b = %T\n", b)

	// 方法三：在初始化的时候，可以省去数据类型，通过值自动匹配当前的变量的数据类型(不推荐)
	var c = 100
	fmt.Println("c = ",c)
	fmt.Printf("Type of c = %T\n", c)

	// 方法四： （常用的方法） 省去var关键字，直接自动匹配
	e :=100
	fmt.Println("e = ",e)
	fmt.Printf("Type of e = %T\n",e)

	f := "abcd"
	fmt.Println("f = ",f)
	fmt.Printf("Type of f = %T\n",f)

	g := 3.14
	fmt.Println("g = ",g)
	fmt.Printf("Type of g = %T\n",g)

	// ----
	fmt.Println("gA = ",gA,", gB = ",gB)


	// 声明多个变量
	var xx, yy int = 100, 200
	fmt.Println("xx = ",xx,", yy = ",yy)
	var kk, ll = 100, "asdtf"
	fmt.Println("kk = ",kk,", ll = ",ll)
	
	// 用多行声明变量
	var (
		vv int = 100
		jj bool = true
	)
	fmt.Println("vv = ",vv,", jj = ",jj)
}
