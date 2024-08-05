package router

import (
	"github.com/julienschmidt/httprouter"
	. "github.com/mrizkisaputra/golang-restfull-starter/cmd/controllers"
	. "github.com/mrizkisaputra/golang-restfull-starter/cmd/repositories"
	. "github.com/mrizkisaputra/golang-restfull-starter/cmd/services"
)

type productHttpRouter struct {
	productRepository     ProductRepositoryInterface
	productService        ProductServiceInterface
	productController     ProductControllerInterface
	webApiErrorController WebApiErrorControllerInterface
	router                *httprouter.Router
}

func NewProductHttpRouter(
	repository ProductRepositoryInterface,
	service ProductServiceInterface,
	controller ProductControllerInterface,
	webApiErrorController WebApiErrorControllerInterface) productHttpRouter {

	return productHttpRouter{
		productRepository:     repository,
		productService:        service,
		productController:     controller,
		webApiErrorController: webApiErrorController,
		router:                httprouter.New(),
	}
}

func (p *productHttpRouter) GetRoute() *httprouter.Router {
	p.router.GET("/api/products", p.productController.GetProductsHandler)
	p.router.GET("/api/products/:id", p.productController.GetProductByIdHandler)
	p.router.POST("/api/product", p.productController.AddProductHandler)
	p.router.PUT("/api/products/:id", p.productController.UpdateProductHandler)
	p.router.DELETE("/api/products/:id", p.productController.RemoveProductHandler)

	p.router.PanicHandler = p.webApiErrorController.PanicErrorHandler

	return p.router
}
