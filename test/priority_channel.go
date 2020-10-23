package main

import (
	"fmt"
	"time"
)

/*
select语句中实现优先级: https://mp.weixin.qq.com/s/3NqPViFwIJbSeL1hrwajTw
有一个函数会持续不间断地从ch1和ch2中分别接收任务1和任务2，
如何确保当ch1和ch2同时达到就绪状态时，优先执行任务1，在没有任务1的时候再去执行任务2呢？
*/

//代码通过嵌套两个select实现了"优先级"，看起来是满足题目要求的。
//但是这代码有点问题，如果ch1和ch2都没有达到就绪状态的话，整个程序不会阻塞而是进入了死循环。
func priority_channel1(ch1, ch2 <-chan int, stopCh chan struct{}) {
	for {
		fmt.Println("runing worker1")
		select {
		case <-stopCh:
			return
		case job1 := <-ch1:
			fmt.Println(job1)
		default:
			select {
			case job2 := <-ch2:
				fmt.Println(job2)
			default:
			}
		}
	}
}

//代码在外层select选中执行job2 := <-ch2时，进入到内层select循环继续尝试执行job1 := <-ch1,
//当ch1就绪时就会一直执行，否则跳出内层select
//如果ch1和ch2都没有达到就绪状态的话,会阻塞，因为没有default
//better way
func priority_channel2(ch1, ch2 <-chan int, stopCh chan struct{}) {
	for {
		fmt.Println("runing worker2")
		select {
		case <-stopCh:
			return
		case job1 := <-ch1:
			fmt.Println(job1)
		case job2 := <-ch2:
		priority:
			for {
				select {
				case job1 := <-ch1:
					fmt.Println(job1)
				default:
					break priority //跳出for循环
				}
			}
			fmt.Println(job2)
		}
	}
}

func main() {
	ch1 := make(chan int, 1)
	ch2 := make(chan int, 1)
	chstop := make(chan struct{}, 1)
	go priority_channel2(ch1, ch2, chstop)
	//go priority_channel1(ch1,ch2,chstop)
	ch1 <- 1
	ch1 <- 1
	ch2 <- 2
	ch1 <- 1
	time.Sleep(3 * time.Second)
}
