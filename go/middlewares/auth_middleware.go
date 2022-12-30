package middlewares

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"net/http"
	"pxj/CloudTravelShopApi/go/constant"
	"pxj/CloudTravelShopApi/go/utils"
	"strings"
)

func Auth() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.GetHeader("Authorization")
		if token == "" || !strings.HasPrefix(token, "Bearer ") {
			c.JSON(http.StatusUnauthorized, gin.H{"code": 401, "message": "没有权限"})
			c.Abort()
			return
		}
		validToken, err := utils.ValidToken(token[7:])
		if err != nil {
			logrus.Errorf("validateTokentoken fail: %v", err)
			c.JSON(http.StatusUnauthorized, gin.H{"code": 401, "message": "没有权限"})
			c.Abort()
			return
		}
		userId := validToken.UserId
		ctx := c.Request.Context()
		ctx = context.WithValue(ctx, constant.CONTEXT_KEY_USER_ID, userId)
		c.Set(constant.CONTEXT_KEY_USER_ID, userId)
		c.Request = c.Request.WithContext(ctx)
		c.Next()
	}
}
