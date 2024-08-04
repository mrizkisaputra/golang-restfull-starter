package controllers

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
)

type ProductControllerInterface interface {
	GetProductsHandler(writer http.ResponseWriter, request *http.Request, param httprouter.Params)
	GetProductByIdHandler(writer http.ResponseWriter, request *http.Request, param httprouter.Params)
	AddProductHandler(writer http.ResponseWriter, request *http.Request, param httprouter.Params)
	UpdateProductHandler(writer http.ResponseWriter, request *http.Request, param httprouter.Params)
	RemoveProductHandler(writer http.ResponseWriter, request *http.Request, param httprouter.Params)
}
