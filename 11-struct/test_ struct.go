package main

import "fmt"

// 声明一种新的数据类型   myint ,是int的一个别名

type myint int


// 定义一个结构体
type Book struct{
	title string
	auth  string
}


func ChangeBook(book Book){
	// 传递的是book的副本
	book.auth = "aaaa"
}

func ChangeBook2(book *Book){
	// 传递的是book的指针
	book.auth = "bbb"
}

func main(){
	var a myint = 10
	fmt.Println("a = ",a)
	fmt.Printf("type of a = %T\n", a)

	var book1 Book
	book1.title = "Golang"
	book1.auth = "zhang3"
	fmt.Printf("%v \n",book1)

	ChangeBook(book1)
	fmt.Printf("%v \n",book1)

	ChangeBook2(&book1)
	fmt.Printf("%v \n",book1)
}