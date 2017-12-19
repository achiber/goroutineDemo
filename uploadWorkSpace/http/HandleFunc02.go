package main
import (
	"io"
	"net/http"
	"fmt"
)

func helloHandler(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "Hello, albert!")
	fmt.Println("sucess!")
}

func main() {
	hh := http.HandlerFunc(helloHandler)
	http.Handle("/", hh)
	http.ListenAndServe(":12345", nil)
}

