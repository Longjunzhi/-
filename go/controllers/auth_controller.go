package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"pxj/CloudTravelShopApi/go/constant"
	"pxj/CloudTravelShopApi/go/services"
	"strconv"
)

// AdminUserLoginByAccount
// @Tags         adminLogin
// @Accept       json
// @Success      200  {object} services.AdminUserLoginByAccountResponse
// @Router       /admin/login/account [post]
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
	if adminUserId, ok := c.Get(constant.CONTEXT_KEY_ADMIN_USER_ID); ok {
		adminUserIdInt, _ := strconv.Atoi(adminUserId.(string))
		req.AdminUserId = uint(adminUserIdInt)
	}
	resp, code, err := services.CurrentUser(req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, "服务器异常")
		return
	}
	c.JSON(code, resp)
}

func Rule(c *gin.Context) {

	c.JSON(http.StatusOK, "")
	//c.JSON(http.StatusUnauthorized, "未授权")
}

func OutLogin(c *gin.Context) {
	c.JSON(http.StatusOK, "退出登录")
}
