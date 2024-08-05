package controllers

import (
	"encoding/json"
	"github.com/go-playground/validator/v10"
	"github.com/mrizkisaputra/golang-restfull-starter/cmd/exceptions"
	"github.com/mrizkisaputra/golang-restfull-starter/cmd/model/dto"
	"net/http"
)

type WebApiErrorController struct {
}

func NewWebApiErrorController() WebApiErrorControllerInterface {
	return &WebApiErrorController{}
}

func (webApiErrorController *WebApiErrorController) writeToResponseBody(writer http.ResponseWriter, webApiResponseError *dto.WebApiResponseError, statusCode int) {
	writer.WriteHeader(statusCode)
	writer.Header().Set("Content-Type", "application/json")
	err := json.NewEncoder(writer).Encode(webApiResponseError)
	exceptions.ErrorInternal(err)
}

func (webApiErrorController *WebApiErrorController) PanicErrorHandler(writer http.ResponseWriter, request *http.Request, err interface{}) {
	switch err.(type) {
	case exceptions.NotFoundError:
		{
			notFound := err.(exceptions.NotFoundError)
			webApiErrorController.notFoundError(writer, notFound)
		}
	case exceptions.InternalServerError:
		{
			errorInternal := err.(exceptions.InternalServerError)
			webApiErrorController.internalServerError(writer, errorInternal)
		}
	case exceptions.BadRequestError:
		{
			errorBadRequest := err.(exceptions.BadRequestError)
			webApiErrorController.badRequestError(writer, errorBadRequest)
		}
	case validator.ValidationErrors:
		{
			// error ini terjadi ketika ada masalah karena validasi data yang gagal
			errorValidated := err.(validator.ValidationErrors)
			webApiErrorController.validationError(writer, errorValidated)
		}
	case validator.InvalidValidationError:
		{
			// error ini terjadi ketika ada masalah dengan cara validasi dilakukan, bukan dengan data yang divalidasi itu sendiri
			errorValidation := err.(*validator.InvalidValidationError)
			webApiErrorController.invalidValidationError(writer, errorValidation)
		}
	}
}

func (webApiErrorController *WebApiErrorController) notFoundError(writer http.ResponseWriter, err exceptions.NotFoundError) {
	webApiResponseError := dto.WebApiResponseError{
		Status:           "error",
		Error:            err.Error(),
		TraceId:          "",
		DocumentationURL: "https://github.com/mrizkisaputra/golang-restfull-starter",
	}
	webApiErrorController.writeToResponseBody(writer, &webApiResponseError, http.StatusNotFound)
}

func (webApiErrorController *WebApiErrorController) internalServerError(writer http.ResponseWriter, err exceptions.InternalServerError) {
	webApiResponseError := dto.WebApiResponseError{
		Status:           "error",
		Error:            err.Error(),
		TraceId:          "",
		DocumentationURL: "https://github.com/mrizkisaputra/golang-restfull-starter",
	}
	webApiErrorController.writeToResponseBody(writer, &webApiResponseError, http.StatusInternalServerError)
}

func (webApiErrorController *WebApiErrorController) badRequestError(writer http.ResponseWriter, err exceptions.BadRequestError) {
	webApiResponseError := dto.WebApiResponseError{
		Status:           "error",
		Error:            err.Error(),
		TraceId:          "",
		DocumentationURL: "https://github.com/mrizkisaputra/golang-restfull-starter",
	}
	webApiErrorController.writeToResponseBody(writer, &webApiResponseError, http.StatusBadRequest)
}

func (webApiErrorController *WebApiErrorController) validationError(writer http.ResponseWriter, err validator.ValidationErrors) {
	var validatedErrors = make(map[string][]any)
	for _, e := range err {
		validatedErrors[e.Field()] = append(validatedErrors[e.Field()], e.Error())
	}
	webApiResponseError := dto.WebApiResponseError{
		Status:           "error",
		Error:            validatedErrors,
		TraceId:          "",
		DocumentationURL: "https://github.com/mrizkisaputra/golang-restfull-starter",
	}
	webApiErrorController.writeToResponseBody(writer, &webApiResponseError, http.StatusBadRequest)
}

func (webApiErrorController *WebApiErrorController) invalidValidationError(writer http.ResponseWriter, err *validator.InvalidValidationError) {
	webApiResponseError := dto.WebApiResponseError{
		Status:           "error",
		Error:            err.Error(),
		TraceId:          "",
		DocumentationURL: "https://github.com/mrizkisaputra/golang-restfull-starter",
	}
	webApiErrorController.writeToResponseBody(writer, &webApiResponseError, http.StatusInternalServerError)
}
