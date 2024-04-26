package base

import "fmt"

/*
函数的定义
函数名首字母小写，就只能作用于当前包使用
函数名首字母大写，可以被其他包使用
*/
func add(a int, b int) int {
	return a + b
}

func Add(a int, b int) int {
	return a + b
}

type person struct {
	Name string
	Age  int
}

func (p person) String() (string, error) {
	return fmt.Sprintf("%s: %d", p.Name, p.Age), nil
}

/*
只有指针类型的才可以修改struct
*/
func (p *person) modifyPerson() {
	p.Age = 30
}

func funcMethod() {
	p := person{
		Name: "alone",
		Age:  20,
	}
	fmt.Println(p.String())
	//p.modifyPerson()
	(&p).modifyPerson() //该操作等同于上面
	fmt.Println(p.String())
}
