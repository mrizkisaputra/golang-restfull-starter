package repositories

import (
	"context"
	"database/sql"
	"errors"
	. "github.com/mrizkisaputra/golang-restfull-starter/cmd/model/entities"
	"github.com/mrizkisaputra/golang-restfull-starter/utils"
)

type productRepository struct {
}

func NewProductRepository() ProductRepositoryInterface {
	return new(productRepository)
}

func (p *productRepository) FindAll(ctx context.Context, tx *sql.Tx) []Product {
	rows, err := tx.QueryContext(ctx, utils.ProductFindAllSQL)
	utils.PanicIfError(err)
	defer rows.Close()
	var products []Product
	for rows.Next() {
		var product = Product{}
		err := rows.Scan(&product.Id, &product.Item, &product.Price, &product.Quantity)
		utils.PanicIfError(err)
		products = append(products, product)
	}
	return products
}

func (p *productRepository) FindById(ctx context.Context, tx *sql.Tx, productId string) (Product, error) {
	rows, err := tx.QueryContext(ctx, utils.ProductFindByIdSQL, productId)
	utils.PanicIfError(err)
	defer rows.Close()
	var product = Product{}
	if rows.Next() {
		err := rows.Scan(&product.Id, &product.Item, &product.Price, &product.Quantity)
		utils.PanicIfError(err)
	} else {
		return Product{}, errors.New("product with id " + productId + " not found")
	}
	return product, nil
}

func (p *productRepository) Create(ctx context.Context, tx *sql.Tx, product Product) Product {
	_, err := tx.ExecContext(ctx, utils.ProductCreateSQL, product.Id, product.Item, product.Price, product.Quantity)
	utils.PanicIfError(err)
	return product
}

func (p *productRepository) Update(ctx context.Context, tx *sql.Tx, product Product) Product {
	_, err := tx.ExecContext(ctx, utils.ProductUpdateSQL, product.Item, product.Price, product.Quantity, product.Id)
	utils.PanicIfError(err)
	return product
}

func (p *productRepository) Delete(ctx context.Context, tx *sql.Tx, productId string) error {
	result, err := tx.ExecContext(ctx, utils.ProductDeleteSQL, productId)
	utils.PanicIfError(err)
	row, err2 := result.RowsAffected()
	utils.PanicIfError(err2)
	if row >= 1 {
		return nil
	}
	return errors.New("Delete data with id " + productId + " fail")
}
