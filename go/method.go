package main

import "fmt"

type Age uint

// 值类型接收者
func (age Age) String(){
	fmt.Println("the age is",age)
}

// Modify 指针类型接收者
func (age *Age) Modify(){
	*age = Age(30)
}

func main() {
	// 定义类型并初始化值
	age:=Age(25)

	age.String()
	age.Modify()
	age.String()

	// 函数指针(将方法作为表达式赋值给一个变量)
	sm := Age.String
	// 原始string()函数没有参数,通过方法表达式调用,第一个参数必须是接收者
	// 然后才是方法自身的参数
	//通过变量，要传一个接收者进行调用也就是age
	sm(age)
}
