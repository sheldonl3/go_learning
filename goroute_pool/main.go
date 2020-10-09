package main

import ( //https://mp.weixin.qq.com/s/Aoc_FUPR6Fm_XDprvr6dqQ
	"log"
	"pool/worker"
	"sync"
	"time"
)
//通过无缓冲通道实现协程（goroutine）池。

var langs = []string{
	"Golang",
	"PHP",
	"JavaScript",
	"Python",
	"Java",
}

type langPrinter struct {
	lang string
}

func (m *langPrinter) Task() {
	log.Println(m.lang)
	time.Sleep(time.Second)
}

func main() {
	p := worker.New(2)

	var wg sync.WaitGroup
	wg.Add(5 * len(langs))

	for i := 0; i < 5; i++ {
		for _, lang := range langs {
			lp := langPrinter{lang}
			go func() {
				p.Run(&lp)
				wg.Done()
			}()
		}
	}

	wg.Wait()
	p.Shutdown()
}
