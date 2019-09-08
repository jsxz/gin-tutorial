package main

/**
curl curl http://localhost:8888/user/afadf

*/
import "github.com/gin-gonic/gin"

func main() {
	r := gin.Default()
	r.GET("/user/*action", func(c *gin.Context) {
		c.String(200, "hello")
	})
	r.Run(":8888")
}
