package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"sync"
	"syscall"

	"github.com/IBM/sarama"
)

// ConsumerGroupHandler 实现了 sarama.ConsumerGroupHandler 接口
type consumerGroupHandler struct{}

func (consumerGroupHandler) Setup(_ sarama.ConsumerGroupSession) error   { return nil }
func (consumerGroupHandler) Cleanup(_ sarama.ConsumerGroupSession) error { return nil }

// ConsumeClaim 是实际处理消息的方法
func (h consumerGroupHandler) ConsumeClaim(session sarama.ConsumerGroupSession, claim sarama.ConsumerGroupClaim) error {
	for message := range claim.Messages() {
		log.Printf("Message claimed: topic=%s, partition=%d, offset=%d, key=%s, value=%s\n",
			message.Topic, message.Partition, message.Offset, string(message.Key), string(message.Value))
		// 标记消息为已处理，提交偏移量
		session.MarkMessage(message, "")
	}
	return nil
}

func main() {
	// 1. 创建消费者组配置
	config := sarama.NewConfig()
	config.Version = sarama.V2_8_1_0 // 根据你的 Kafka 版本调整
	config.Consumer.Group.Rebalance.Strategy = sarama.BalanceStrategyRange
	config.Consumer.Offsets.Initial = sarama.OffsetNewest // 从最新的偏移量开始消费

	// 2. 创建消费者组
	groupID := "test-go-group"
	brokers := []string{"192.168.0.185:9092"}
	topic := "test-topic"

	consumerGroup, err := sarama.NewConsumerGroup(brokers, groupID, config)
	if err != nil {
		log.Fatalf("Error creating consumer group: %v", err)
	}
	defer consumerGroup.Close()

	// 3. 使用上下文和信号量控制消费循环
	ctx, cancel := context.WithCancel(context.Background())
	wg := &sync.WaitGroup{}
	wg.Add(1)

	go func() {
		defer wg.Done()
		handler := consumerGroupHandler{}

		for {
			// 消费消息。Consume 方法会在重平衡时自动重新启动
			if err := consumerGroup.Consume(ctx, []string{topic}, handler); err != nil {
				log.Printf("Error from consumer: %v", err)
				// 可以根据错误类型决定是否退出
			}
			// 如果上下文被取消，则退出循环
			if ctx.Err() != nil {
				return
			}
		}
	}()

	log.Println("Sarama consumer up and running!...")

	// 4. 等待中断信号，优雅退出
	sigterm := make(chan os.Signal, 1)
	signal.Notify(sigterm, syscall.SIGINT, syscall.SIGTERM)
	<-sigterm
	cancel()
	wg.Wait()
	log.Println("Shutting down consumer...")
}
