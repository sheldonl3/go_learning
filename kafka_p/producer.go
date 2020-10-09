package main

import (
	"fmt"
	"github.com/Shopify/sarama"
	"math/rand"
)

//zk start
//sudo bin/zookeeper-server-start.sh config/zookeeper.properties
//kafka start
//sudo bin/kafka-server-start.sh config/server.properties
//create topic
//sudo bin/kafka-topics.sh --create --zookeeper localhost:2181 --replication-factor 3 --partitions 3 --topic test
//list topic
//./bin/kafka-topics.sh --zookeeper localhost:2181 --list
//delete topic
//./bin/kafka-topics.sh --delete --zookeeper localhost:2181 --topic test
func main() {
	config := sarama.NewConfig()
	config.Producer.RequiredAcks = sarama.WaitForAll          // 发送完数据需要leader和follow都确认
	config.Producer.Partitioner = sarama.NewRandomPartitioner // 新选出一个partition
	config.Producer.Return.Successes = true                   // 成功交付的消息将在success channel返回

	// 构造一个消息
	msg := &sarama.ProducerMessage{}
	msg.Topic = "kafka_web"
	s:=fmt.Sprintf("this is a test log %d",rand.Intn(100))
	msg.Value = sarama.StringEncoder(s)
	// 连接kafka
	client, err := sarama.NewSyncProducer([]string{"127.0.0.1:9092"}, config)
	if err != nil {
		fmt.Println("producer closed, err:", err)
		return
	}
	defer client.Close()
	// 发送消息
	pid, offset, err := client.SendMessage(msg)
	if err != nil {
		fmt.Println("send msg failed, err:", err)
		return
	}
	fmt.Printf("pid:%v offset:%v\n", pid, offset)
}
