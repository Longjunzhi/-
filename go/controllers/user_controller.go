package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// ApiLogin
// @Summary      Show an account
// @Description  get string by ID
// @Tags         login
// @Accept       json
// @Param        id   path      int  true  "Account ID"
// @Success      200  {string} welcome
// @Router       /api/login [post]
func ApiLogin(c *gin.Context) {
	c.JSON(http.StatusOK, "")
}
