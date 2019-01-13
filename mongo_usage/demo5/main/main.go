package main

import (
	"github.com/mongodb/mongo-go-driver/mongo"
	"context"
	"github.com/mongodb/mongo-go-driver/mongo/clientopt"
	"time"
	"fmt"
)

// 任务的执行时间点
type TimePoint struct {
	StartTime int64	`bson:"startTime"`
	EndTime int64	`bson:"endTime"`
}

// 一条日志
type LogRecord struct {
	JobName string	`bson:"jobName"` // 任务名
	Command string `bson:"command"` // shell命令
	Err string `bson:"err"` // 脚本错误
	Content string `bson:"content"`// 脚本输出
	TimePoint TimePoint `bson:"timePoint"`// 执行时间点
}

type TimeBeforeCond struct {
	Before int64`bson:"$lt"`
} 

type DeleteCond struct {
	beforeCond TimeBeforeCond  `bson:"timePoint.startTime"`
}

func main() {
	var (
		client *mongo.Client
		err error
		database *mongo.Database
		collection *mongo.Collection
		delCond *DeleteCond
		delResult *mongo.DeleteResult
	)
	// 1, 建立连接
	if client, err = mongo.Connect(context.TODO(), "mongodb://36.111.184.221:27017", clientopt.ConnectTimeout(5 * time.Second)); err != nil {
		fmt.Println(err)
		return
	}

	// 2, 选择数据库my_db
	database = client.Database("cron")

	// 3, 选择表my_collection
	collection = database.Collection("log")

	delCond =&DeleteCond{
		beforeCond:TimeBeforeCond{
			Before:time.Now().Unix(),
		},
	}


	// 发起插入
	// 5, 查询（过滤 +翻页参数）
	if delResult, err = collection.DeleteMany(context.TODO(), delCond); err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(delResult.DeletedCount)
}
