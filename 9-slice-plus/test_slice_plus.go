package main

import "fmt"

func main(){
	// 定义一个切片，并分配一个容器为5的空间，前三个值的初始值为0
	var numbers = make([]int, 3, 5)

	fmt.Printf("len = %d, cap = %d, slice = %v\n",len(numbers),cap(numbers),numbers)
	// 向numbers切片追加一个元素1， numbers_len = 4, [0,0,0,1], cap = 5
	numbers = append(numbers, 1)
	fmt.Printf("len = %d, cap = %d, slice = %v\n",len(numbers),cap(numbers),numbers)


	numbers = append(numbers, 2)
	fmt.Printf("len = %d, cap = %d, slice = %v\n",len(numbers),cap(numbers),numbers)
	// 当再进行追加时，cap会翻倍，另外开辟一个cap的空间，但没有初始值
	numbers = append(numbers, 3)
	fmt.Printf("len = %d, cap = %d, slice = %v\n",len(numbers),cap(numbers),numbers)

	fmt.Println("---------------------")
	var numbers2 = make([]int, 3)
	fmt.Printf("len = %d, cap = %d, slice = %v\n",len(numbers2),cap(numbers2),numbers2)
	numbers2 = append(numbers2, 1)
	fmt.Printf("len = %d, cap = %d, slice = %v\n",len(numbers2),cap(numbers2),numbers2)

	s := numbers[:2]  // s其实是指向numbers的一个指针
	fmt.Println("s = ", s)
	s[1] = 100      //修改后，numbers会发生改变
	fmt.Println("s = ", s)
	fmt.Println("numbers = ", numbers)

	// copy 可以将底层数组slice一起进行拷贝
	s2 := make([]int,3)    // s2 = [0,0,0]
	copy(s2,s)
	fmt.Println("s2 = ", s2)
	s[0] = 100
	fmt.Println("s2 = ", s2)
}	