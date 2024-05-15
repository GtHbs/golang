package sdk

import (
	"fmt"
	"github.com/gorilla/context"
	"net/http"
	"strconv"
)

func InitGorillaContextHttpServer() {
	//http.Handle("/", http.HandlerFunc(GorillaHandler))
	// 可以使用该方法自动清理context
	http.Handle("/", context.ClearHandler(http.HandlerFunc(GorillaHandler)))
	http.ListenAndServe(":8080", nil)
}
func GorillaHandler(rw http.ResponseWriter, r *http.Request) {
	// 底层可以看出，是使用了map存储request和request创建时间，也使用读写锁保证goroutine下安全
	context.Set(r, "user", "alone")
	context.Set(r, "age", 18)
	GorillaDoHandler(rw, r)
}

func GorillaDoHandler(rw http.ResponseWriter, r *http.Request) {
	user := context.Get(r, "user").(string)
	age := context.Get(r, "age").(int)
	value, exist := context.GetOk(r, "user")
	if exist {
		fmt.Println(value)
	}
	valueMap := context.GetAll(r)
	for key, val := range valueMap {
		fmt.Println(key, val)
	}
	// 保留10s内的的request
	context.Purge(10)
	context.Delete(r, "user")
	context.Clear(r)
	rw.WriteHeader(http.StatusOK)
	rw.Write([]byte(user + "," + strconv.Itoa(age)))
}
