package main

import (
	"fmt"
	"time"
)

type Queue struct {
	queue chan interface{}
}

func (q *Queue) Producter(req interface{}) {
	q.queue <- req
}

func (q *Queue) Consurmer() {
	go func() {
		for {
			data := <-q.queue
			fmt.Println("got ", data)
		}
	}()
}

func main() {
	q:=new(Queue)
	q.queue=make(chan interface{},3)
	q.Consurmer()
	time.Sleep(time.Second*1)
	q.Producter(1)
	time.Sleep(time.Second*2)
}
