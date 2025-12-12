package main

import "fmt"

func printArray(myAarry []int) {
	//slice传递是一种引用传递
	for _, value := range myAarry {
		fmt.Println("value = ", value)
	}
	myAarry[0] = 100
}

func main() {

	myAarry5 := []int{1, 2, 3, 4} //动态数组，切片 slice
	fmt.Printf("myArray type is %T\n", myAarry5)

	printArray(myAarry5)

	for _, value := range myAarry5 {
		fmt.Println("value = ", value)
	}
}
