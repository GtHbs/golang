package test

import (
	"io/ioutil"
	"log"
	"net/http"
	"testing"
)

func TestAdd(t *testing.T) {

	if Add(1, 2) == 3 {
		t.Log("1+2=3")
	}

	if Subtract(2, 3) == -1 {
		t.Log("2-2=-1")
	}

	/**
	表组测试
	*/
	sum := Add(2, 2)
	if sum == 4 {
		t.Log("2+2=4")
	} else {
		t.Fatal("result is wrong")
	}

	sum = Subtract(2, 2)
	if sum == 0 {
		t.Log("2-2=0")
	} else {
		t.Fatal("result is wrong")
	}
}

/**
模拟服务端响应
*/
//func TestSendJson(t *testing.T) {
//	Routes()
//	req, err := http.NewRequest(http.MethodGet, "/sendJson", nil)
//	if err != nil {
//		t.Fatal(err)
//	}
//	// 利用http.ResponseWriter模拟服务端响应
//	rw := httptest.NewRecorder()
//	http.DefaultServeMux.ServeHTTP(rw, req)
//	log.Println("code:", rw.Code)
//	log.Println("body:", rw.Body.String())
//}

func TestSendJson(t *testing.T) {
	server := MockServer()
	defer server.Close()
	resp, err := http.Get("http://" + server.Listener.Addr().String() + "/json")
	if err != nil {
		t.Fatal(err)
	}
	defer resp.Body.Close()
	log.Println("code:", resp.StatusCode)
	json, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		t.Fatal(err)
	}
	log.Println(string(json))
}

// 使用命令，go test -v -coverprofile=c.out ，指定输出文件
// 使用命令，go tool cover -html=c.out -o=tag.html，来输出哪些用例未测试
func TestTag(t *testing.T) {
	Tag("1")
	Tag("2")
	Tag("3")
	Tag("4")
}
