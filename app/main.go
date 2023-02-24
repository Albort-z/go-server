package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"

	"go-task/conf"
	"go-task/core"
)

// 在main中，只关心这个app是否成功启动，不关心app内是什么。
//所以main里的逻辑：加载配置、初始化app、启动app、监听信号、收到停止信号或者reload信号时，分别触发app的停止和reload、退出。

func main() {
	// 加载配置
	c, err := conf.ReadConfig()
	if err != nil {
		log.Fatal("读取配置失败, err:", err)
	}

	// 初始化app
	ser, err := core.NewCore(c)
	if err != nil {
		log.Fatal("初始化服务失败, err:", err)
	}
	log.Println("服务已启动")

	// 启动app
	defer func() {
		// 终止app
		ser.Stop()
		log.Println("服务已退出")
	}()

	// 监听信号
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGHUP, syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT)
	for {
		s := <-sigChan
		log.Printf("系统信号:%s\n", s.String())
		switch s {
		case syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT:
			return
		case syscall.SIGHUP:
		default:
		}
	}
}
