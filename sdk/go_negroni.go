package sdk

import (
	"github.com/urfave/negroni"
	"io"
	"net/http"
)

func InitNegorniServer() {
	server := negroni.Classic()
	server.UseHandler(negroniHandler())
	server.Run(":8080")

}

func negroniHandler() http.Handler {
	return http.HandlerFunc(negroniHandle)
}

func negroniHandle(rw http.ResponseWriter, r *http.Request) {
	rw.Header().Add("content-type", "text/json")
	io.WriteString(rw, `{"alive": true}`)
}

func InitNegroniMuxServer() {
	server := negroni.Classic()
	mux := http.NewServeMux()
	mux.Handle("/", negroniHandler())
	mux.HandleFunc("/test", func(rw http.ResponseWriter, r *http.Request) {
		rw.Header().Add("content-type", "text/json")
		io.WriteString(rw, `{"name": "alone"}`)
	})
	server.UseHandler(mux)
	server.Run(":8080")
}
