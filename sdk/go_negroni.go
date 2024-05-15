package sdk

import (
	"github.com/urfave/negroni"
	"io"
	"net/http"
)

func NegroniClassic() {
	// New(NewRecovery(), NewLogger(), NewStatic(http.Dir("public")))
	n := negroni.Classic()
	n.UseHandler(NegroniHandle())
	// 这里会从环境变量中探测PORT变量，替换端口号
	n.Run(":8080")
}

func NegroniHandle() http.Handler {
	return http.HandlerFunc(NegroniHandler)
}

func NegroniHandler(rw http.ResponseWriter, r *http.Request) {
	rw.Header().Set("Content-Type", "application/json")
	rw.WriteHeader(http.StatusOK)
	io.WriteString(rw, `{ "Name": "world" }`)
}
