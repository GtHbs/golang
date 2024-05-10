package base

import (
	"bytes"
	"fmt"
	"io/ioutil"
)

func IoInit() {
	// 定义零值类型的buffer
	var b bytes.Buffer
	// 使用write写入字符串
	b.Write([]byte("hello"))
	// 对buffer进行拼接
	fmt.Fprint(&b, " ", "world")
	// 将buffer中的内容输出到控制台
	//b.WriteTo(os.Stdout)
	// 除了使用WriteTo将数据写出，还可以使用io.Reader将数据读到buffer中
	//var p [100]byte
	//n, err := b.Read(p[:])
	//if err != nil {
	//	fmt.Println(err)
	//}
	//fmt.Println(n, err, string(p[:n]))
	data, err := ioutil.ReadAll(&b)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(data), err)
}
