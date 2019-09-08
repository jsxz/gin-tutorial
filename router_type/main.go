package main

/**
curl -X GET "http://localhost:8888/get"
curl -X POST "http://localhost:8888/post"
curl -X DELETE "http://localhost:8888/delete"
curl -X DELETE "http://localhost:8888/any"
curl -X GET "http://localhost:8888/any"
*/
import "github.com/gin-gonic/gin"

func main() {
	r := gin.Default()
	r.GET("/get", func(c *gin.Context) {
		c.String(200, "get")
	})
	r.POST("/post", func(c *gin.Context) {
		c.String(200, "post")
	})
	r.Handle("DELETE", "/delete", func(c *gin.Context) {
		c.String(200, "delete")
	})
	r.Any("/any", func(c *gin.Context) {
		c.String(200, "any")
	})

	r.Run(":8888")
}
