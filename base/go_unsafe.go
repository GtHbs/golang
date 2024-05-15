package base

import (
	"fmt"
	"unsafe"
)

func UsingSizeOf() {
	fmt.Println("-------------------")
	// 1 byte
	fmt.Println(unsafe.Sizeof(true))
	// 1
	fmt.Println(unsafe.Sizeof(int8(0)))
	// 2
	fmt.Println(unsafe.Sizeof(int16(0)))
	// 4
	fmt.Println(unsafe.Sizeof(int32(0)))
	// 8
	fmt.Println(unsafe.Sizeof(int64(0)))
	// 8
	fmt.Println(unsafe.Sizeof(int(1000000)))
}

func UsingAlignOf() {
	// 该函数用于内存对齐
	var b bool
	var i8 int8
	var i16 int16
	var i32 int32
	var i64 int64
	var f32 float32
	var f64 float64
	var s string
	var m map[int]int
	var p *int32

	fmt.Println("------------------")
	// 1
	fmt.Println(unsafe.Alignof(b))
	// 1
	fmt.Println(unsafe.Alignof(i8))
	// 2
	fmt.Println(unsafe.Alignof(i16))
	// 4
	fmt.Println(unsafe.Alignof(i32))
	// 8
	fmt.Println(unsafe.Alignof(i64))
	// 4
	fmt.Println(unsafe.Alignof(f32))
	// 8
	fmt.Println(unsafe.Alignof(f64))
	// 8
	fmt.Println(unsafe.Alignof(s))
	// 8
	fmt.Println(unsafe.Alignof(m))
	// 等价于 reflect.TypeOf(x).Align()
	fmt.Println(unsafe.Alignof(p))
}

func UsingOffsetOf() {
	var user UserInfo
	fmt.Println("-----------------")
	// 16，因为Age在Name之前，Age占用8个字节，所以偏移量是8byte,这个偏移量是相对于结构体的开始位置
	fmt.Println(unsafe.Offsetof(user.Name))
}

type StructSizeA struct {
	// 1xxx
	b byte
	// 1111
	i int32
	// 11111111
	j int64
}

type StructSizeB struct {
	// 1xxxxxxx
	b byte
	// 11111111
	j int64
	// 11111111
	i int32
}

/*
StructSizeOf
内存对齐规则
1. 具体类型，对齐值=min(编译器默认对齐值，类型大小sizeof长度)。即在默认设置的对齐值和类型的内存占用大小之间，取最小值为该类型的对齐值
2. struct在每个字段对齐之后，本身也要进行对齐。对齐值=min(默认对齐值，字段最大类型长度)。即在struct中找最大的字段和默认对齐值进行对比，取最小。
*/
func StructSizeOf() {
	var s1 StructSizeA
	var s2 StructSizeB
	// 16
	// byte占用1位，不需要填充
	// int32占用4位，需要填充3位，共计8位
	// int64占用8位，前面位8位，不需要填充，共计16位
	// 结构体共计16位是8的倍数，不需要填充
	fmt.Println(unsafe.Sizeof(s1))
	// 24 可以看出，字段排列顺序不同会导致结构体大小也不同，主要是因为内存对齐的原因
	// byte占用1位，不需要填充
	// int64占用8位，前面需要填充7位，共计16位
	// int32占用4位，前面位16位正好为8的倍数，所以不用填充，共计20位
	// 结构体现在共计20位，不是8的倍数，所以需要填充4位，共计24位
	fmt.Println(unsafe.Sizeof(s2))
}
