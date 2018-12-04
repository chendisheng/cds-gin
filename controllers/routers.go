package controllers

import (
	"cds-gin/model"
	"github.com/gin-gonic/gin"
	"net/http"
)

func RegProductRoute(r *gin.Engine){

	  r.POST("/product" , func(c *gin.Context) {
		  var product model.Product
		  if err:= c.ShouldBind(&product); err !=nil {
			  c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			  return
		  }

		  if product.Code == "" || product.Price == 0 {
			  c.JSON(http.StatusBadRequest, gin.H{"errr": "param invalid"})
			  return
		  }

		  if err := model.ProductDao.Insert(product); err != nil{
			  c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			  return
		  }
		  c.JSON(http.StatusOK,gin.H{"status":"ok"})

	  })
}
