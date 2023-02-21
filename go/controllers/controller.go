package controllers

import (
	"github.com/gin-gonic/gin"
	"pxj/CloudTravelShopApi/go/services"
)

func Success(c *gin.Context, code int, data *services.Response) {

	c.JSON(code, data)
}
func Fail(c *gin.Context, code int, data *services.Response) {

	c.JSON(code, data)
}

func Json(c *gin.Context, code int, data *services.Response) {

	c.JSON(code, data)
}

func Response(c *gin.Context, code int, data *services.Response) {

	c.JSON(code, data)
}
