package main

/**
curl http://localhost:8888/assets/a.html
curl http://localhost:8888/static/b.html
*/
import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	r := gin.Default()
	r.Static("/assets", "./assets")
	r.StaticFS("/static", http.Dir("static"))
	r.StaticFile("/favicon.ico", ".favicon.ico")
	r.Run(":8888")
}
