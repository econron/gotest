package main

import(
	"fmt"
	"net/http"
)

type MyHandler struct{}

func (h *MyHandler) ServerHTTP(w http.ResponseWriter, r *http.Request){
	fmt.Fprintf(w, "Hello World!")
}

func main(){
	handler := MyHandler{}
	server := http.Server{
		Addr: "127.0.0.1:8090",
		// テキストの通りだとエラーになる。
		Handler: &handler,
	}
	server.ListenAndServe()
}