package main

import (
	"go.etcd.io/etcd/clientv3"
	"time"
	"fmt"
)

func main()  {
	var (
		config clientv3.Config
		client *clientv3.Client
		err error
	)
	config=clientv3.Config{
		Endpoints:[]string{"10.101.12.7:2379"},
		DialTimeout:5*time.Second,
	}
	if client, err = clientv3.New(config);err!=nil{
		fmt.Println(err)
		return
	}
	client=client

}
