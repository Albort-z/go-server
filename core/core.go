package core

import (
	"context"
	"log"

	"go-task/conf"
	"go-task/core/service"
	"go-task/dao"
	"go-task/server/http"
)

// Core 内核
type Core struct {
	cancel context.CancelFunc
}

func NewCore(c *conf.Config) (*Core, error) {
	s := &Core{}
	s.Start(c)
	return s, nil
}

// Start 启动服务
func (s *Core) Start(c *conf.Config) {
	svc := service.NewService(c, dao.NewDao(c))

	var ctx context.Context
	ctx, s.cancel = context.WithCancel(context.Background())

	var err error
	var errChan = make(chan error)

	// 启动服务器
	server, err := http.NewServer(c, svc)
	if err != nil {
		errChan <- err
		return
	}

	go s.run(server, errChan)

	go func() {
		select {
		case <-ctx.Done():
			// 收到信号退出
			err = server.ShutdownServer()
			if err != nil {
				log.Println("服务退出异常， err:", err)
			}
		case err = <-errChan:
			log.Fatal("服务执行异常退出, err:", err)
		}
	}()
}

// Stop 停止服务
func (s *Core) Stop() {
	// 再停止服务
	s.cancel()
	// 最后释放资源
}

// Reload 平滑重启服务
func (s *Core) Reload(c *conf.Config) error {
	// 使用配置创建新服务
	// 使用新服务
	return nil
}

// run 服务运行时
func (s *Core) run(server *http.Server, errChan chan error) {
	log.Println("服务启动中...")
	defer func() {
		log.Println("服务退出中...")
	}()

	err := server.RunServer()
	if err != nil {
		errChan <- err
		return
	}
}
