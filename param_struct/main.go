package main

/**
curl -X POST -d "name=abc&address=gg&birthday=2019-08-09" "http://localhost:8888/test"


*/
import (
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

type Person struct {
	Name     string    `form:"name"`
	Address  string    `form:"address"`
	Birthday time.Time `form:"birthday" time_format:"2006-01-02"`
}

func main() {
	r := gin.Default()
	r.Any("test", testing)
	r.Run(":8888")
}
func testing(c *gin.Context) {
	var persion Person
	if err := c.ShouldBind(&persion); err != nil {
		c.String(http.StatusBadRequest, "%v", err)
	} else {
		c.JSON(http.StatusOK, persion)
	}
}
