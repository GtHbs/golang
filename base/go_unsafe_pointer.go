package base

import (
	"fmt"
	"unsafe"
)

/*
任何指针都可以转换为unsafe.Pointer
unsafe.Pointer可以转换为任何指针
*/
func PointerTransfer() {
	i := 10
	ip := &i
	fmt.Println(ip)
	//var fp *float64 = (*float64)(ip)
	var fp *float64 = (*float64)(unsafe.Pointer(ip))
	fmt.Println(fp)
}

func UsingPointer() {
	user := new(UserInfo)
	fmt.Println(user)
	// 取第一个字段的指针，不需要偏移
	pName := (*string)(unsafe.Pointer(user))
	*pName = "alones"
	// 取第二个指针需要获取结构体的地址+第二个字段的偏移量
	pAge := (*int)(unsafe.Pointer(uintptr(unsafe.Pointer(user)) + unsafe.Offsetof(user.Age)))
	*pAge = 20
	fmt.Println(*user)
}
