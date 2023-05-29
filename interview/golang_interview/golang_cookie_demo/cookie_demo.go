package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.GET("/set-cookie", func(c *gin.Context) {
		c.SetCookie("username", "john", 3600, "/", "localhost", false, true)
		c.JSON(200, gin.H{"message": "cookie set successfully"})
	})

	r.GET("/get-cookie", func(c *gin.Context) {
		username, err := c.Cookie("username")
		if err != nil {
			c.JSON(400, gin.H{"message": "cookie not found"})
			return
		}
		c.JSON(200, gin.H{"message": "cookie found", "username": username})
	})

	r.Run(":8080")
}

//在这个示例中，我们使用 SetCookie 方法来设置一个名为 "username" 的cookie，然后使用 JSON 方法返回一个成功消息。
//在另一个处理程序中，我们使用 Cookie 方法来获取名为 "username" 的cookie，并将其返回给客户端。如果cookie不存在，则返回一个错误消息。
//
//SetCookie 方法接受以下参数：
//
//name：cookie的名称。
//value：cookie的值。
//maxAge：cookie的最大存活时间（以秒为单位）。
//path：cookie的路径。
//domain：cookie的域名。
//secure：cookie是否只能通过HTTPS传输。
//httpOnly：cookie是否只能通过HTTP传输。
//Cookie 方法接受一个参数 name，表示要获取的cookie的名称。如果cookie不存在，则返回一个错误。
