package main

import "fmt"

func swap(a int , b int){
	temp := a
	a = b
	b = temp
}

func pswap(a *int , b *int){
	temp := *a
	*a = *b
	*b = temp
}



func main(){
	var a int = 10
	b := 20
	swap(a,b)
	fmt.Println("a = ",a,", b = ",b)
	pswap(&a,&b)
	fmt.Println("a = ",a,", b = ",b)
}