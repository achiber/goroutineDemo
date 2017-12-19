package main

import (
	"io"
	"fmt"
	"net/http"
)

func helloHandle(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "Hello, wwh!\n")
	fmt.Println("sucess!")
}

func main() {
	http.HandleFunc("/", helloHandle)
	http.ListenAndServe(":12345", nil)
}
