package controllers

import "net/http"

type WebApiErrorControllerInterface interface {
	PanicErrorHandler(writer http.ResponseWriter, request *http.Request, err interface{})
}
