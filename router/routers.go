package router

import (
	"IMProject/controller"
	"github.com/gin-gonic/gin"
)

func CommonRouter(r *gin.Engine) *gin.Engine {
	r1 := r.Group("/")
	{
		r1.GET("/", controller.Verify)
	}
	return r
}

func UserRouter(r *gin.Engine) *gin.Engine {
	r1 := r.Group("/user")
	{
		r1.POST("/login", controller.Login)
	}
	return r
}
