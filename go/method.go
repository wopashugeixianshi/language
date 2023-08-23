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
	sm(age)
}
