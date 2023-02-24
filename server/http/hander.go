package http

import (
	"go-task/server/middleware"
	"net/http"

	"github.com/gin-gonic/gin"

	"go-task/core/service"
)

func setupHandler(svc *service.Service) http.Handler {
	h := &Handler{svc: svc}

	r := gin.Default()

	PublicGroup := r.Group("")
	// 公开接口

	PublicGroup.GET("/hi", h.sayHi)

	PrivateGroup := r.Group("")
	PrivateGroup.Use(middleware.LoadTls())
	// 私有接口

	return r
}

type Handler struct {
	svc *service.Service
}

func (h *Handler) sayHi(c *gin.Context) {
	resp, err := h.svc.SayHI(c)
	if err != nil {
		c.AbortWithStatusJSON(500, gin.H{"error": err.Error()})
	}
	c.JSON(200, resp)
}
