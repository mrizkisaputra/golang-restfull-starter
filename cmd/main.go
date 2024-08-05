package main

import (
	"context"
	"github.com/mrizkisaputra/golang-restfull-starter/cmd/controllers"
	"github.com/mrizkisaputra/golang-restfull-starter/cmd/repositories"
	"github.com/mrizkisaputra/golang-restfull-starter/cmd/router"
	"github.com/mrizkisaputra/golang-restfull-starter/cmd/services"
	"github.com/mrizkisaputra/golang-restfull-starter/config"
	"github.com/mrizkisaputra/golang-restfull-starter/utils"
	"net/http"
)

var db = config.GetConnectDB()

func main() {
	webApiErrorController := controllers.NewWebApiErrorController()
	productRepository := repositories.NewProductRepository()
	productService := services.NewProductService(productRepository, db)
	productController := controllers.NewProductController(productService, context.Background())
	productHttpRouter := router.NewProductHttpRouter(productRepository, productService, productController, webApiErrorController)
	route := productHttpRouter.GetRoute()

	server := http.Server{
		Addr:    "localhost:3000",
		Handler: route,
	}
	err := server.ListenAndServe()
	utils.PanicIfError(err)
}
