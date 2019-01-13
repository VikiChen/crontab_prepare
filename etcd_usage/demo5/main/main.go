package main

import (
	"go.etcd.io/etcd/clientv3"
	"time"
	"fmt"
	"context"
	"go.etcd.io/etcd/mvcc/mvccpb"
)

func main()  {
	var (
		config clientv3.Config
		client *clientv3.Client
		err error
		kv clientv3.KV
		delResp *clientv3.DeleteResponse
		idx int
		kvpair *mvccpb.KeyValue
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
	if delResp,err=kv.Delete(context.TODO(),"/cron/jobs/",clientv3.WithPrefix());err!=nil{
		fmt.Println(err)
		return
	}
	if len(delResp.PrevKvs)!=0{
		for idx ,kvpair =range delResp.PrevKvs{
			idx=idx
			fmt.Println("删除了",string(kvpair.Key),kvpair.Value)
		}
	}



}
