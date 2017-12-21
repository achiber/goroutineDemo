package main

import "fmt"

func main() {
	intChan := make(chan int, 10)
	for i := 0; i < 10; i++ {
		intChan  <- i
	}
	close(intChan)
	syncChan := make(chan struct{}, 1)
	go func() {
		loop:
			for {
				select {
				case e, ok := <- intChan:
					if !ok {
						fmt.Println("End!")
						break loop
					}
					fmt.Printf("Received: %v\n", e)
				}
			}
			syncChan <- struct{}{}
	}()
	<- syncChan
}