package main

import (
	"cds-gin/controllers"
	"cds-gin/util/config"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {

	value1 := config.String("key1")
	fmt.Println("key1 value =",value1)

	router := gin.Default()


	router.GET("/ping", func(c *gin.Context) {

		c.JSON(http.StatusOK, "pong")
	})

	v1 := router.Group("/v1")
	{
		v1.POST("/ping", func(c *gin.Context) {
			c.JSON(http.StatusOK, "pong")
		})
	}

	//注册路由
	controllers.RegProductRoute(router)

	// listen and serve on 0.0.0.0:8080
	router.Run()
}
