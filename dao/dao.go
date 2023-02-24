package dao

import "go-task/conf"

type Dao struct {
}

func NewDao(c *conf.Config) *Dao {
	return &Dao{}
}
