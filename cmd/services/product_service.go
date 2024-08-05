package services

import (
	"context"
	"database/sql"
	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
	"github.com/mrizkisaputra/golang-restfull-starter/cmd/exceptions"
	"github.com/mrizkisaputra/golang-restfull-starter/cmd/model/dto"
	. "github.com/mrizkisaputra/golang-restfull-starter/cmd/model/entities"
	. "github.com/mrizkisaputra/golang-restfull-starter/cmd/repositories"
	"github.com/mrizkisaputra/golang-restfull-starter/utils"
)

type productService struct {
	productRepository ProductRepositoryInterface
	db                *sql.DB
	validate          *validator.Validate
}

func NewProductService(repository ProductRepositoryInterface, db *sql.DB) ProductServiceInterface {
	return &productService{
		productRepository: repository,
		db:                db,
		validate:          validator.New(validator.WithRequiredStructEnabled()),
	}
}

func toProductResponseBody(product Product) dto.ProductResponseBody {
	return dto.ProductResponseBody{
		Id:       product.Id,
		Item:     product.Item,
		Price:    product.Price,
		Quantity: product.Quantity,
	}
}

func (service *productService) GetAllProduct(ctx context.Context) []dto.ProductResponseBody {
	tx, err := service.db.Begin()
	exceptions.ErrorInternal(err)
	defer utils.RollbackIfPanic(tx)
	products := service.productRepository.FindAll(ctx, tx)

	var productResponseBody []dto.ProductResponseBody
	for _, product := range products {
		result := toProductResponseBody(product)
		productResponseBody = append(productResponseBody, result)
	}
	return productResponseBody
}

func (service *productService) GetProductById(ctx context.Context, param dto.ProductRequestParam) dto.ProductResponseBody {
	// validate request parameter
	validateParamError := service.validate.Struct(&param)
	utils.PanicIfError(validateParamError)

	tx, err := service.db.Begin()
	exceptions.ErrorInternal(err)
	defer utils.RollbackIfPanic(tx)

	product, err := service.productRepository.FindById(ctx, tx, param.Id)
	utils.PanicIfError(err)
	return toProductResponseBody(product)
}

func (service *productService) AddProduct(ctx context.Context, body dto.ProductRequestBody) dto.ProductResponseBody {
	// validate request payload body
	validateBodyError := service.validate.Struct(&body)
	utils.PanicIfError(validateBodyError)

	tx, err := service.db.Begin()
	exceptions.ErrorInternal(err)
	defer utils.RollbackIfPanic(tx)

	random, err := uuid.NewRandom()
	exceptions.ErrorInternal(err)
	var product Product = Product{
		Id:       random.String(),
		Item:     body.Item,
		Price:    body.Price,
		Quantity: body.Quantity,
	}
	product = service.productRepository.Create(ctx, tx, product)
	return toProductResponseBody(product)
}

func (service *productService) UpdateProduct(ctx context.Context, body dto.ProductRequestBody, param dto.ProductRequestParam) dto.ProductResponseBody {
	// validate request payload body
	validateBodyError := service.validate.Struct(&body)
	utils.PanicIfError(validateBodyError)

	// validate request parameter
	validateParamError := service.validate.Struct(&param)
	utils.PanicIfError(validateParamError)

	tx, err := service.db.Begin()
	exceptions.ErrorInternal(err)
	defer utils.RollbackIfPanic(tx)

	product, err := service.productRepository.FindById(ctx, tx, param.Id)
	utils.PanicIfError(err)
	product.Item = body.Item
	product.Price = body.Price
	product.Quantity = body.Quantity

	product = service.productRepository.Update(ctx, tx, product)
	return toProductResponseBody(product)
}

func (service *productService) RemoveProduct(ctx context.Context, param dto.ProductRequestParam) {
	// validate request parameter
	validateParamError := service.validate.Struct(&param)
	utils.PanicIfError(validateParamError)

	tx, err := service.db.Begin()
	exceptions.ErrorInternal(err)
	defer utils.RollbackIfPanic(tx)

	errDelete := service.productRepository.Delete(ctx, tx, param.Id)
	utils.PanicIfError(errDelete)
}
