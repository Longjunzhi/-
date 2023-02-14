package main

import (
	"context"
	"errors"
	"fmt"
	"github.com/sirupsen/logrus"
	_ "github.com/swaggo/files"
	swaggerFiles "github.com/swaggo/files"
	_ "github.com/swaggo/gin-swagger"
	ginSwagger "github.com/swaggo/gin-swagger"
	"log"
	"net/http"
	"os"
	"os/signal"
	"pxj/CloudTravelShopApi/go/config"
	_ "pxj/CloudTravelShopApi/go/config"
	"pxj/CloudTravelShopApi/go/database"
	docs "pxj/CloudTravelShopApi/go/docs"
	"pxj/CloudTravelShopApi/go/models"
	"pxj/CloudTravelShopApi/go/routes"
	"syscall"
	"time"
)

func main() {
	// 初始化
	models.DbInit()
	database.DbInit()
	defer func() {
		if err := recover(); err != nil {
			println(err.(string))
		}
	}()
	router := routes.Routes
	docs.SwaggerInfo.BasePath = ""
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	address := fmt.Sprintf("%s:%d", config.AppConf.ServerConf.Host, config.AppConf.ServerConf.Port)
	server := &http.Server{
		Addr:           address,
		Handler:        router,
		ReadTimeout:    61 * time.Second,
		WriteTimeout:   61 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	logrus.Infof("runing server at : %v", address)
	go func() {
		err := server.ListenAndServe()
		if err != nil && errors.Is(err, http.ErrServerClosed) {
			logrus.Errorf("Server Listen Error: %s\n ", err)
		}
	}()
	quit := make(chan os.Signal)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Println("Shutting down server ...")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := server.Shutdown(ctx); err != nil {
		log.Fatal("Server forced to shutdown:", err)
	}
	log.Println("Server exiting")
}
