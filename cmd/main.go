package main

import (
	"context"
	"fmt"
	"github.com/mrizkisaputra/golang-restfull-starter/cmd/auth"
	"github.com/mrizkisaputra/golang-restfull-starter/cmd/controllers"
	"github.com/mrizkisaputra/golang-restfull-starter/cmd/exceptions"
	"github.com/mrizkisaputra/golang-restfull-starter/cmd/repositories"
	"github.com/mrizkisaputra/golang-restfull-starter/cmd/router"
	"github.com/mrizkisaputra/golang-restfull-starter/cmd/services"
	"github.com/mrizkisaputra/golang-restfull-starter/config"
	"github.com/mrizkisaputra/golang-restfull-starter/helper"
	"github.com/sirupsen/logrus"
	"io"
	"net/http"
	"os"
	"os/signal"
	"path/filepath"
	"syscall"
	"time"
)

var Log = logrus.New()

func main() {
	if err := setupLogging(); err != nil {
		Log.Error(err)
	}

	if err := run(); err != nil {
		Log.Error(err)
	}

}

func setupLogging() error {
	fileLocation, err := filepath.Abs("./log")
	if err != nil {
		return fmt.Errorf("filepath abs error : %v", err)
	}
	file, err := os.OpenFile(filepath.Join(fileLocation, "/logs.log"), os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		return fmt.Errorf("open file error : %v", err)
	}

	Log.SetLevel(logrus.InfoLevel)
	Log.SetFormatter(&logrus.JSONFormatter{
		TimestampFormat: time.RFC3339,
	})
	Log.SetOutput(io.MultiWriter(file, os.Stdout))
	return nil
}

func run() error {
	Log.WithField("filename", "main.go").Info("Started")
	defer Log.WithField("filename", "main.go").Info("Completed")

	// open database connection
	db, errDB := config.GetConnectDB()
	if errDB != nil {
		exceptions.ErrorInternal(errDB)
		return fmt.Errorf("open connection database error : %v", errDB)
	}
	defer helper.CloseDB(db, Log)

	productRepository := repositories.NewProductRepository(Log)
	productService := services.NewProductService(productRepository, db, Log)
	productController := controllers.NewProductController(productService, context.Background(), Log)
	webApiErrorController := controllers.NewWebApiErrorController()
	productHttpRouter := router.NewProductHttpRouter(productRepository, productService, productController, webApiErrorController)
	route := productHttpRouter.GetRoute()
	authMiddleware := auth.NewAuthMiddleware(route)

	// parameter server
	server := http.Server{
		Addr:         "localhost:3000",
		Handler:      authMiddleware,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 5 * time.Second,
	}

	// start listening server
	serveErrorChannel := make(chan error, 1)
	defer close(serveErrorChannel)
	shutdownChannel := make(chan os.Signal, 1)
	signal.Notify(shutdownChannel, os.Interrupt, syscall.SIGTERM)
	go listenAndServe(serveErrorChannel, &server)

	select {
	case err := <-serveErrorChannel:
		{
			return fmt.Errorf("start server error : %v", err)
		}
	case <-shutdownChannel:
		{
			ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
			defer cancel()
			if err := server.Shutdown(ctx); err != nil {
				Log.Error("gracefully shutting down server : %s", err.Error())
				if err := server.Close(); err != nil {
					return fmt.Errorf("closing server error : %v", err)
				}
			}
		}
	}

	return nil
}

func listenAndServe(channel chan<- error, server *http.Server) {
	err := server.ListenAndServe()
	if err != nil {
		channel <- err
	}
	defer close(channel)
}
