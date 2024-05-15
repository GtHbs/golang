package base

import (
	"encoding/json"
	"fmt"
	"reflect"
)

func JsonConvStr() {
	var user UserInfo
	info := `{"Name":"alone","Age":18}`
	// json反序列化,原理是利用UserInfo的字段Tag，然后把解析的json对应的值赋值给他们
	err := json.Unmarshal([]byte(info), &user)
	if err != nil {
		panic(err)
	} else {
		fmt.Println(user)
	}
	value, err := json.Marshal(user)
	if err != nil {
		panic(err)
	} else {
		fmt.Println(string(value))
	}
}

func GetFieldTag() {
	var user UserInfo
	type_ := reflect.TypeOf(user)
	for i := 0; i < type_.NumField(); i++ {
		field := type_.Field(i)
		fmt.Println(i, field.Tag)
	}
}

type JsonTag struct {
	Key string `json:"key" bson:"keys"`
}

func GetJsonTag() {
	var jsonTag JsonTag
	type_ := reflect.TypeOf(jsonTag)
	for i := 0; i < type_.NumField(); i++ {
		field := type_.Field(i)
		fmt.Println(field.Tag.Get("json"), field.Tag.Get("bson"))
	}
}
