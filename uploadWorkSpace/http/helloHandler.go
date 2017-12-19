package main

import (
	//"io"
	"net/http"
	"fmt"
)

type helloHandler struct {}
func (h *helloHandler)ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello, zxh"))
	fmt.Println("sucess")
}



func main() {
	http.Handle("/", &helloHandler{})
	http.ListenAndServe(":12345", nil)
}
