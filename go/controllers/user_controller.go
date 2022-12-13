package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func ApiLogin(c *gin.Context) {
	c.JSON(http.StatusOK, "")
}

func AdminLogin(c *gin.Context) {
	c.JSON(http.StatusOK, map[string]any{
		"currentAuthority": [1]string{"超级管理员"},
		"status":           "ok",
		"token":            "124",
		"type":             "account",
	})
}
