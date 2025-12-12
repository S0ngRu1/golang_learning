package main

import "fmt"

func printArray(myAarry [10]int) {
	// 值拷贝
	for index, value := range myAarry {
		fmt.Println("index = ", index, ", value = ", value)
	}

}

func main() {

	//固定长度的数组
	var myAarry1 [10]int

	myAarry2 := [10]int{1, 2, 3, 4}
	for i := 0; i < len(myAarry1); i++ {
		fmt.Println(myAarry1[i])
	}

	for index, value := range myAarry2 {
		fmt.Println("index = ", index, ", value = ", value)
	}

	// 查看数组的数据类型
	fmt.Printf("myArray1 types = %T\n", myAarry1)
	fmt.Printf("myArray2 types = %T\n", myAarry2)

	printArray(myAarry2)
}
