package router

import (
	"github.com/julienschmidt/httprouter"
	. "github.com/mrizkisaputra/golang-restfull-starter/cmd/controllers"
	. "github.com/mrizkisaputra/golang-restfull-starter/cmd/repositories"
	. "github.com/mrizkisaputra/golang-restfull-starter/cmd/services"
)

type productHttpRouter struct {
	productRepository ProductRepositoryInterface
	productService    ProductServiceInterface
	productController ProductControllerInterface
	router            *httprouter.Router
}

func NewProductHttpRouter(
	repository ProductRepositoryInterface,
	service ProductServiceInterface,
	controller ProductControllerInterface) productHttpRouter {

	return productHttpRouter{
		productRepository: repository,
		productService:    service,
		productController: controller,
		router:            httprouter.New(),
	}
}

func (p *productHttpRouter) GetRoute() *httprouter.Router {
	p.router.GET("/api/products", p.productController.GetProductsHandler)
	p.router.GET("/api/products/:id", p.productController.GetProductByIdHandler)
	p.router.POST("/api/product", p.productController.AddProductHandler)
	p.router.PUT("/api/products/:id", p.productController.UpdateProductHandler)
	p.router.DELETE("/api/products/:id", p.productController.RemoveProductHandler)

	return p.router
}
