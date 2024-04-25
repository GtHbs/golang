package base

import "fmt"

func init() {
	// 切片是动态数组，底层用数组生成，类似于Java的List
	// 创建一个长度为5，容量为10（底层数组的长度）的切片，容量必须大于等于长度
	slice := make([]int, 5, 10)
	fmt.Println(slice)
	// 创建容量和大小都为5的切片
	slice2 := make([]int, 5)
	fmt.Println(slice2)
	// 不使用make创建切片
	slice3 := []int{1, 2, 3}
	fmt.Println(slice3)
}

func sliceType() {
	// nil切片表示一个不存在的切片
	var nilSlice []int
	fmt.Printf("%T\n", nilSlice)
	// 空切片表示切片指向的底层数组为空
	slice := []int{}
	fmt.Printf("%T\n", slice)
}

func sliceArr() {
	arr := [5]int{1, 2, 3, 4, 5}
	slice1 := arr[:]  // 拷贝数据到切片
	slice2 := arr[:2] // 拷贝0-1下标的数据到切片
	slice3 := arr[3:] // 拷贝3及之后的数据到切片
	fmt.Println(slice1, slice2, slice3)

	slice1[0] = 100
	fmt.Println(slice1, slice2, slice3) //可以看出，修改slice1的元素后，1、2以及原数组元素都发生了变化，说明底层用的是同一个数组

	// 使用len计算切片长度，cap计算切片容量
	fmt.Println(len(slice1), cap(slice1))
	// 最后一个元素表示生成一个容量为3的切片，该值不能超过数组最大长度
	slice4 := arr[1:2:3]
	fmt.Println(slice1, slice2, slice4)
}

func useSlice() {
	slice := make([]int, 5)
	fmt.Println(slice[0])
	// 使用append追加元素，如果append的时候，原数组容量不足，会重新创建一个数组，然后复制原数组的元素到新数组
	// 这里增长的方式是，容量小于1000时，每次成倍增长，超过1000后，每次增长25%
	slice = append(slice, 3)

	slice2 := slice[2:3]
	// 将slice2追加到slice中，追加方式是追加到头位置
	slice = append(slice2, slice...)

	// 切片的使用和数组一致，切片在函数间的传递和数组不同，不是传递实际数组，而是指针，所以不需要在入参出用指针接收
}
