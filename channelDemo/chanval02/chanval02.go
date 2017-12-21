package main

import (
	"fmt"
	"time"
)

type Counter struct {
	count int
}

var mapChan = make(chan map[string]*Counter, 1)

func (Counter *Counter) String() string {
	return fmt.Sprintf("{count:%d}", Counter.count)
}

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
		countMap := map[string]*Counter{
			"count": &Counter{count: 0},
		}
		for i := 0; i < 5; i++ {
			mapChan <- countMap
			time.Sleep(time.Millisecond)
			fmt.Println(countMap)
		}
		close(mapChan)
		synChan <- struct{}{}
	}()
	<-synChan
	<-synChan
}
