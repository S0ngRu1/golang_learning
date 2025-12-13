package main

import "fmt"

// 如果类名首字母大写，表示其他包也能够访问
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



func (this Hero) GetName() string {
	// 这个this是调用方法的对象的一个副本（值传递）
	return this.Name
}

func (this *Hero) SetName(newName string){
	// 这个this是调用方法对象的地址
	this.Name = newName 
} 
func main() {
	// 创建一个对象
	hero := Hero{Name: "zhang3",Ad: 100, Level: 1}
	hero.Show()
	hero.GetName()
	hero.SetName("李四")
	hero.Show()
}
