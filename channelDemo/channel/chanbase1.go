package main

import (
	"fmt"
	"time"
)

var strChan = make(chan string, 3)

func main() {
	synChan1 := make(chan struct{}, 1)
	synChan2 := make(chan struct{}, 2)
	go func() {
		<-synChan1
		fmt.Println("Received a sync signal and wait a second... [receiver]")
		time.Sleep(time.Second)
		for {
			if elem, ok := <-strChan; ok {
				fmt.Println("Received:", elem, "[receive]")
			} else {
				break
			}
		}
		fmt.Println("Stopped. [receiver]")
		synChan2 <- struct{}{}
	}()
	go func() {
		for _, elem := range []string{"a", "b", "c", "d"} {
			strChan <- elem
			fmt.Println("Sent: ", elem, "[sender")
			if elem == "c" {
				synChan1 <- struct{}{}
				fmt.Println("Sent a sync signal. [Sender]")
			}
		}
		fmt.Println("Wait 2 seconds ... [sender]")
		time.Sleep(time.Second * 2)
		close(strChan)
		synChan2 <- struct{}{}
	}()
	<-synChan2
	<-synChan2
}
