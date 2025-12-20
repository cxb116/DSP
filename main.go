package main

import (
	clientv3 "go.etcd.io/etcd/client/v3"
	"time"
)

func NewClient() (*clientv3.Client, error) {
	return clientv3.New(clientv3.Config{
		Endpoints:   []string{"117.72.67.186:2379"},
		DialTimeout: 5 * time.Second,
	})
}

func main() {

}
