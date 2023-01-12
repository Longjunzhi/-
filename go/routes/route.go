package routes

import (
	"github.com/gin-gonic/gin"
	"pxj/CloudTravelShopApi/go/controllers"
	"pxj/CloudTravelShopApi/go/middlewares"
)

var Routes *gin.Engine

func init() {
	Routes = gin.Default()

	// api前端接口
	authApi := Routes.Group("/api")
	authApi.Use(middlewares.ApiAuth())
	api := Routes.Group("/api")

	// 不需要token接口
	api.POST("/login", controllers.ApiLogin)
	// 需要token接口
	authApi.POST("/users/getUserById", controllers.ApiLogin)

	// admin后台接口
	authAdmin := Routes.Group("/admin")
	authAdmin.Use(middlewares.AdminAuth())
	admin := Routes.Group("/admin")

	// 不需要token接口
	admin.POST("/login/account", controllers.AdminUserLoginByAccount)
	// 需要token接口
	authAdmin.POST("/currentUser", controllers.CurrentUser)
	authAdmin.POST("/login/outLogin", controllers.OutLogin)
	authAdmin.POST("/rule", controllers.Rule)

}
