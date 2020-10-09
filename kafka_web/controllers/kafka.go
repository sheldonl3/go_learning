package controllers

import (
	"fmt"
	"github.com/Shopify/sarama"
	"github.com/astaxie/beego"
	"sync"
)

type KafkaController struct {
	beego.Controller
}

func (c *KafkaController) Get() {
	res := Consumer("kafka_web")
	c.Data["data"] = res
	c.Data["Website"] = "beego.me"
	c.Data["Email"] = "astaxie@gmail.com"
	c.TplName = "index.tpl"

}

var wg sync.WaitGroup

func Consumer(topic string) []string {
	data := []string{}
	config := sarama.NewConfig()
	config.Consumer.Return.Errors = true
	config.Version = sarama.V0_11_0_2

	consumer, err := sarama.NewConsumer([]string{"127.0.0.1:9092"}, config)
	if err != nil {
		fmt.Printf("fail to start consumer, err:%v\n", err)
		return nil
	}
	partitionList, err := consumer.Partitions(topic) // 根据topic取到所有的分区,消费者组
	if err != nil {
		fmt.Printf("fail to get list of partition:err%v\n", err)
		return nil
	}
	fmt.Println("partitionList", partitionList)
	for partition := range partitionList { // 遍历所有的分区
		// 针对每个分区创建一个对应的分区消费者
		fmt.Printf("int partition %d\n", partition)
		wg.Add(1)
		pc, err := consumer.ConsumePartition(topic, int32(partition), sarama.OffsetOldest)//消费者组
		if err != nil {
			fmt.Printf("failed to start consumer for partition %d,err:%v\n", partition, err)
			return nil
		}
		defer pc.AsyncClose()
		// 异步从每个分区消费信息
		go func(sarama.PartitionConsumer) {
			fmt.Println("start go route")
			defer wg.Done()
			//for msg := range pc.Messages() {
			//	fmt.Printf("Partition:%d Offset:%d Value:%v\n", msg.Partition, msg.Offset, msg.Value)
			//	data[int(partition)] = append(data[int(partition)], string(msg.Value))
			//}
			for {
				select {
				case msg := <-pc.Messages():
					fmt.Printf("Partition:%d Offset:%d Value:%v\n", msg.Partition, msg.Offset, msg.Value)
					data = append(data,string(msg.Value))
					if len(pc.Messages()) <= 0 { // 如果现有数据量为0，跳出循环
						return
					}
				case err := <-pc.Errors():
					fmt.Printf("err :%s\n", err.Error())
				}
			}
			fmt.Println("go route return")
		}(pc)
	}
	wg.Wait()
	fmt.Println("get kafka data successfully")
	return data
}
