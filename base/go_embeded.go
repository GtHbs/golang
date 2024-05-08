package base

import "fmt"

/*
	嵌入类型，可以把已有的类型声明在新的类型里的一种方式，即组合
*/

type Reader interface {
	Read(b []byte) (n int, err error)
}

type Writer interface {
	Write(b []byte) (n int, err error)
}

type Closer interface {
	Close() error
}

type ReadWriter interface {
	Reader
	Writer
}

type ReadCloser interface {
	Reader
	Closer
}

type WriteCloser interface {
	Writer
	Closer
}

type Behavior interface {
	say()
}

type User struct {
	name  string
	email string
}

//func (admin Admin) say() {
//	fmt.Println("admin say")
//}

func (user User) say() {
	fmt.Println("user say")
}

type Admin struct {
	user  User
	level int
}

func sayHello(behavior Behavior) {
	behavior.say()
}

func hello() {
	admin := Admin{user: User{name: "admin", email: "admin@admin.com"}, level: 0}
	//admin.say()
	admin.user.say()
	//sayHello(admin)
}
