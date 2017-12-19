package main

import(
	"fmt"
	"time"
)

func main(){
	for i := 0; i < 5000; i++ {
		go printHelloWorld(i)
	}
	time.Sleep(time.Millisecond*10)
}
func printHelloWorld(i int) {
	for {
		fmt.Println(i)
	}
}


