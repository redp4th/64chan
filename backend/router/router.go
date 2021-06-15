package router

import (
	"nmb/middleware"
	v1 "nmb/router/api/v1"

	"github.com/gin-gonic/gin"
)

func InitRouter() (router *gin.Engine) {
	router = gin.Default()
	router.Use(middleware.Cors())
	// router.Use(middleware.TLSHandler())

	apiv1 := router.Group("/api/v1")
	{
		apiv1.GET("/getcookie", v1.AssignCookie)
		apiv1.GET("/ws", v1.ServeWS)
		apiv1.POST("/channel", v1.CreateChannel)
		// apiv1.GET("/test", func(c *gin.Context) {
		// 	c.String(http.StatusOK, "Hello world!\n")
		// })
	}

	return
}
