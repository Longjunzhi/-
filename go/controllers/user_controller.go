package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func ApiLogin(c *gin.Context) {
	c.JSON(http.StatusOK, "")
}
