package main

import (
	"github.com/gorhill/cronexpr"
	"time"
	"fmt"
)

type CronJob struct {
	expr *cronexpr.Expression
	nextTime time.Time
}

func main()  {
	//一个调度协程 定时检查cron  过期就执行
var(
	cronJob *CronJob
	expr *cronexpr.Expression
	now time.Time
	schedule map[string]*CronJob
)
  schedule =make(map[string]*CronJob)
   now =time.Now()
	expr = cronexpr.MustParse("*/5 * * * * * *")
    cronJob =&CronJob{
    	expr:expr,
    	nextTime:expr.Next(now),
	}
	schedule["job1"]=cronJob

	expr = cronexpr.MustParse("*/5 * * * * * *")
	cronJob =&CronJob{
		expr:expr,
		nextTime:expr.Next(now),
	}
	schedule["job2"]=cronJob

	go func() {
		var (
			jobName string
			cronJob *CronJob
			now time.Time
		)
		for {
			now =time.Now()
		    for jobName,cronJob =range schedule{
		    	if cronJob.nextTime.Before(now)||cronJob.nextTime.Equal(now){
		    		go func(jobName string) {
		    			fmt.Println("执行",jobName)
					}(jobName)
		    		//计算下次调度时间
		    		cronJob.nextTime=cronJob.expr.Next(now)
		    		fmt.Println("下次执行时间",cronJob.nextTime)
				}
			}

			select {
			case <-time.NewTimer(100*time.Millisecond).C:
				
			}

			}
	}()
    time.Sleep(100*time.Second)
}
