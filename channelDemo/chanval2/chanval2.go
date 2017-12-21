package main

import (
	"fmt"
	"time"
)

type Counter struct {
	count int
}

var mapChan = make(chan map[string]Counter, 1)

func main() {
	synChan := make(chan struct{}, 2)
	go func() {
		for {
			if elem, ok := <-mapChan; ok {
				counter := elem["count"]
				counter.count++
			} else {
				break
			}
		}
		fmt.Println("Stopped. [receiver")
		synChan <- struct{}{}
	}()
	go func() {
		countMap := map[string]Counter{
			"count": Counter{count: 1},
		}
		for i := 0; i < 5; i++ {
			mapChan <- countMap
			time.Sleep(time.Millisecond)
			fmt.Printf("The count map : %v .[sender]\n", countMap)
		}
		close(mapChan)
		synChan <- struct{}{}
	}()
	<-synChan
	<-synChan
}
