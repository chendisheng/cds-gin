package routers

import (
	"cds-gin/middleware/jwt"
	"cds-gin/routers/api"
	"cds-gin/routers/api/v1"
	"github.com/gin-gonic/gin"
	"github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
)


//路由初始化
func InitRouter() *gin.Engine {
	r := gin.New()

	r.Use(gin.Logger())

	r.Use(gin.Recovery())

	gin.SetMode("debug")

	r.GET("/auth", api.GetAuth)
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

    contextPath := "/demo"
	apiv1 := r.Group(contextPath+"/api/v1")
	apiv1.Use(jwt.JWT())
	{

		////新建产品
		apiv1.POST("/product", v1.AddProduct)

	}

	return r
}