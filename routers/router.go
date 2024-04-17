package routers

import (
	"gin-api/api"

	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	apiv1 := r.Group("/api/v1")

	apiv1.GET("/ping", api.GetPing)

	apiv1.POST("/pong", api.AddPing)

	return r
}
