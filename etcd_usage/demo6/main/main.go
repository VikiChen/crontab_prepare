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
		lease clientv3.Lease
		leaseGrantResp *clientv3.LeaseGrantResponse
		leaseId clientv3.LeaseID
		putResp *clientv3.PutResponse
		kv clientv3.KV
		getResp *clientv3.GetResponse
		keepResp *clientv3.LeaseKeepAliveResponse
		keepRespChan <-chan *clientv3.LeaseKeepAliveResponse

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
	if leaseGrantResp,err=lease.Grant(context.TODO(),10);err!=nil {
		fmt.Println(err)
		return
	}
	leaseId=leaseGrantResp.ID
    //或者kv 对象
    kv=clientv3.NewKV(client)
    if putResp,err=kv.Put(context.TODO(),"/cron/lock/job6","",clientv3.WithLease(leaseId));err!=nil{
    	fmt.Println(err)
    	return
	}
	fmt.Println("success",putResp.Header.Revision)

    ctx,_:=context.WithTimeout(context.TODO(),5*time.Second)
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

    for {
		if	getResp,err= kv.Get(context.TODO(),"/cron/lock/job6"); err != nil {
			fmt.Println(err)
			return
		}
		if getResp.Count==0{
			fmt.Println("过期了")
			break
		}
		fmt.Println("还没过期",getResp.Kvs)
		time.Sleep(2 *time.Second)
	}
}
