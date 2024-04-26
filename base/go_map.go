package base

import "fmt"

func init() {
	// 创建一个key为string，value为int的字典
	dict := make(map[string]string)
	fmt.Println(dict)
	dict["name"] = "alone"
	// 创建字典，默认可以初始化为空{}
	dict2 := map[string]string{
		"name": "alone",
	}
	fmt.Println(dict2)
}
func usage() {
	dict := map[string]string{}
	dict["name"] = "alone"
	dict["age"] = "19"
	// 获取字典中的item
	age, exist := dict["age"]
	fmt.Println(age, exist)

	// 删除字典中item
	delete(dict, "age")

	for k, v := range dict {
		fmt.Println(k, v)
	}
}

func refers(dict map[string]string) {
	fmt.Println(dict)
}

func referDict() {
	dict := map[string]string{}
	// 在函数间传递字典，默认传递引用和数组不同
	refers(dict)
}
