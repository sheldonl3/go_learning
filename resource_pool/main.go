package main

import ( //https://mp.weixin.qq.com/s/ubMCwHM5X1OLKM8kM7yq_g
	"io"
	"log"
	"math/rand"
	"pool/pool"
	"sync"
	"sync/atomic"
	"time"
)

/*
通过缓冲通道（channel）实现共享资源池。
该资源池可用于管理任意数量的协程（goroutine）之间共享的资源（比如数据库连接），
如果某个协程需要从资源池获取资源（比如从数据库连接池获取数据库连接），可以从共享资源池申请（如果没有的话需要新建），并且在使用完成后将其归还到共享资源池。
*/

const (
	maxGoroutines   = 5
	pooledResources = 2
)

type dbConnection struct {
	ID int32
}

func (dbConn *dbConnection) Close() error {
	log.Println("Close: Connection", dbConn.ID)
	return nil
}

var idCounter int32

func createConnection() (io.Closer, error) {
	id := atomic.AddInt32(&idCounter, 1)
	log.Println("Create: New Connection", id)

	return &dbConnection{id}, nil
}

func main() {
	var wg sync.WaitGroup
	wg.Add(maxGoroutines)

	p, err := pool.New(createConnection, pooledResources)
	if err != nil {
		log.Println(err)
	}

	for query := 0; query < maxGoroutines; query++ {
		go func(q int) {
			performQueries(q, p)
			wg.Done()
		}(query)
	}

	wg.Wait()

	log.Println("Shutdown Program.")
	p.Close()
}

func performQueries(query int, p *pool.Pool) {
	conn, err := p.Acquire()
	if err != nil {
		log.Println(err)
		return
	}
	defer p.Release(conn)

	// 模拟 SQL 查询
	time.Sleep(time.Duration(rand.Intn(1000)) * time.Millisecond)
	log.Printf("QID[%d] CID[%d]\n", query, conn.(*dbConnection).ID)
}
