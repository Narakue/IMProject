package controller

import (
	"IMProject/model"
	"IMProject/service"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"net/http"
)

func Verify(c *gin.Context) {
	println("hello")
}

func Login(c *gin.Context) {
	user := model.User{}
	err := c.ShouldBindJSON(&user)
	if err != nil {
		logrus.Error("msg:", err)
		c.JSON(http.StatusOK, gin.H{"msg": err.Error()})
		return
	}
	username := user.Username
	password := user.Password
	result := ""
	if username != "" && password != "" {
		result = service.Login(username, password)
		if result == "success" {
			c.SetCookie("user_cookie", username, 1000, "/", "localhost", false, true)
		}
	} else {
		result = "fail"
	}
	c.JSON(http.StatusOK, gin.H{"msg": result})
}
