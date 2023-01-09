package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"pxj/CloudTravelShopApi/go/services"
)

func AdminUserLoginByAccount(c *gin.Context) {
	req := &services.AdminUserLoginByAccountRequest{}
	err := c.BindJSON(&req)
	if err != nil {
		c.JSON(http.StatusBadRequest, "参数错误!")
		return
	}
	resp, code, err := services.AdminUserLoginByAccount(req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, "服务器异常")
		return
	}
	c.JSON(code, resp)
}

func CurrentUser(c *gin.Context) {
	req := &services.CurrentUserRequest{}
	resp, code, err := services.CurrentUser(req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, "服务器异常")
		return
	}
	c.JSON(code, resp)
}

func Rule(c *gin.Context) {

	c.JSON(http.StatusUnauthorized, "未授权")
}

func OutLogin(c *gin.Context) {
	c.JSON(http.StatusOK, "退出登录")
}
