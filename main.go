package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	// 解决跨域问题
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

var pws *websocket.Conn

func main() {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong1312313a啊啊范德萨123aaa0.0.1",
		})

	})
	r.POST("/msg", msg)

	r.GET("/ws", func(c *gin.Context) {
		ws, err := upgrader.Upgrade(c.Writer, c.Request, nil)
		if err != nil {
			log.Print("upgrade:", err)
			return
		}
		pws = ws
		defer ws.Close()
		for {
			// err = ws.WriteMessage(1, []byte("message123"))
			// if err != nil {
			// 	log.Println("write:", err)
			// 	break
			// }
			time.Sleep(time.Second * 1)
		}
	})
	r.Run(":80") // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}

func msg(c *gin.Context) {
	jsonData, _ := ioutil.ReadAll(c.Request.Body)
	err := pws.WriteMessage(1, jsonData)
	if err != nil {
		log.Println("write:", err)

	}
	fmt.Println(string(jsonData))
}
