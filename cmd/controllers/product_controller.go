package controllers

import (
	"context"
	"encoding/json"
	"github.com/julienschmidt/httprouter"
	"github.com/mrizkisaputra/golang-restfull-starter/cmd/model/dto"
	. "github.com/mrizkisaputra/golang-restfull-starter/cmd/services"
	"github.com/mrizkisaputra/golang-restfull-starter/utils"
	"net/http"
)

type productController struct {
	productService ProductServiceInterface
	context        context.Context
}

func NewProductController(productService ProductServiceInterface, ctx context.Context) ProductControllerInterface {
	return &productController{
		productService: productService,
		context:        ctx,
	}
}

func buildResponseEntity(products []dto.ProductResponseBody) dto.WebApiResponseSuccess {
	return dto.WebApiResponseSuccess{
		Status: "Success",
		Data:   products,
	}
}

func (services *productController) GetProductsHandler(writer http.ResponseWriter, request *http.Request, param httprouter.Params) {
	products := services.productService.GetAllProduct(services.context)
	webApiResponseSuccess := buildResponseEntity(products)
	encode := json.NewEncoder(writer)
	err := encode.Encode(webApiResponseSuccess)
	utils.PanicIfError(err)
	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(http.StatusOK)
}

func (services *productController) GetProductByIdHandler(writer http.ResponseWriter, request *http.Request, param httprouter.Params) {
	productId := param.ByName("id")
	requestParam := dto.ProductRequestParam{
		Id: productId,
	}
	product := services.productService.GetProductById(services.context, requestParam)
	var response []dto.ProductResponseBody
	response = append(response, product)
	webApiResponseSuccess := buildResponseEntity(response)
	encode := json.NewEncoder(writer)
	err := encode.Encode(webApiResponseSuccess)
	utils.PanicIfError(err)
	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(http.StatusOK)
}

func (services *productController) AddProductHandler(writer http.ResponseWriter, request *http.Request, param httprouter.Params) {
	if content := request.Header.Get("Content-Type"); content != "application/json" {
		writer.WriteHeader(http.StatusUnsupportedMediaType)
	}

	// decode reqest json
	var requestBody = dto.ProductRequestBody{}
	decode := json.NewDecoder(request.Body)
	errDecoded := decode.Decode(&requestBody)
	utils.PanicIfError(errDecoded)

	product := services.productService.AddProduct(services.context, requestBody)
	var response []dto.ProductResponseBody
	response = append(response, product)
	webApiResponseSuccess := buildResponseEntity(response)
	errEncoded := json.NewEncoder(writer).Encode(webApiResponseSuccess)
	utils.PanicIfError(errEncoded)
	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(http.StatusCreated)
}

func (services *productController) UpdateProductHandler(writer http.ResponseWriter, request *http.Request, param httprouter.Params) {
	id := param.ByName("id")
	var productRequestParam = dto.ProductRequestParam{Id: id}
	var productRequestBody = dto.ProductRequestBody{}
	errDecoded := json.NewDecoder(request.Body).Decode(&productRequestBody)
	utils.PanicIfError(errDecoded)

	product := services.productService.UpdateProduct(services.context, productRequestBody, productRequestParam)
	var response []dto.ProductResponseBody
	response = append(response, product)

	webApiResponseSuccess := buildResponseEntity(response)
	encode := json.NewEncoder(writer)
	errEncoded := encode.Encode(webApiResponseSuccess)
	utils.PanicIfError(errEncoded)

	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(http.StatusAccepted)
}

func (services *productController) RemoveProductHandler(writer http.ResponseWriter, request *http.Request, param httprouter.Params) {
	id := param.ByName("id")
	productRequestParam := dto.ProductRequestParam{Id: id}
	services.productService.RemoveProduct(services.context, productRequestParam)
	err := json.NewEncoder(writer).Encode(buildResponseEntity(nil))
	utils.PanicIfError(err)
	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(http.StatusOK)
}
