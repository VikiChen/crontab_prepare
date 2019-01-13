package main

import (
	"os/exec"
	"context"
	"time"
	"fmt"
)

type result struct {
	err error
	output []byte
}


func main() {

	var (
		ctx context.Context
		cancelFunc context.CancelFunc
		cmd *exec.Cmd
		resultChan chan *result
		res *result
	)
	resultChan =make(chan *result,1000)
	ctx, cancelFunc = context.WithCancel(context.TODO())
	go func() {
		var (
			output []byte
			err error
		)
		cmd =exec.CommandContext(ctx,"D:\\Git\\bin\\bash.exe", "-c", "sleep 4;echo hello")

		output,err=cmd.CombinedOutput()
		resultChan <-&result{output:output,err:err}
	}()


	time.Sleep(1*time.Second)
	cancelFunc() //关闭上下文 中断任务

	res =<-resultChan

	// 打印任务执行结果
	fmt.Println(res.err, string(res.output))

}
