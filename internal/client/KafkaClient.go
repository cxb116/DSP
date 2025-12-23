package client

import (
	"fmt"
	"github.com/IBM/sarama"
	"github.com/cxb116/DSP/global"
	"log"
	"time"
)

var AsyncProducer sarama.AsyncProducer

func InitKafkaAsyncProducer(brokers []string) error {
	cfg := sarama.NewConfig()
	fmt.Println(global.EngineConfig.Kafka)

	// Kafka 最大性能配置
	cfg.Producer.RequiredAcks = sarama.WaitForLocal                                       // 本地确认即可
	cfg.Producer.Compression = sarama.CompressionSnappy                                   // 压缩吞吐高
	cfg.Producer.Flush.Frequency = global.EngineConfig.Kafka.Frequency * time.Millisecond // 批量发送
	cfg.Producer.Flush.Bytes = 256 * 1024
	cfg.Producer.Flush.Messages = global.EngineConfig.Kafka.Messages // 批次 1000 条
	cfg.Producer.Return.Successes = global.EngineConfig.Kafka.ReturnSuccess

	cfg.Version = sarama.MaxVersion // kafka最大版本

	var err error
	AsyncProducer, err = sarama.NewAsyncProducer(brokers, cfg)
	if err != nil {
		fmt.Println("Error creating the async producer:", err)
		return err
	}

	// 错误处理（异步）
	go func() {
		for err := range AsyncProducer.Errors() {
			log.Println("Kafka async error:", err)
		}
	}()

	fmt.Println("Kafka async producer started")

	return nil
}

// 发送自定义topic
func KafkaAsyncSend(topic, key string, msg string) {
	AsyncProducer.Input() <- &sarama.ProducerMessage{
		Topic: topic,
		Key:   sarama.ByteEncoder(key),
		Value: sarama.ByteEncoder(msg),
	}
}

// ssp 将请求响应 dsp 请求响应 超时,都去放到这个topic中
func TopicSspKafkaAsyncSend(key string, msg string) { // key是  ssp:媒体广告位id:预算广告位id:时间戳
	AsyncProducer.Input() <- &sarama.ProducerMessage{
		Topic: global.EngineConfig.Kafka.TopicSsp,
		Key:   sarama.ByteEncoder(key),
		Value: sarama.ByteEncoder(msg),
	}
}

// 统计请求的次数的topic
func TopicSspReqKafkaAsyncSend(key string, msg string) { // key是  ssp:媒体广告位id:预算广告位id:时间戳
	AsyncProducer.Input() <- &sarama.ProducerMessage{
		Topic: global.EngineConfig.Kafka.TopicSsp,
		Key:   sarama.ByteEncoder(key),
		Value: sarama.ByteEncoder(msg),
	}
}
