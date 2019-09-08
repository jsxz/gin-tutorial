package main

/**
curl -X POST -d "first_name=abc" "http://localhost:8888/test"


*/
import (
	"bytes"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"net/http"
)

func main() {
	r := gin.Default()
	r.Any("/test", func(c *gin.Context) {
		bodyBytes, err := ioutil.ReadAll(c.Request.Body)
		if err != nil {
			c.String(http.StatusBadRequest, err.Error())
			c.Abort()
		}
		// bodyBytes 放回requst.body ,下面才能取到值
		c.Request.Body = ioutil.NopCloser(bytes.NewBuffer(bodyBytes))
		firstName := c.PostForm("first_name")
		lastNmae := c.DefaultPostForm("last_name", "anjun")
		c.String(http.StatusOK, "%s,%s", firstName, lastNmae)
	})
	r.Run(":8888")
}
