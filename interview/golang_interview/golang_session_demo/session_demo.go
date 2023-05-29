package main

import (
	"github.com/gin-gonic/gin"
	"github.com/gorilla/sessions"
)

var store = sessions.NewCookieStore([]byte("secret"))

func main() {
	r := gin.Default()

	r.GET("/login", func(c *gin.Context) {
		session, _ := store.Get(c.Request, "session-name")
		session.Values["authenticated"] = true
		session.Save(c.Request, c.Writer)
		c.JSON(200, gin.H{"message": "login successful"})
	})

	r.GET("/logout", func(c *gin.Context) {
		session, _ := store.Get(c.Request, "session-name")
		session.Values["authenticated"] = false
		session.Save(c.Request, c.Writer)
		c.JSON(200, gin.H{"message": "logout successful"})
	})

	r.GET("/protected", func(c *gin.Context) {
		session, _ := store.Get(c.Request, "session-name")
		if auth, ok := session.Values["authenticated"].(bool); !ok || !auth {
			c.AbortWithStatus(401)
			return
		}
		c.JSON(200, gin.H{"message": "protected content"})
	})

	r.Run(":8080")
}

// 我们使用gorilla/sessions库来管理session。我们首先创建一个session存储对象，
//然后在登录和注销处理程序中设置和清除session值。在受保护的处理程序中，我们检查session值是否已经设置，如果没有设置，则返回401未授权状态码。
