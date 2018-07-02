package lib

import (
	"github.com/gin-gonic/gin"
	"fmt"
)

type Payload struct {
	_Payload string `json:"payload"`
}

func GetHooks(c *gin.Context) {
	payload := Payload{}
	c.BindJSON(&payload)
	fmt.Println(payload)

}
