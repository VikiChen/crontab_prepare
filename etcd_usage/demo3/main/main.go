package main

import (
	"go.etcd.io/etcd/clientv3"
	"time"
	"fmt"
	"context"
)

func main()  {
	var (
		config clientv3.Config
		client *clientv3.Client
		err error
		kv clientv3.KV
		gerResp *clientv3.GetResponse
	)
	config=clientv3.Config{
		Endpoints:[]string{"10.101.12.7:2379"},
		DialTimeout:5*time.Second,
	}
	if client, err = clientv3.New(config);err!=nil{
		fmt.Println(err)
		return
	}
	kv=clientv3.NewKV(client)
	if gerResp,err=kv.Get(context.TODO(),"/cron/jobs/job1",clientv3.WithCountOnly());err!=nil{
		fmt.Println(err)
	}else {
		fmt.Println("Revision",gerResp.Kvs)
		fmt.Println("Revision",gerResp.Count)

	}

}
