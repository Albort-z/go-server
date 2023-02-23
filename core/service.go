package core

import (
	"context"
	"log"
	"math/rand"
	"sync"
	"sync/atomic"
	"time"

	"go-task/conf"
)

type Service struct {
	cancel     context.CancelFunc
	inShutdown atomic.Bool // true when server is in shutdown
	dealer     sync.WaitGroup
}

func NewService(c *conf.Config) (*Service, error) {
	s := &Service{}
	s.Start()
	return s, nil
}

// Start 启动服务
func (s *Service) Start() {
	var ctx context.Context
	ctx, s.cancel = context.WithCancel(context.Background())

	var err error
	var errChan = make(chan error)
	go s.run(errChan)

	go func() {
		select {
		case <-ctx.Done():
			// 收到信号退出
		case err = <-errChan:
			log.Fatal("服务执行异常退出, err:", err)
		}
	}()
}

// Stop 停止服务
func (s *Service) Stop() {
	// 先关闭接口
	s.inShutdown.Store(true)
	s.dealer.Wait()

	// 再停止服务
	s.cancel()
	// 最后释放资源
}

// Reload 平滑重启服务
func (s *Service) Reload(c *conf.Config) error {
	return nil
}

// run 服务运行时
func (s *Service) run(errChan chan error) {
	log.Println("服务启动中...")
	defer func() {
		log.Println("服务退出中...")
	}()

	for true {
		time.Sleep(time.Second * 20)
		if s.inShutdown.Load() {
			break
		}
		// 服务没有在关闭，就可以继续申请请求
		reqID := rand.Int()
		s.dealer.Add(1)
		go func(qid int) {
			defer func() {
				s.dealer.Done()
			}()
			log.Println("有新的请求接入:", qid)
			defer func() {
				log.Println("请求处理完成:", qid)
			}()
			time.Sleep(time.Second * 3)
		}(reqID)

	}
}
