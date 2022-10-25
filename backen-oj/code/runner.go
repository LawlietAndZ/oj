package main

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"os/exec"
)

func main() {
	// go run code-user/main.go
	cmd := exec.Command("go", "run", "code-user/main.go")
	var out, stderr bytes.Buffer
	//错误信息
	cmd.Stderr = &stderr
	//输出信息
	cmd.Stdout = &out

	//管道，可用于写入数据
	stdinPipe, err := cmd.StdinPipe()
	if err != nil {
		fmt.Println(err.Error(),stderr.String())
	}
	//写入数据，计算两数之和
	io.WriteString(stdinPipe, "23 11\n")
	if err := cmd.Run(); err != nil {
		log.Fatalln(err, stderr.String())
	}
	fmt.Println(out.String())

}
