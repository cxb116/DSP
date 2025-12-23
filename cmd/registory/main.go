package main

import (
	"context"
	"fmt"
	"github.com/cxb116/DSP/global"
	"github.com/cxb116/DSP/internal/client"
	clientv3 "go.etcd.io/etcd/client/v3"
)

func main() {

	global.EngineViper = client.NewClientViper()
	global.EngineDB = client.NewClientMysql()
	global.EngineRedis = client.NewRedisClient()

	//client.InitKafkaAsyncProducer(global.EngineConfig.Kafka.Brokers)

	client, err := client.NewEtcdClient()
	if err != nil {
		fmt.Println("链接异常 err:", err.Error())
		return
	}

	resp, err := client.Get(context.Background(), "/dsp/key", clientv3.WithPrefix())
	if err != nil {
		fmt.Println("Get err:", err.Error())
	}

	config := make(map[string]string)
	for _, kv := range resp.Kvs {
		config[string(kv.Key)] = string(kv.Value)
	}

}
