package main

import (
	"github.com/gin-gonic/gin"
	"webhook/lib"
	"log"
)

func main() {
	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()
	r.Handle("POST","/payload",lib.GetHooks)
	r.Handle("GET","/", func(c *gin.Context) {
		c.JSON(200,"welcome")
	})
	log.Println(r.Run(":8585"))
}
