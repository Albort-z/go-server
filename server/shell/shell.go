package shell

import (
	"fmt"
	"io"
	"log"
	"os"
	"os/exec"
)

func startShell() {
	cmd := exec.Command("ls", "-lah")
	out, err := cmd.CombinedOutput()
	if err != nil {
		log.Fatalf("cmd.Run() failed with %s\n", err)
	}
	os.NewFile()
	io.MultiReader()

	var stdout, stderr []byte
	var errStdout, errStderr error
	stdoutIn, _ := cmd.StdoutPipe()
	stderrIn, _ := cmd.StderrPipe()
	cmd.Start()

	go func() {
		stdout, errStdout = copyAndCapture(os.Stdout, stdoutIn)
	}()
	go func() {
		stderr, errStderr = copyAndCapture(os.Stderr, stderrIn)
	}()
	err := cmd.Wait()
	if err != nil {
		log.Fatalf("cmd.Run() failed with %s\n", err)
	}
	if errStdout != nil || errStderr != nil {
		log.Fatalf("failed to capture stdout or stderr\n")
	}
	outStr, errStr := string(stdout), string(stderr)
	fmt.Printf("\nout:\n%s\nerr:\n%s\n", outStr, errStr)
}

// 启动服务后，服务开始监听一个端口，并收听信号
// 编译新的服务二进制文件
// 将新服务二进制文件通过接口传入到旧服务
// 旧服务通过exec.Command拉起新服务，拉起时的参数中添加标识，使子进程知道自己是被优雅启动的
// https://grisha.org/blog/2014/06/03/graceful-restart-in-golang/
// 子进程初始化，知道自己是被优雅启动
// f := os.NewFile(3, "") 子进程直接继承3文件符
// 子进程通过信号通知父进程结束服务并自杀
// 父进程收到信号后优雅退出
