package controllers

import (
	"context"
	"encoding/json"
	"github.com/julienschmidt/httprouter"
	"github.com/mrizkisaputra/golang-restfull-starter/cmd/exceptions"
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

func (services *productController) writeToResponseBody(writer http.ResponseWriter, webApiResponse dto.WebApiResponseSuccess, httpStatusCode int) {
	writer.WriteHeader(httpStatusCode)
	writer.Header().Set("Content-Type", "application/json charset=UTF-8")
	err := json.NewEncoder(writer).Encode(&webApiResponse)
	exceptions.ErrorInternal(err)
}

func (services *productController) readFromRequestBody(request *http.Request, requestBody *dto.ProductRequestBody) {
	decode := json.NewDecoder(request.Body)
	errDecoded := decode.Decode(requestBody)
	exceptions.ErrorInternal(errDecoded)
}

func (services *productController) GetProductsHandler(writer http.ResponseWriter, request *http.Request, param httprouter.Params) {
	products := services.productService.GetAllProduct(services.context)
	webApiResponseSuccess := dto.WebApiResponseSuccess{
		Status: "success",
		Data:   products,
	}
	services.writeToResponseBody(writer, webApiResponseSuccess, http.StatusOK)
}

func (services *productController) GetProductByIdHandler(writer http.ResponseWriter, request *http.Request, param httprouter.Params) {
	productId := param.ByName("id")
	requestParam := dto.ProductRequestParam{
		Id: productId,
	}
	product := services.productService.GetProductById(services.context, requestParam)
	webApiResponseSuccess := dto.WebApiResponseSuccess{
		Status: "success",
		Data:   product,
	}
	services.writeToResponseBody(writer, webApiResponseSuccess, http.StatusOK)
}

func (services *productController) AddProductHandler(writer http.ResponseWriter, request *http.Request, param httprouter.Params) {
	if content := request.Header.Get("Content-Type"); content != "application/json" {
		writer.WriteHeader(http.StatusUnsupportedMediaType)
		utils.PanicIfError(exceptions.NewBadRequestError("UNSUPPORTED CONTENT_TYPE"))
	}

	var productRequestBody = dto.ProductRequestBody{}
	services.readFromRequestBody(request, &productRequestBody)
	product := services.productService.AddProduct(services.context, productRequestBody)

	webApiResponseSuccess := dto.WebApiResponseSuccess{
		Status: "success",
		Data:   product,
	}
	services.writeToResponseBody(writer, webApiResponseSuccess, http.StatusCreated)
}

func (services *productController) UpdateProductHandler(writer http.ResponseWriter, request *http.Request, param httprouter.Params) {
	id := param.ByName("id")
	var productRequestParam = dto.ProductRequestParam{Id: id}
	var productRequestBody = dto.ProductRequestBody{}

	services.readFromRequestBody(request, &productRequestBody)

	product := services.productService.UpdateProduct(services.context, productRequestBody, productRequestParam)
	webApiResponseSuccess := dto.WebApiResponseSuccess{
		Status: "success",
		Data:   product,
	}
	services.writeToResponseBody(writer, webApiResponseSuccess, http.StatusAccepted)
}

func (services *productController) RemoveProductHandler(writer http.ResponseWriter, request *http.Request, param httprouter.Params) {
	id := param.ByName("id")
	productRequestParam := dto.ProductRequestParam{Id: id}
	services.productService.RemoveProduct(services.context, productRequestParam)
	webApiResponseSuccess := dto.WebApiResponseSuccess{
		Status: "success",
		Data:   nil,
	}
	services.writeToResponseBody(writer, webApiResponseSuccess, http.StatusOK)
}
