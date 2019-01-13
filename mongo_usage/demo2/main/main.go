package main

import (
	"github.com/mongodb/mongo-go-driver/mongo"
	"context"
	"github.com/mongodb/mongo-go-driver/mongo/clientopt"
	"time"
	"fmt"
	"github.com/mongodb/mongo-go-driver/bson/objectid"
)

type TimePoint struct {
	StartTime int64 `json:"startTime"`
	EndTime int64	`json:"endTime"`
}


type LogRecord struct {
	jobName string `bson:"jobName"`
	Command string `bson:"command"`
	Err string `bson:"err"`
	Conent string `bson:"conent"`
	TimePoint TimePoint `bson:"timePoint"`
}


func main() {
	var(
		client *mongo.Client
		err error
		datebase *mongo.Database
		collection *mongo.Collection
		record *LogRecord
		result *mongo.InsertOneResult
		docId objectid.ObjectID
	)

	if client,err=mongo.Connect(context.TODO(),"mongodb://10.101.12.7:27017",clientopt.ConnectTimeout(5*time.Second));err!=nil{
		fmt.Println(err)
		return
	}
	datebase =client.Database("cron")
	collection=datebase.Collection("log")
	record =&LogRecord{
		jobName:"job10",
		Command:"echo hello",
		Err:"",
		Conent:"hello",
		TimePoint:TimePoint{StartTime:time.Now().Unix(),EndTime:time.Now().Unix()+10},
	}
	if result,err=collection.InsertOne(context.TODO(),record);err!=nil{
		fmt.Println(result)
		return
	}
	docId=result.InsertedID.(objectid.ObjectID)
	fmt.Println("自增id",docId.Hex())
}
