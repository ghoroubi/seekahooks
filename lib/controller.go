package lib

import (
	"github.com/gin-gonic/gin"
	"fmt"
	"io/ioutil"
	"log"
)

type Payload struct {
	_Payload string `json:"payload"`
}

func GetHooks(c *gin.Context) {
	payload := Payload{}
	c.BindJSON(&payload)

	body,err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		log.Fatal(err)
		return
	}


	fmt.Println(string(body))

}
