package controllers

import (
	"context"
	"github.com/gin-gonic/gin"
	"net/http"
	"pxj/CloudTravelShopApi/go/services"
)

func ApiLogin(c *gin.Context) {
	c.JSON(http.StatusOK, "")
}

func AdminLogin(c *gin.Context) {
	req := &services.LoginByPasswordRequest{}
	ctx := context.Background()
	err := c.BindJSON(&req)
	if err != nil {
		c.JSON(http.StatusBadRequest, "参数错误!")
		return
	}
	resp, code, err := services.LoginByPassword(ctx, req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, "服务器异常")
		return
	}
	c.JSON(code, resp)
}
