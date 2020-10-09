package main

import (
	"fmt"
	"sync"
	"time"
)

var ch chan interface{}
var wait sync.WaitGroup

func consumer(ch chan interface{}) {
	defer func() {
		fmt.Println("channel close quit")
		wait.Done()
	}()
	for {
		select {
		case data, is_open := <-ch: //channel关闭之后，读取到的是空值
			fmt.Println("get from channel ", data, "is open ", is_open)
			if is_open == false {
				return
			}
			time.Sleep(time.Second * 1)
		default:
			time.Sleep(time.Second * 1)
			fmt.Println("got nil")
		}
	}
}

func produser(ch chan interface{}) {
	for i := 0; i < 10; i++ {
		ch <- i
	}
	close(ch)
	wait.Done()
}

func main() {
	wait.Add(2)
	ch = make(chan interface{}, 9)
	go consumer(ch)
	go produser(ch)
	wait.Wait()
}
