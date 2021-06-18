package router

import (
	"net/http"
	"nmb/middleware"
	v1 "nmb/router/api/v1"

	"github.com/gin-gonic/gin"
)

func InitRouter() (router *gin.Engine) {
	router = gin.Default()
	router.Use(middleware.Cors())
	router.Use(middleware.TLSHandler())

	apiv1 := router.Group("/api/v1")
	{
		apiv1.GET("/getcookie", v1.AssignCookie)
		apiv1.GET("/ws", v1.ServeWS)
		apiv1.POST("/channel", v1.CreateChannel)
		apiv1.POST("/upload", v1.FileHandler)
		apiv1.POST("/pic", v1.PicHandler)
	}
	router.StaticFS("/static", http.Dir("/Users/redpath/work/nmb/storage"))
	return
}
