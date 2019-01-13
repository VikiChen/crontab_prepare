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
		putRes *clientv3.PutResponse
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
	if putRes,err=kv.Put(context.TODO(),"/cron/jobs/job1","bye",clientv3.WithPrevKV());err!=nil{
		fmt.Println(err)
	}else {
		fmt.Println("Revision",putRes.Header.Revision)
		if putRes.PrevKv!=nil{
			fmt.Println("prevvalue",string(putRes.PrevKv.Value))
		}
	}

}
