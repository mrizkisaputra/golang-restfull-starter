package services

import (
	"context"
	"github.com/mrizkisaputra/golang-restfull-starter/cmd/model/dto"
)

type ProductServiceInterface interface {
	GetAllProduct(ctx context.Context) []dto.ProductResponseBody
	GetProductById(ctx context.Context, param dto.ProductRequestParam) dto.ProductResponseBody
	AddProduct(ctx context.Context, body dto.ProductRequestBody) dto.ProductResponseBody
	UpdateProduct(ctx context.Context, body dto.ProductRequestBody, param dto.ProductRequestParam) dto.ProductResponseBody
	RemoveProduct(ctx context.Context, param dto.ProductRequestParam)
}
