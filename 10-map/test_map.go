package main

import "fmt"

func map_declare(){
	// 声明myMap是一种map类型，key是string ，value 是string
	var myMap map[string]string
	if myMap == nil {
		fmt.Println("myMap是一个空的map")
	}
	// 在使用map前， 需要先用make给map分配数据空间
	myMap = make(map[string]string, 10)
	myMap["one"] = "java"
	myMap["two"] = "go"
	myMap["three"] = "c++"
	
	fmt.Println(myMap)
	// 第二种声明，直接开辟空间
	myMap2 := make(map[int]string, 10)
	myMap2[1] = "java"
	myMap2[2] = "go"
	myMap2[3] = "c++"
	
	fmt.Println(myMap2)

	// 第三种声明，直接设定值

	myMap3 := map[string]string{
		"1":"java",
		"2":"java",
		"3":"java",
	}
	myMap3["1"] = "java"
	myMap3["2"] = "go"
	myMap3["3"] = "c++"
	
	fmt.Println(myMap3)

}

func printMap(myMap map[string]string){
	// myMap 是一个引用传递
	for key, value := range myMap {
		fmt.Println("key = ", key)
		fmt.Println("value = ", value)
	}
}

func changeValue(myMap map[string]string){
	myMap["England"] = "London"
}

func map_use(){
	citeMape := make(map[string]string)

	// 添加
	citeMape["China"] = "Beijing"
	citeMape["Japan"] = "Tokyo"
	citeMape["USA"] = "NewYork"
	
	// 遍历
	printMap(citeMape)

	// 删除
	delete(citeMape, "China")
	fmt.Println("----------------")

	// 修改
	citeMape["USA"] = "DC"
	changeValue(citeMape)
	// 遍历
	printMap(citeMape)
}

func main(){
	map_use()
}