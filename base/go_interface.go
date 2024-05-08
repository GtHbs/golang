package base

import (
	"bytes"
	"fmt"
	"io"
)

func print() {
	var b bytes.Buffer
	// 将数据写到缓冲区
	fmt.Fprintf(&b, "Hello World")
	// 将缓冲区的数据写出
	fmt.Println(b.String())

	var buffer bytes.Buffer
	fmt.Fprintf(&buffer, "Hello World")
	var writer io.Writer
	// bytes.Buffer实现了io.Writer接口,所以这里可以赋值,类似于Java中的接口实现类
	writer = &buffer
	fmt.Println(writer)
}

type animal interface {
	printInfo()
}

type cat int
type dog int
type mouse int

func (c cat) printInfo() {
	fmt.Println("cat")
}

// 方法参数中传入类类型，即实现接口
func (d dog) printInfo() {
	fmt.Println("dog")
}

func (m *mouse) printInfo() {
	fmt.Println("mouse")
}

func printAnimal() {
	var a animal
	var c cat
	a = c
	a.printInfo()
	var d dog
	a = d
	d.printInfo()

	// 实体类型以值接收者实现接口的时候，不管是实体类型的值，还是实体类型的指针，都实现了该接口
	invoke(&c)

	// 如果实现方法接收参数为指针，则必须传入指针。
	// 如果实现方法接收参数为值类型，则可以传入指针，也可以传入值
	var m mouse
	invoke(&m)
}

func invoke(a animal) {
	a.printInfo()
}
