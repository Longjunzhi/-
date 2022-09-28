package middlewares

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

func Cors() gin.HandlerFunc {
	return func(c *gin.Context) {
		fmt.Println("cors 中间件")
		c.Next()
	}
}
