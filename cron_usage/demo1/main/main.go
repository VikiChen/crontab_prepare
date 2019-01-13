package main

import (
	"github.com/gorhill/cronexpr"
	"fmt"
	"time"
)

func main()  {
	var (
		expr *cronexpr.Expression
		err error
		now time.Time
		nextTime time.Time
	)
	if expr, err = cronexpr.Parse("* * * * *");err!=nil{
		fmt.Println(err)
		return
	}
	if expr,err=cronexpr.Parse("*/3 * * * * * *");err!=nil{
		fmt.Println(err)
		return
	}
	now =time.Now()
	nextTime=expr.Next(now)
	time.AfterFunc(nextTime.Sub(now), func() {
		fmt.Println("gogogo",nextTime)
	})
	time.Sleep(5*time.Second)

	}
