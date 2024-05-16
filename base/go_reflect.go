package base

import (
	"fmt"
	"reflect"
)

type UserInfo struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

type CallMe interface {
	CallMe(gender bool) string
}

func (userInfo UserInfo) String() string {
	return fmt.Sprintf("name: %s, age: %d", userInfo.Name, userInfo.Age)
}

func (userInfo UserInfo) CallMe(gender bool) string {
	return fmt.Sprintf("name: %s, age: %d, gender:%t", userInfo.Name, userInfo.Age, gender)
}

/*
UsingTypeOf

	go的反射中，任意对象分为reflect.Value和reflect.Type，而reflect.Value又同时持有一个对象reflect.Value和reflect.Type，
	所以可以通过reflect.Value的Interface还原对象
	golang最基础类型
	const (
		Invalid Kind = iota
		Bool
		Int
		Int8
		Int16
		Int32
		Int64
		Uint
		Uint8
		Uint16
		Uint32
		Uint64
		Uintptr
		Float32
		Float64
		Complex64
		Complex128
		Array
		Chan
		Func
		Interface
		Map
		Ptr
		Slice
		String
		Struct
		UnsafePointer
	)
*/
func UsingTypeOf() {
	user := UserInfo{
		Name: "alone",
		Age:  18,
	}
	type_ := reflect.TypeOf(user)
	// base.UserInfo
	fmt.Println(type_)
	// 获取type类型
	fmt.Printf("%T\n", user)
	value := reflect.ValueOf(user)
	fmt.Println(value)
	// reflect.Value
	fmt.Printf("%T\n", value)
	// 获取原始value
	fmt.Printf("%v\n", value)
	// 将reflect.Value转换为UserInfo
	user2 := value.Interface().(UserInfo)
	fmt.Println(user2)
	// base.UserInfo，从这里可以看出reflect.Value持有的Type确实为原有类型
	fmt.Println(value.Type())
	// 获取类型的底层类型（基础数据类型或结构型）
	// struct
	fmt.Println(type_.Kind())
}

func RecursionField() {
	user := UserInfo{
		Name: "alone",
		Age:  18,
	}
	type_ := reflect.TypeOf(user)
	for i := 0; i < type_.NumField(); i++ {
		fmt.Println(type_.Field(i).Name)
	}
	for i := 0; i < type_.NumMethod(); i++ {
		fmt.Println(type_.Method(i).Name)
	}
}

func ModifyField() {
	user := UserInfo{
		Name: "alone",
		Age:  18,
	}
	fmt.Println(user)
	// 获取结构体地址
	value := reflect.ValueOf(&user)
	// 根据字段名修改字段值
	field := value.Elem().FieldByName("Age")
	fmt.Printf("%T\n", field)
	if field.IsValid() && field.CanSet() {
		field.SetInt(30)
	}
	// 修改完后原始对象和反射对象值均被修改
	fmt.Println(value)
	fmt.Println(user)
}

func ReflectInvoke() {
	user := UserInfo{
		Name: "alone",
		Age:  18,
	}
	value := reflect.ValueOf(user)
	// 获取方法信息
	method_ := value.MethodByName("CallMe")
	if method_.IsValid() {
		// 获取参数信息
		args := []reflect.Value{
			reflect.ValueOf(true),
		}
		fmt.Println(method_.Call(args))
	}

}
