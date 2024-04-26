package base

import (
	"bytes"
	"fmt"
)

func print() {
	var b bytes.Buffer
	// 将数据写到缓冲区
	fmt.Fprintf(&b, "Hello World")
	// 将缓冲区的数据写出
	fmt.Println(b.String())
}
