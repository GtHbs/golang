package main

import "fmt"

func main() {
	//f.Println(os.Args)
	arr := [5]int{1, 2, 3, 4, 5}
	slice1 := arr[:]  // 拷贝数据到切片
	slice2 := arr[:2] // 拷贝0-1下标的数据到切片
	fmt.Println(slice1, slice2)
	slice2[0] = 10
	slice2[1] = 20

	//slice1 = append(slice2, 1)
	//fmt.Println(slice1, slice2)
	slice1 = append(slice2, slice1...)
	fmt.Println(slice1, slice2)
}
