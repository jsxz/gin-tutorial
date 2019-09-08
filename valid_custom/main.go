package main

/**
curl -X POST -d "check_in=2019-09-21&check_out=2019-09-22" "http://localhost:8888/test"


*/
import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"gopkg.in/go-playground/validator.v8"
	"log"
	"net/http"
	"reflect"
	"time"
)

type Booking struct {
	CheckIn  time.Time `form:"check_in" binding:"required,bookabledate" time_format:"2006-01-02"`
	CheckOut time.Time `form:"check_out" binding:"required,gtfield=CheckIn" time_format:"2006-01-02"`
}

func main() {
	r := gin.Default()
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		v.RegisterValidation("bookabledate", bookableDate)
	}
	r.Any("test", testing)
	r.Run(":8888")
}
func testing(c *gin.Context) {
	var booking Booking
	if err := c.ShouldBind(&booking); err != nil {
		c.String(http.StatusBadRequest, "error: %v", err)
	} else {
		c.JSON(http.StatusOK, booking)
	}
}
func bookableDate(v *validator.Validate, topStruct reflect.Value, currentStructOrField reflect.Value, field reflect.Value, fieldType reflect.Type, fieldKind reflect.Kind, param string) bool {

	if date, ok := field.Interface().(time.Time); ok {
		today := time.Now()
		log.Printf("date:%v", date)
		log.Printf("today:%v", today)
		if date.Unix() > today.Unix() {
			return true
		}
	}

	return false
}
