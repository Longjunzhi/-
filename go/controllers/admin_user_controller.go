package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"pxj/CloudTravelShopApi/go/services"
)

func AdminUserGet(c *gin.Context) {
	req := &services.AdminUserGetRequest{}
	err := c.BindJSON(&req)
	if err != nil {
		c.JSON(http.StatusBadRequest, "参数错误!")
		return
	}
	resp, code, err := services.AdminUserGet(req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, "服务器异常")
		return
	}
	Success(c, code, resp)
}
