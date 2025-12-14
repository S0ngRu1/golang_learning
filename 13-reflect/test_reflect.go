package main

import (
	"fmt"
	"reflect"
)

type User struct {
	Id   int
	Name string
	Age  int
}

func (this *User) Call() {
	fmt.Println("user is called ..")
	fmt.Printf("%v \n", this)
}

func reflectNum(arg interface{}) {
	fmt.Println("type : ", reflect.TypeOf(arg))
	fmt.Println("value : ", reflect.ValueOf(arg))
}

func main() {
	var num float64 = 1.234
	reflectNum(num)

	fmt.Println("=---------------------------------=")
	user := &User{1, "abc", 18}
	DoFiledAndMethod(user)
}

func DoFiledAndMethod(input interface{}) {
	
	t := reflect.TypeOf(input)
	v := reflect.ValueOf(input)
	// 如果是指针，获取它指向的类型和值
	if t.Kind() == reflect.Ptr{
		t = t.Elem()   // 获取 *User -> User
		v = v.Elem()   // 获取 *User 的值 -> User 的值
	}

	// 现在 t 是 User（结构体类型），可以安全遍历字段
    fmt.Println("Struct type:", t.Name())
	// 通过type 获取里面的字段
	// 1.获取interface的reflaect.Type,通过Type得到NumField，进行遍历
	// 2.获得每一个field的数据类型
	// 3. 通过field的Interface()方法得到对应的value
	    for i := 0; i < t.NumField(); i++ {
        	field := t.Field(i)
        	value := v.Field(i).Interface()
        fmt.Printf("%s: %v = %v\n", field.Name, field.Type, value)
    }

	// 方法：方法属于原始类型 *User
    // 所以我们用原始的 Type（未 Elem 的）来获取方法
    origType := reflect.TypeOf(input)
    fmt.Printf("\nMethods (%d):\n", origType.NumMethod())
	// 通过type获取里面的方法，调用
	for i := 0; i < origType.NumMethod(); i++ {
        m := origType.Method(i)
        fmt.Printf("- %s: %v\n", m.Name, m.Type)
    }
}
