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
		gerResp *clientv3.GetResponse
		watchStartRevision int64
		watcher clientv3.Watcher
		watchResponseChan  <-chan clientv3.WatchResponse
		watchResp clientv3.WatchResponse
		event *clientv3.Event
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

	go func() {
		for {
			kv.Put(context.TODO(),"/cron/jobs/job7","l an job7")
			kv.Delete(context.TODO(),"/cron/jobs/job7")
			time.Sleep(1*time.Second)
		}
	}()
	if gerResp,err=kv.Get(context.TODO(),"/cron/jobs/job7",clientv3.WithCountOnly());err!=nil{
		fmt.Println(err)
		return
	}
	if len(gerResp.Kvs)!=0{
		fmt.Println("当前值",string(gerResp.Kvs[0].Value))
	}
	//当前etcd事务id ，单调递增
	watchStartRevision=gerResp.Header.Revision+1
	//创建一个watch
	watcher=clientv3.NewWatcher(client)
	ctx,cancelFunc :=context.WithCancel(context.TODO())
	time.AfterFunc(5*time.Second, func() {
		cancelFunc()
	})
	watchResponseChan = watcher.Watch(ctx,"/cron/jobs/job7",clientv3.WithRev(watchStartRevision))

	for watchResp=range watchResponseChan{
		for _,event= range watchResp.Events{
			switch event.Type {
			case mvccpb.PUT:
				fmt.Println("修改为",string(event.Kv.Value),"Revision",event.Kv.CreateRevision,event.Kv.ModRevision)
			case mvccpb.DELETE:
				fmt.Println("删除了",event.Kv.ModRevision)
			}
		}
	}
	}
