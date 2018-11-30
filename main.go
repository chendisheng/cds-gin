package main

import (
	"cds-beego/util/config"
	"fmt"
	"github.com/gin-gonic/gin"
)

func main() {

	value1 := config.String("key1")
	fmt.Println("key1 value =",value1)

	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {

		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.Run() // listen and serve on 0.0.0.0:8080
}
