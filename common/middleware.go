package common

import "github.com/gin-gonic/gin"

func Middleware(c *gin.Context) {
	cookie, _ := c.Cookie("user_cookie")
	if cookie == "" {
		println("not login")
		c.Abort()
	} else {
		println("already login")
		c.Next()
	}
}