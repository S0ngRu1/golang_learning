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

type Human struct{
	Name string
	Sex string
}

func (this *Human) Eat(){
	fmt.Println("Human.Eat()...")
}

func (this *Human) Walk(){
	fmt.Println("Human.Walk()...")
}



// 继承
type SuperMan struct{
	Human  // SuperMan 类继承了Human类的方法

	level int
}
func (this *SuperMan) Eat(){
	fmt.Println("SuperMan.Eat()")
}

func (this *SuperMan) Fly(){
	fmt.Println("SuperMan.Fly()")
}



func test_class() {
	// // 创建一个对象
	// hero := Hero{Name: "zhang3",Ad: 100, Level: 1}
	// hero.Show()
	// hero.GetName()
	// hero.SetName("李四")
	// hero.Show()

	zhangsan := Human{"zhangsan", "female"}

	zhangsan.Eat()
	zhangsan.Walk()

	// 定义一个子类对象
	// 子类名{父类名{父类属性}，子类属性}
	s := SuperMan{Human{"li","female"},88}
	s.Walk()
	s.Eat()
	s.Fly()

}
