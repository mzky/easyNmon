package utils

import (
	"bytes"
	"context"
	"os/exec"
	"time"
)

//仅执行命令，不要用这个函数执行脚本
func Exec(command string) (string, error) {
	//设置超时时间 默认60秒
	ctxt, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	defer cancel()

	//函数返回一个*Cmd，用于使用给出的参数执行name指定的程序
	cmd := exec.CommandContext(ctxt, "/bin/bash", "-c", command)

	//读取io.Writer类型的cmd.Stdout，再通过bytes.Buffer(缓冲byte类型的缓冲器)将byte类型转化为string类型(out.String():这是bytes类型提供的接口)
	var stdout bytes.Buffer
	var stderr bytes.Buffer

	cmd.Stdout = &stderr
	cmd.Stdout = &stdout

	//Run执行c包含的命令，并阻塞直到完成. 这里stdout被取出，cmd.Wait()无法正确获取stdin,stdout,stderr，则阻塞在那了
	if err := cmd.Start(); err != nil {
		return stdout.String(), err
	}

	if err := cmd.Wait(); err != nil {
		return stdout.String(), err
	}

	return stdout.String(), nil
}
