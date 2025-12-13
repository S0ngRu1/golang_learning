package main

import "fmt"

type Hero struct {
	Name  string
	Ad    int
	Level int
}

func (this Hero) Show() {
	// 这个this是调用方法的对象的一个副本（值传递）
	fmt.Println("Name = ", this.Name)
	fmt.Println("Ad = ", this.Ad)
	fmt.Println("Level = ", this.Level)
}

func main() {

}
