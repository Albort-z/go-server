package http

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"

	"go-task/core/service"
)

func setupHandler(svc *service.Service) http.Handler {
	h := &Handler{svc: svc}

	r := gin.Default()
	r.MaxMultipartMemory = 100 << 20

	// 公开接口
	PublicGroup := r.Group("")
	PublicGroup.Static("web", "web")
	PublicGroup.GET("/hi", h.sayHi)

	// 私有接口
	PrivateGroup := r.Group("")
	// PrivateGroup.Use(gin.BasicAuth(map[string]string{"fuck": "123"}))
	PrivateGroup.Static("data", "data")
	PrivateGroup.GET("fuck", h.sayHi)

	PrivateGroup.POST("upload", h.upload)
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

func (h *Handler) upload(c *gin.Context) {
	// 单文件
	file, err := c.FormFile("file")
	if err != nil {
		log.Println("未取到文件:", err.Error())
		c.String(400, "Failed")
		return
	}
	log.Println(file.Filename)

	dst := "./" + file.Filename
	// 上传文件至指定的完整文件路径
	c.SaveUploadedFile(file, dst)

	c.String(http.StatusOK, fmt.Sprintf("'%s' uploaded!", file.Filename))
}
