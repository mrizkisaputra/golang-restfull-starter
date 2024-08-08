package auth

import (
	"encoding/json"
	"github.com/julienschmidt/httprouter"
	"github.com/mrizkisaputra/golang-restfull-starter/cmd/model/dto"
	"github.com/mrizkisaputra/golang-restfull-starter/helper"
	"net/http"
)

var example_username = "root"
var example_password = "root"

type AuthMiddleware struct {
	handler http.Handler
}

func NewAuthMiddleware(handler *httprouter.Router) http.Handler {
	return AuthMiddleware{handler: handler}
}

func (authMiddleware AuthMiddleware) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	username, password, ok := request.BasicAuth()
	if !ok {
		webApiResponseError := dto.WebApiResponseError{
			Status:           "error",
			Error:            "UNAUTHORIZED",
			DocumentationURL: "https://github.com/mrizkisaputra/golang-restfull-starter",
		}
		writer.WriteHeader(http.StatusBadRequest)
		err := json.NewEncoder(writer).Encode(webApiResponseError)
		helper.PanicIfError(err)
		return
	}

	if username != example_username || password != example_password {
		webApiResponseError := dto.WebApiResponseError{
			Status:           "error",
			Error:            "UNAUTHORIZED",
			DocumentationURL: "https://github.com/mrizkisaputra/golang-restfull-starter",
		}
		writer.WriteHeader(http.StatusUnauthorized)
		err := json.NewEncoder(writer).Encode(webApiResponseError)
		helper.PanicIfError(err)
		return
	}

	authMiddleware.handler.ServeHTTP(writer, request)
}
