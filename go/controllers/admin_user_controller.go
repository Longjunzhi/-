package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func AdminUserGet(c *gin.Context) {
	c.JSON(http.StatusOK, "")

}
