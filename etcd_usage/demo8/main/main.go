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
		putOp clientv3.Op
		getOp clientv3.Op
		opResp clientv3.OpResponse
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
	putOp=clientv3.OpPut("/cron/jobs/job8","")
	if opResp,err=kv.Do(context.TODO(),putOp);err!=nil{
		fmt.Println(err)
		return
	}
	fmt.Println("写入的是",opResp.Put().Header.Revision)

	getOp=clientv3.OpGet("/cron/jobs/job8")
	if opResp,err=kv.Do(context.TODO(),getOp);err!=nil{
		fmt.Println(err)
		return
	}
	fmt.Println("读到的是",opResp.Get().Kvs[0].ModRevision)
	fmt.Println("读到的是",opResp.Get().Kvs[0].Value)
	}
