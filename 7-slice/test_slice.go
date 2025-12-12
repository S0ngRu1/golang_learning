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


	slice := []int{1,2,3}
	fmt.Printf("len = %d , slice = %v\n",len(slice),slice)

	// 声明slice 1是一个切片，但是并没有给slice分配空间
	var slice1 []int
	fmt.Printf("len = %d , slice = %v\n",len(slice1),slice1)
	// 如果一个slice没有任何空间，是不能被赋值的
	// 开辟容量，分配3个空间，默认值为0
	slice1 = make([]int,3)
	slice1[1] = 3   // 赋值
	fmt.Printf("len = %d , slice = %v\n",len(slice1),slice1)

	// 声明slice是一个切片，同时给slice分配空间，三个空间，初始值为0
	var slice2 []int = make([]int, 3)
	fmt.Printf("len = %d , slice = %v\n",len(slice2),slice2)

	// 声明slice3 是一个切片，同时给slice分配空间，3个空间，初始化值为0，通过:=推导出slice是一个切片
	slice3 := make([]int ,3)
	fmt.Printf("len = %d , slice = %v\n",len(slice3),slice3)

	// 判断一个slice是否是0
	if slice1 == nil {
		fmt.Println("slice1 是一个空切片")
	} else {
		fmt.Println("slice1 是有空间的")
	}
}
