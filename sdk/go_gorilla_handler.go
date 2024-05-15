package sdk

import (
	"github.com/gorilla/handlers"
	"io"
	"net/http"
	"os"
)

func InitGorillaHandlerHttpServer() {
	//http.Handle("/", http.HandlerFunc(GorillaHandler))
	// 可以使用该方法自动清理context
	http.Handle("/", UsingCanonicalHostHandler(GorillaHandlerHandler()))
	http.Handle("/zip", UsingCompressHandler(GorillaHandlerHandler()))
	http.ListenAndServe(":8080", nil)
}

func GorillaHandlerHandler() http.Handler {
	return http.HandlerFunc(GorillaZipHandlerHandle)
}

func GorillaHandlerHandle(rw http.ResponseWriter, r *http.Request) {
	rw.WriteHeader(http.StatusOK)
	io.WriteString(rw, "Hello World!")
}

func GorillaZipHandlerHandle(rw http.ResponseWriter, r *http.Request) {
	rw.WriteHeader(http.StatusOK)
	rw.Header().Set("Content-Type", "text/plain")
	io.WriteString(rw, "Hello World")
}

/*
UsingLoggingHandler
该方法是对handler进行包装
*/
func UsingLoggingHandler(next http.Handler) http.Handler {
	return handlers.LoggingHandler(os.Stdout, next)
}

func UsingCombinedLoggingHandler(handler http.Handler) http.Handler {
	return handlers.CombinedLoggingHandler(os.Stdout, handler)
}

func UsingCompressHandler(handler http.Handler) http.Handler {
	return handlers.CompressHandler(handler)
}

func UsingContentTypeHandler(handler http.Handler) http.Handler {
	return handlers.ContentTypeHandler(handler, "application/json")
}

// 重定向
func UsingCanonicalHostHandler(handler http.Handler) http.Handler {
	return handlers.CanonicalHost("http://www.baidu.com", http.StatusFound)(handler)
}
