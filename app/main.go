package main

import (
	"github.com/gin-gonic/gin"
	"webhook/lib"
	"log"
)

func main() {
	r := gin.Default()
	r.Handle("POST","/payload",lib.GetHooks)
	log.Fatal(r.Run(":8585"))
}
