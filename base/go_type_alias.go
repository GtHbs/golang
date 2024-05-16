package base

import "fmt"

/*
作用：用于已经定义的类型，在package之间的移动时兼容。
例如：原有包base/T1需要移动到sdk/T1，这样移动会导致所有用到base/T1的代码都要修改路径，否则无法使用
*/

type age1 int   // 创建新类型age1
type age2 = int // 创建别名age2，使用=来定义别名

func UsingTypeAlias() {
	var i int = 1
	//var i1 age1 = i	// error 类型转换错误,这是因为age1是一个单独的类型，不是int所以不能强制类型转换
	//var i2 age2 = i	// age2是int的别名，本质上还是int，所以是可以类型转换的
	//fmt.Println(i1, i2)
	fmt.Println(i)
	var u1 user1
	var u2 user2
	u1.m1()
	u2.m2()

	fmt.Println("---------------")
	var u3 UserInfo
	var i1 I = u2
	var i2 I = u3 // 因为u2是UserInfo的别名，其实现了I接口，所以可以推导出来u3也实现了I接口
	fmt.Println(i1, i2)
}

func (i age1) m1() {
	fmt.Println("age1")
}

//func (i age2) m2() {		// 编译失败，因为int是基础类型，无法为其重写方法
//	fmt.Println("age2")
//}

type user1 UserInfo
type user2 = UserInfo

func (u user1) m1() {
	fmt.Println("user1")
}
func (u user2) m2() { // 编译可以通过，这里为user2实现m2方法就等于是为UserInfo实现m2方法
	fmt.Println("user2")
}

type I interface {
	m3()
}

func (u user2) m3() {
	fmt.Println("user3")
}

type I1 I
type I2 = I

type MyI int

func (u MyI) m3() {
	fmt.Println("myI")
}

func UsingInterfaceType() {
	// MyI实现了接口I，I1是I的类型，I2是I的别名，所以三者之间可以互相赋值
	var i I = MyI(1)
	var i1 I1 = MyI(1)
	var i2 I2 = MyI(1)
	i = i1
	i = i2
	i1 = i2
	i1 = i
	i2 = i
	i2 = i1
}

type T1 struct {
	// 不允许出现循环引用
	//T2
}

type T2 = T1

type MyStruct struct {
	T2
	T1
}

func (t T1) m3() {
	fmt.Println("t1")
}

func StructAlias() {
	//my := MyStruct{}
	// base/go_type_alias.go:95:5: ambiguous selector my.m3
	// 因为结构体中有T1和T2两种类型，并且T2为T1的别名，T1实现了m3接口，等于T2也实现了，
	// 所以这里调用m3才会产生错误

	//my.m3()

}

//type byte = uint8
//type rune = int32

type unPushed struct {
	Name string `json:"name"`
}

func (u unPushed) getName() string {
	return u.Name
}

type UnPushed = unPushed // 可以使用这种方式将不可以导出的类型导出
