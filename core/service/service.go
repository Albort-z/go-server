package service

import (
	"context"
	"log"

	"go-task/conf"
	"go-task/dao"
)

type Service struct {
	c   *conf.Config
	dao *dao.Dao
}

func NewService(c *conf.Config, dao *dao.Dao) *Service {
	return &Service{c: c, dao: dao}
}

func (s *Service) SayHI(ctx context.Context) (string, error) {
	log.Println("HI")
	return "你好", nil
}
