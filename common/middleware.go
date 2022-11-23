package common

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func MiddleLog(c *gin.Context) {
	m := make(map[string]interface{})
	m["url"] = c.Request.URL.Path
	//logrus.Info("info msg")
}

func VerifyLogin(c *gin.Context) {
	cookie, _ := c.Cookie("user_cookie")
	if cookie == "" {
		println("not login")
		c.Redirect(http.StatusMovedPermanently, "/user/login")
		c.Abort()
	} else {
		println("already login")
		c.Next()
	}
}
