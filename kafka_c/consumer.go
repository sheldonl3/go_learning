package main

import (
	"fmt"
	"sync"
	"github.com/Shopify/sarama"
)

// kafka consumer

func main() {
	wg := sync.WaitGroup{}

	config := sarama.NewConfig()
	config.Consumer.Return.Errors = true
	config.Version = sarama.V0_11_0_2

	consumer, err := sarama.NewConsumer([]string{"127.0.0.1:9092"}, config)
	if err != nil {
		fmt.Printf("fail to start consumer, err:%v\n", err)
		return
	}
	partitionList, err := consumer.Partitions("kafka_web") // 根据topic取到所有的分区
	if err != nil {
		fmt.Printf("fail to get list of partition:err%v\n", err)
		return
	}
	fmt.Println(partitionList)
	for partition := range partitionList { // 遍历所有的分区
		// 针对每个分区创建一个对应的分区消费者
		wg.Add(1)
		pc, err := consumer.ConsumePartition("kafka_web", int32(partition), sarama.OffsetOldest)
		if err != nil {
			fmt.Printf("failed to start consumer for partition %d,err:%v\n", partition, err)
			return
		}
		defer pc.AsyncClose()
		// 异步从每个分区消费信息
		go func(sarama.PartitionConsumer) {
			//for msg := range pc.Messages() {
			//	fmt.Printf("Partition:%d Offset:%d Key:%v Value:%v", msg.Partition, msg.Offset, msg.Key, msg.Value)
			//}
			for {
				select {
				case msg := <-pc.Messages():
					fmt.Printf("msg offset: %d, partition: %d, timestamp: %s, value: %s\n",
						msg.Offset, msg.Partition, msg.Timestamp.String(), string(msg.Value))
				case err := <-pc.Errors():
					fmt.Printf("err :%s\n", err.Error())
				}
			}
			wg.Done()
		}(pc)
	}
	wg.Wait()
	fmt.Println("consumer exit successfully")
}
