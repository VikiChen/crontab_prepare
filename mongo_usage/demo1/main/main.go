package main

import (
	"github.com/mongodb/mongo-go-driver/mongo"
	"context"
	"github.com/mongodb/mongo-go-driver/mongo/clientopt"
	"time"
	"fmt"
)

func main() {
	var(
		client *mongo.Client
		err error
		datebase *mongo.Database
		collection *mongo.Collection
	)

	if client,err=mongo.Connect(context.TODO(),"mongodb://10.101.12.7:27017",clientopt.ConnectTimeout(5*time.Second));err!=nil{
		fmt.Println(err)
		return
	}
	datebase =client.Database("my_db")
	collection=datebase.Collection("my_collection")
	collection=collection
}
