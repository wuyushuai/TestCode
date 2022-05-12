package main

import (
	"fmt"
	"io/ioutil"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong1312313a啊啊范德萨123aaa0.0.1",
		})

	})
	r.POST("/msg", msg)
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}

func msg(c *gin.Context) {
	jsonData, _ := ioutil.ReadAll(c.Request.Body)
	fmt.Println(string(jsonData))
}
