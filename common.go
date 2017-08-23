package frango

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"

	"github.com/gin-gonic/gin"
)

func ParseRequestBody(c *gin.Context) map[string]interface{} {
	body, _ := ioutil.ReadAll(c.Request.Body)
	var data map[string]interface{}

	if err := json.Unmarshal(body, &data); err != nil {
		return data
	}

	return data
}

func LogErr(err error) {
	if err != nil {
		log.Println(err.Error())
	}
}

func PrintErr(err error) {
	if err != nil {
		fmt.Println(err.Error())
	}
}
