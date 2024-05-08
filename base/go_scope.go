package base

import "fmt"

/*
	在go中，如果类型是以大写开头，则表示其他包可以访问，小写开头只能在当前包中访问
*/

// 该变量只能在当前包中使用
type count int

func New() count {
	return count(0)
}

func login() {
	// 使用 := 符号可以推断变量类型，从而达到访问其他包私有变量的功能
	c := New()
	fmt.Println(c)

	// 这里只对外暴露了接口，但没有暴露接口具体的实现
	loginer := NewLoginer()
	loginer.Login()
}

func NewLoginer() Loginer {
	return defaultLoginer(0)
}

type Loginer interface {
	Login()
}

type defaultLoginer int

func (d defaultLoginer) Login() {
	fmt.Println("login")
}
