package main

import (
	"os/exec"
	"fmt"
)

func main() {
	var (
		cmd *exec.Cmd
		err error
		outPut []byte
	)

	// cmd = exec.Command("/bin/bash", "-c", "echo 1;echo2;")

	cmd = exec.Command("D:\\Git\\bin\\bash.exe", "-c", "echo hello")
	outPut, err = cmd.CombinedOutput()
	if err!=nil{
		fmt.Println(err)
		return
	}
	fmt.Println(string(outPut))

}
