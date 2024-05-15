package sdk

import (
	"context"
	"net/http"
	"strconv"
)

func InitContextHttpServer() {
	http.Handle("/", http.HandlerFunc(ContextHandler))
	http.ListenAndServe(":8080", nil)
}
func ContextHandler(rw http.ResponseWriter, r *http.Request) {
	// 底层可以看出，是使用了map存储request和request创建时间，也使用读写锁保证goroutine下安全
	userContext := context.WithValue(context.Background(), "user", "alone")
	ageContext := context.WithValue(userContext, "age", 22)
	rContext := r.WithContext(ageContext)
	ContextDoHandler(rw, rContext)
}

func ContextDoHandler(rw http.ResponseWriter, r *http.Request) {
	user := r.Context().Value("user").(string)
	age := r.Context().Value("age").(int)
	rw.WriteHeader(http.StatusOK)
	rw.Write([]byte(user + "," + strconv.Itoa(age)))
}
