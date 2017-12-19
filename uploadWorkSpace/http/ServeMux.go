package main
import (
	"io"
	"net/http"
	"fmt"
)
func helloHandler(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "小惠惠，臭狗屎!")
	fmt.Println("sucess02!")
}

func echoHandler(w http.ResponseWriter, r *http.Request) {
	//io.WriteString(w, r.URL.Path)
	io.WriteString(w, "haha -^-^- !")
	fmt.Println("sucess01!")
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/hello", helloHandler)
	mux.HandleFunc("/echo", echoHandler)
	http.ListenAndServe(":12345", mux)
}
