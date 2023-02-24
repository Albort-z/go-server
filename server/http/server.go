package http

import (
	"context"
	"log"
	"net/http"
	"time"

	"go-task/conf"
	"go-task/core/service"
)

// 该文件为http服务的初始化文件，负责初始化接口

func init() {
}

// Server 服务器
type Server struct {
	c      *conf.Config
	svc    *service.Service
	server *http.Server
}

func NewServer(c *conf.Config, svc *service.Service) (*Server, error) {
	var s = Server{
		c:   c,
		svc: svc,
	}
	log.Println("初始化服务器")
	return &s, nil
}

// RunServer 启动http服务接口，传入配置和核心服务对象，通过调用核心服务对象的方法来对外提供服务
func (s *Server) RunServer() error {
	log.Println("运行服务器")

	s.server = &http.Server{
		Handler: setupHandler(s.svc),
		Addr:    s.c.Addr,
	}

	err := s.server.ListenAndServe()
	return err
}

func (s *Server) ShutdownServer() error {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()
	return s.server.Shutdown(ctx)
}
