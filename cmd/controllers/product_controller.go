package controllers

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/julienschmidt/httprouter"
	"github.com/mrizkisaputra/golang-restfull-starter/cmd/exceptions"
	"github.com/mrizkisaputra/golang-restfull-starter/cmd/model/dto"
	. "github.com/mrizkisaputra/golang-restfull-starter/cmd/services"
	"github.com/mrizkisaputra/golang-restfull-starter/helper"
	"github.com/sirupsen/logrus"
	"net/http"
)

type productController struct {
	productService ProductServiceInterface
	context        context.Context
	Log            *logrus.Logger
}

func NewProductController(productService ProductServiceInterface, ctx context.Context, log *logrus.Logger) ProductControllerInterface {
	return &productController{
		productService: productService,
		context:        ctx,
		Log:            log,
	}
}

func (services *productController) writeToResponseBody(writer http.ResponseWriter, webApiResponse dto.WebApiResponseSuccess, httpStatusCode int) error {
	writer.WriteHeader(httpStatusCode)
	writer.Header().Set("Content-Type", "application/json charset=UTF-8")
	err := json.NewEncoder(writer).Encode(&webApiResponse)
	if err != nil {
		return fmt.Errorf("error while encoding response : %v", err)
	}
	return nil
}

func (services *productController) readFromRequestBody(request *http.Request, requestBody *dto.ProductRequestBody) error {
	decode := json.NewDecoder(request.Body)
	errDecoded := decode.Decode(requestBody)
	if errDecoded != nil {
		return fmt.Errorf("error while decoding request body : %v", errDecoded)
	}
	return nil
}

func (services *productController) GetProductsHandler(writer http.ResponseWriter, request *http.Request, param httprouter.Params) {
	products := services.productService.GetAllProduct(services.context)
	webApiResponseSuccess := dto.WebApiResponseSuccess{
		Status: "success",
		Data:   products,
	}
	if err := services.writeToResponseBody(writer, webApiResponseSuccess, http.StatusOK); err != nil {
		services.Log.Error(err)
		exceptions.ErrorInternal(err)
	}
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
	if err := services.writeToResponseBody(writer, webApiResponseSuccess, http.StatusOK); err != nil {
		services.Log.Error(err)
		exceptions.ErrorInternal(err)
	}
}

func (services *productController) AddProductHandler(writer http.ResponseWriter, request *http.Request, param httprouter.Params) {
	if content := request.Header.Get("Content-Type"); content != "application/json" {
		writer.WriteHeader(http.StatusUnsupportedMediaType)
		helper.PanicIfError(exceptions.NewBadRequestError("UNSUPPORTED CONTENT_TYPE"))
	}

	var productRequestBody = dto.ProductRequestBody{}
	if err := services.readFromRequestBody(request, &productRequestBody); err != nil {
		services.Log.Error(err.Error())
		exceptions.ErrorInternal(err)
	}
	product := services.productService.AddProduct(services.context, productRequestBody)

	webApiResponseSuccess := dto.WebApiResponseSuccess{
		Status: "success",
		Data:   product,
	}

	if err := services.writeToResponseBody(writer, webApiResponseSuccess, http.StatusCreated); err != nil {
		services.Log.Error(err)
		exceptions.ErrorInternal(err)
	}
}

func (services *productController) UpdateProductHandler(writer http.ResponseWriter, request *http.Request, param httprouter.Params) {
	id := param.ByName("id")
	var productRequestParam = dto.ProductRequestParam{Id: id}
	var productRequestBody = dto.ProductRequestBody{}

	if err := services.readFromRequestBody(request, &productRequestBody); err != nil {
		services.Log.Error(err)
		exceptions.ErrorInternal(err)
	}

	product := services.productService.UpdateProduct(services.context, productRequestBody, productRequestParam)
	webApiResponseSuccess := dto.WebApiResponseSuccess{
		Status: "success",
		Data:   product,
	}
	if err := services.writeToResponseBody(writer, webApiResponseSuccess, http.StatusAccepted); err != nil {
		services.Log.Error(err)
		exceptions.ErrorInternal(err)
	}
}

func (services *productController) RemoveProductHandler(writer http.ResponseWriter, request *http.Request, param httprouter.Params) {
	id := param.ByName("id")
	productRequestParam := dto.ProductRequestParam{Id: id}
	services.productService.RemoveProduct(services.context, productRequestParam)
	webApiResponseSuccess := dto.WebApiResponseSuccess{
		Status: "success",
		Data:   nil,
	}
	if err := services.writeToResponseBody(writer, webApiResponseSuccess, http.StatusOK); err != nil {
		services.Log.Error(err)
		exceptions.ErrorInternal(err)
	}
}
