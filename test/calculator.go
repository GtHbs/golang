package test

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
)

func Add(a, b int) int {
	return a + b
}

func Subtract(a, b int) int {
	return a - b
}

func Routes() {
	http.HandleFunc("/sendJson", SendJson)
}

func SendJson(rw http.ResponseWriter, req *http.Request) {
	u := struct {
		Name string
	}{
		Name: "John Doe",
	}
	rw.Header().Set("Content-Type", "application/json")
	rw.WriteHeader(http.StatusOK)
	json.NewEncoder(rw).Encode(u)
}

func MockServer() *httptest.Server {
	sendJson := func(w http.ResponseWriter, r *http.Request) {
		u := struct {
			Name string
		}{
			Name: "John Doe",
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(u)
	}
	// HandlerFunc接口实现为匿名函数
	return httptest.NewServer(http.HandlerFunc(sendJson))
}

func Tag(tag string) string {
	switch tag {
	case "1":
		return "A"
	case "2":
		return "B"
	case "3":
		return "C"
	default:
		return "D"
	}
}
