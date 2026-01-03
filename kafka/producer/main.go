package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/IBM/sarama"
)

func main() {
	// 1. 创建生产者配置
	config := sarama.NewConfig()
	config.Producer.RequiredAcks = sarama.WaitForAll // 等待所有副本确认
	config.Producer.Retry.Max = 5                    // 重试次数
	config.Producer.Return.Successes = true          // 成功交付的消息将在success channel返回

	// 2. 连接 Kafka 集群
	producer, err := sarama.NewSyncProducer([]string{"192.168.0.185:9092"}, config)
	if err != nil {
		log.Fatalf("Error creating producer: %v", err)
	}
	defer producer.Close()

	// 3. 构造并发送消息
	topic := "test-topic"
	msg := &sarama.ProducerMessage{
		Topic: topic,
		Value: sarama.StringEncoder("Hello, Kafka from Go!"),
		Key:   nil,
	}

	partition, offset, err := producer.SendMessage(msg)
	if err != nil {
		log.Fatalf("Error sending message: %v", err)
	}

	// 4. 打印消息发送结果
	log.Printf("Message sent successfully! Partition: %d, Offset: %d\n", partition, offset)

	// 5. 通过信号量保持运行，优雅关闭
	signals := make(chan os.Signal, 1)
	signal.Notify(signals, os.Interrupt, syscall.SIGTERM)
	<-signals
	log.Println("Shutting down producer...")
}
