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
		lease clientv3.Lease
		leaseGrantResp *clientv3.LeaseGrantResponse
		leaseId clientv3.LeaseID
		keepRespChan <- chan *clientv3.LeaseKeepAliveResponse
		keepResp *clientv3.LeaseKeepAliveResponse
		ctx context.Context
		cancenlFunc context.CancelFunc
		tex clientv3.Txn
		txnResp *clientv3.TxnResponse
	)
	config=clientv3.Config{
		Endpoints:[]string{"10.101.12.7:2379"},
		DialTimeout:5*time.Second,
	}
	if client, err = clientv3.New(config);err!=nil{
		fmt.Println(err)
		return
	}

	//申请一个lease 租约
	lease=clientv3.NewLease(client)
	if leaseGrantResp,err=lease.Grant(context.TODO(),5);err!=nil {
		fmt.Println(err)
		return
	}
	leaseId=leaseGrantResp.ID

	ctx,cancenlFunc=context.WithCancel(context.TODO())
	defer cancenlFunc()
	defer lease.Revoke(context.TODO(),leaseId)

	if keepRespChan ,err=lease.KeepAlive(ctx,leaseId);err !=nil{
		fmt.Println(err)
		return
	}

	go func() {
		for{
			select{
			case keepResp=<-keepRespChan:
				if keepRespChan ==nil{
					fmt.Println("租约失效了")
					goto END
				}else {
					fmt.Println("收到自动续租应答",keepResp.ID)
				}
			}
		}
	END:
	}()

	//抢key  if 不存在 then 设置 else失败
	kv=clientv3.NewKV(client)
	tex =kv.Txn(context.TODO())
	//定义事务
	//key不存在
	tex.If(clientv3.Compare(clientv3.CreateRevision("/cron/locak/job99"),"=",0)).
		Then(clientv3.OpPut("/cron/locak/job99","xx",clientv3.WithLease(leaseId))).
		Else(clientv3.OpGet("/cron/locak/job99"))//否则抢锁事变

	if txnResp,err =tex.Commit();err!=nil{
		fmt.Println(err)
		return
	}

	if !txnResp.Succeeded{
		fmt.Println("被占用",string(txnResp.Responses[0].GetResponseRange().Kvs[0].Value))
		return
	}

	fmt.Println("处理任务")
	time.Sleep(5*time.Second)
	//putOp=clientv3.OpPut("/cron/jobs/job8","")
	//if opResp,err=kv.Do(context.TODO(),putOp);err!=nil{
	//	fmt.Println(err)
	//	return
	//}
	//fmt.Println("写入的是",opResp.Put().Header.Revision)
	//
	//getOp=clientv3.OpGet("/cron/jobs/job8")
	//if opResp,err=kv.Do(context.TODO(),getOp);err!=nil{
	//	fmt.Println(err)
	//	return
	//}
	//fmt.Println("读到的是",opResp.Get().Kvs[0].ModRevision)
	//fmt.Println("读到的是",opResp.Get().Kvs[0].Value)
	}
