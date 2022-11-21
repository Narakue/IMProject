package controller

import (
	"Test/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Verify(c *gin.Context) {
	println("hello")
}

func Login(c *gin.Context) {
	username := c.Query("username")
	password := c.Query("password")
	if username != "" && password != "" {
		result := service.Login(username, password)
		if result == "success" {
			c.SetCookie("user_cookie", username, 1000, "/", "localhost", false, true)
		}
		c.JSON(http.StatusOK, gin.H{"result": result})
	}
}
