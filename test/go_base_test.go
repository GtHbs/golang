package test

import (
	"fmt"
	"strconv"
	"testing"
)

/*
基准测试
基准测试的代码文件必须以_test.go结尾
基准测试的函数必须以Benchmark开头，必须是可导出的
基准测试函数必须接受一个指向Benchmark类型的指针作为唯一参数
基准测试函数不能有返回值
b.ResetTimer是重置计时器，这样可以避免for循环之前的初始化代码的干扰
最后的for循环很重要，被测试的代码要放到循环里
b.N是基准测试框架提供的，表示循环的次数，因为需要反复调用测试的代码，才可以评估性能
go test -bench=. -run=none
-bench：表示接收一个表达式作为参数，匹配基准测试的函数，.表示运行所有基准测试
-run：匹配一个从来没有的单元测试方法，过滤掉单元测试的输出
-benchtime=3s 表示测试时间
-benchmem 显示每次操作分配内存的次数，以及每次操作分配的字节数
goos: darwin
goarch: arm64
pkg: golang/test			表示执行测试的包
BenchmarkSprintf-8      31162503   调用次数  38.60 ns/op 每次执行耗时  2 B/op 每次执行分配内存大小 1 allocs/op 每次执行分配内存次数
BenchmarkFormat-8       571271356  调用次数  2.131 ns/op 每次执行耗时  0 B/op 每次执行分配内存大小 0 allocs/op 每次执行分配内存次数
BenchmarkItoa-8         595410376  调用次数  2.355 ns/op 每次执行耗时  0 B/op 每次执行分配内存大小 0 allocs/op 每次执行分配内存次数
PASS
ok      golang/test     1.245s
*/
func BenchmarkSprintf(b *testing.B) {
	num := 10
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = fmt.Sprintf("%d", num)
	}
}

func BenchmarkFormat(b *testing.B) {
	num := int64(10)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		strconv.FormatInt(num, 10)
	}
}

func BenchmarkItoa(b *testing.B) {
	num := 10
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		strconv.Itoa(num)
	}
}
