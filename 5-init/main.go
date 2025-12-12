package main

import (
	_ "5-init/lib1"   //包的匿名，可以避免不使用报错
	mylib2 "5-init/lib2"  //为包起别名mylib2
	// . "5-init/lib2"  //可直接调用该包的方法

)

func main(){
	// lib1.Lib1Test()
	mylib2.Lib2Test()
	// Lib2Test()
}