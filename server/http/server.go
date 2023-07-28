package http

import (
	"context"
	"log"
	"net/http"
	"time"

	"go-task/conf"
	"go-task/core/service"
)

// Server 服务器
type Server struct {
	c      *conf.Config
	svc    *service.Service
	server *http.Server
	// 证书与密钥文件路径
	certFile, keyFile string
}

func NewServer(c *conf.Config, svc *service.Service) (*Server, error) {
	var s = Server{
		c:        c,
		svc:      svc,
		certFile: c.CertFile,
		keyFile:  c.KeyFile,
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

	var err error
	if s.certFile == "" {
		err = s.server.ListenAndServe()
	} else {
		err = s.server.ListenAndServeTLS(s.certFile, s.keyFile)
	}
	return err
}

func (s *Server) ShutdownServer() error {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()
	return s.server.Shutdown(ctx)
}
