package main

import (
	"fmt"
	"sync"
)

var w sync.WaitGroup

func main() {

	chan_n := make(chan bool)
	chan_c := make(chan bool, 1)

	w.Add(2)
	go func() {
		for i := 0; i < 9; i += 2 {
			chan_c <- true
			fmt.Printf("%d ", i)
			fmt.Printf("%d ", i+1)
			chan_n <- true
		}
		w.Done()
	}()
	char_lis := []string{"a", "b", "c", "d", "e", "f", "g", "h", "i", "g"}
	go func() {
		for i := 0; i < 9; i += 2 {
			<-chan_n
			fmt.Printf("%v ", char_lis[i])
			fmt.Printf("%v ", char_lis[i+1])
			<-chan_c
		}
		w.Done()
	}()
	w.Wait()
}
