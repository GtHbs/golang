package base

import (
	"fmt"
	"go/types"
)

func baseType() {
	// 基础类型包括数值、浮点、字符、布尔类型，这种类型的操作会直接返回一个新的副本，在函数间传递也是值传递，即传递的是值的副本
}

func referType() {
	/*
		引用类型包括切片、map、函数、以及chan
		引用类型实际上是创建了个指针，指向了底层的数据结构
	*/
}

type Address struct {
	Country  string
	Province string
	City     string
}

type Person struct {
	Name    string
	Age     int
	Address Address
}

func structType() {
	/*
		结构性类似于Java的对象
	*/
	var person Person // 默认初始化struct，struct的默认值为字段的零值
	person.Name = "alone"
	fmt.Println(person)

	//aloneAddress := Address{Country: "China", Province: "Beijing", City: "Beijing"}
	// 你们初始化嵌套struct
	alone := Person{Name: "Alone", Age: 18, Address: Address{
		Country:  "China",
		Province: "Beijing",
		City:     "Beijing",
	}} // 显式的初始化struct

	fmt.Println(alone)
}

/*
*
只有引用传递的结构体才能修改结构体的值
*/
func modifyStruct(person *Person) {
	person.Name = "alone2"
}

func referStruct() {
	alone := Person{Name: "Alone", Age: 18, Address: Address{
		Country:  "China",
		Province: "Beijing",
		City:     "Beijing",
	}}
	modifyStruct(&alone)
}

type Duration int64

// 自定义类型可以为结构体、基础类型、引用类型
type Student Person

type CustomSlice types.Slice

func customType() {
	// 自定义类型
	duration := Duration(10)
	fmt.Println(duration)
	//var durationBase int64
	//durationBase = duration		// 虽然类型一样，但是不能强行赋值
}
