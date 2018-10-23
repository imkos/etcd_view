package main

import (
	"context"
	"flag"
	"fmt"
	"go.etcd.io/etcd/client"
	"log"
	"strings"
	"time"
)

const (
	ectdServer = "http://127.0.0.1:2379"
)

var (
	sKey string
)

// ReadEtcdConfig 读取etcd相关环境配置
func ReadEtcdConfig(key string) (string, bool) {
	if ectdServer == "" {
		return "", false
	}
	etcdServers := strings.Split(ectdServer, ",")
	cfg := client.Config{
		Endpoints:               etcdServers,
		Transport:               client.DefaultTransport,
		HeaderTimeoutPerRequest: time.Second,
	}

	c, err := client.New(cfg)
	if err != nil {
		log.Printf("New err:%v", err)
		return "", false
	}

	kapi := client.NewKeysAPI(c)
	resp, err := kapi.Get(context.Background(), key, nil)
	if err != nil {
		log.Printf("Kapi Get Key:%s err:%v", key, err)
		return "", false
	}
	log.Printf("Etcd Config Key:%s Value:%s", key, resp.Node.Value)
	return resp.Node.Value, true
}

func init() {
	flag.StringVar(&sKey, "key", "", "get values by etcd key!")
}

func main() {
	flag.Parse()
	fmt.Println(ReadEtcdConfig(sKey))
}
