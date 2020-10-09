package main

import (
	"fmt"
	"runtime"
	"sync"
)

//使用pool保存和复用临时对象，以减少内存分配，降低CG压力。
//每次需要临时变量的话，从pool中取，pool没有的话会自动创建；有的话直接取出来用
func poolforsclice() {
	//我们创建一个Pool，并实现New()函数
	sp := sync.Pool{
		//New()函数的作用是当我们从Pool中Get()对象时，如果Pool为空，则先通过New创建一个对象，插入Pool中，然后返回对象。
		New: func() interface{} {
			return make([]int, 16)
		},
	}
	item := sp.Get()
	//打印可以看到，我们通过New返回的大小为16的[]int
	fmt.Println("item : ", item)

	//然后我们对item进行操作
	//New()返回的是interface{}，我们需要通过类型断言来转换
	for i := 0; i < len(item.([]int)); i++ {
		item.([]int)[i] = i
	}
	fmt.Println("item : ", item)

	//使用完后，我们把item放回池中，让对象可以重用
	sp.Put(item)
	runtime.GC() //有可能被gc回收
	//再次从池中获取对象
	item2 := sp.Get()
	//注意这里获取的对象就是上面我们放回池中的对象
	fmt.Println("item2 : ", item2)
	//我们再次获取对象
	item3 := sp.Get()
	//因为池中的对象已经没有了，所以又重新通过New()创建一个新对象，放入池中，然后返回
	//所以item3是大小为16的空[]int
	fmt.Println("item3 : ", item3)

	//测试sync.Pool保存socket长连接池
	//testTcpConnPool()
}

func main() {
	poolforsclice()
}
