package main

import (
	"context"
	"errors"
	"fmt"
	"github.com/sirupsen/logrus"
	"log"
	"net/http"
	"os"
	"os/signal"
	"pxj/CloudTravelShopApi/go/config"
	_ "pxj/CloudTravelShopApi/go/config"
	"pxj/CloudTravelShopApi/go/routes"
	"syscall"
	"time"
)

func main() {
	defer func() {
		if err := recover(); err != nil {
			println(err.(string))
		}
	}()
	router := routes.Routes
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
