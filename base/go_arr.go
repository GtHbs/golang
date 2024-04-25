package base

import "fmt"

func init() {
	// 1. 先声明变量
	var arr [5]int
	// 2. 初始化数组
	arr = [5]int{1, 2, 3, 4, 5}
	fmt.Println(arr)

	// 声明和初始化一步生成
	arr1 := [5]int{1, 2, 3, 4, 5}
	fmt.Println(arr1)

	// 不指定数组长度
	arr2 := [...]int{1, 2, 3, 4, 5}
	fmt.Println(arr2)

	// 指定初始化部分位置的元素
	arr3 := [...]int{0, 1, 0, 2, 3}
	fmt.Println(arr3)

	// 指定初始化1和2位置的元素
	arr4 := [...]int{1: 1, 2: 3}
	fmt.Println(arr4)

	// 访问固定位置的元素
	fmt.Println(arr4[2])
	// 修改固定位置的元素
	arr4[2] = 2

	// 遍历数组
	for i := 0; i < len(arr4); i++ {
		fmt.Println(arr4[i])
	}

	// 只遍历数组下标
	for i := range arr4 {
		fmt.Println(arr4[i])
	}

	// 同时遍历下标和元素
	for i, v := range arr4 {
		fmt.Println(i, v)
	}
}

func refer() {
	arr := [5]int{1, 2, 3, 4, 5}
	var arr2 [5]int = arr
	fmt.Println(arr2)
	//var arr3 [4]int = arr // 只有相同长度的数组才可以相互赋值
}

func pointerArr() {
	// 指针数组，初始化1和3下标位置，为其开辟了内存空间，其他位置为nil
	arr := [5]*int{1: new(int), 2: new(int)}
	// 注意指针数组和数组指针
	// 指针数组：[...]*int
	// 数组指针: *[...]int{}
	fmt.Println(arr)
	// 为指针赋值，只能给初始化内存的下标地址赋值
	*arr[1] = 1
	// 初始化
	arr[0] = new(int)
	// 获取下标为1的指针指向的值
	fmt.Println(*arr[1])
}

/*
*
修改数组，入参为大小为5的int数组
*/
func modifyArr(arr [5]int) {
	arr[1] = 2
	fmt.Println(arr)
}

func transmitArr() {
	arr := [5]int{1, 2, 3, 4, 5}
	// 这里入参有个问题，传入的是值传递，即将arr复制了一份传给modifyArr，modifyArr修改后当前数组没影响
	modifyArr(arr)
	fmt.Println(arr)
}

func modifyArrByPointer(arr *[5]int) {
	arr[0] = 2
	fmt.Println(*arr)
}

func transmitArrByPointer() {
	arr := [5]int{1, 2, 3, 4, 5}
	// 数组长度不固定的不能用指针传递
	modifyArrByPointer(&arr)
	fmt.Println(arr)
}
