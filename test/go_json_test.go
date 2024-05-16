package test

import (
	"encoding/json"
	"github.com/json-iterator/go"
	"golang/base"
	"testing"
)

var json2 = jsoniter.ConfigCompatibleWithStandardLibrary

func BenchmarkBaseJson(t *testing.B) {
	var user base.UserInfo
	t.ResetTimer()
	s := `{"name":"alone","age":20}`
	// 33100773               779.4 ns/op           256 B/op          6 allocs/op
	for i := 0; i < t.N; i++ {
		json.Unmarshal([]byte(s), &user)
	}
}

func BenchmarkJson2(t *testing.B) {
	var user base.UserInfo
	t.ResetTimer()
	s := `{"name":"alone","age":20}`
	// 123913396              200.6 ns/op            37 B/op          2 allocs/op
	for i := 0; i < t.N; i++ {
		json2.Unmarshal([]byte(s), &user)
	}
}
