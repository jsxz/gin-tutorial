package main

/**
curl "http://localhost:8888/test?first_name=a"

*/
import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	r := gin.Default()
	r.GET("/test", func(c *gin.Context) {
		firstName := c.Query("first_name")
		lastNmae := c.DefaultQuery("last_name", "anjun")
		c.String(http.StatusOK, "%s,%s", firstName, lastNmae)
	})
	r.Run(":8888")
}
