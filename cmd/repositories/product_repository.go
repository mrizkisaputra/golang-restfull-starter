package repositories

import (
	"context"
	"database/sql"
	"github.com/mrizkisaputra/golang-restfull-starter/cmd/exceptions"
	. "github.com/mrizkisaputra/golang-restfull-starter/cmd/model/entities"
	"github.com/mrizkisaputra/golang-restfull-starter/helper"
	"github.com/sirupsen/logrus"
)

type productRepository struct {
	Log *logrus.Logger
}

func NewProductRepository(log *logrus.Logger) ProductRepositoryInterface {
	productRepository := new(productRepository)
	productRepository.Log = log
	return productRepository
}

func (p *productRepository) FindAll(ctx context.Context, tx *sql.Tx) []Product {
	rows, err := tx.QueryContext(ctx, helper.ProductFindAllSQL)
	if err != nil {
		p.Log.WithFields(logrus.Fields{
			"layer": "repositories/product_repository/FindAll()",
		}).Errorf("error because : %v", err.Error())
		exceptions.ErrorInternal(err)
	}
	defer helper.CloseQuery(rows, p.Log)
	var products []Product
	for rows.Next() {
		var product = Product{}
		err := rows.Scan(&product.Id, &product.Item, &product.Price, &product.Quantity)
		if err != nil {
			p.Log.WithFields(logrus.Fields{
				"layer": "repositories/product_repository/FindAll()",
			}).Errorf("error because : %v", err.Error())
			exceptions.ErrorInternal(err)
		}
		products = append(products, product)
	}
	return products
}

func (p *productRepository) FindById(ctx context.Context, tx *sql.Tx, productId string) (Product, error) {
	rows, err := tx.QueryContext(ctx, helper.ProductFindByIdSQL, productId)
	if err != nil {
		p.Log.WithFields(logrus.Fields{
			"layer": "repositories/product_repository/FindById()",
		}).Errorf("error because : %v", err.Error())
		exceptions.ErrorInternal(err)
	}
	defer helper.CloseQuery(rows, p.Log)
	var product = Product{}
	if rows.Next() {
		err := rows.Scan(&product.Id, &product.Item, &product.Price, &product.Quantity)
		if err != nil {
			p.Log.WithFields(logrus.Fields{
				"layer": "repositories/product_repository/FindById()",
			}).Errorf("error because : %v", err.Error())
			exceptions.ErrorInternal(err)
		}
	} else {
		p.Log.Errorf("product with id (%s) not found", productId)
		return Product{}, exceptions.NewNotFoundError("NOT FOUND")
	}
	return product, nil
}

func (p *productRepository) Create(ctx context.Context, tx *sql.Tx, product Product) Product {
	_, err := tx.ExecContext(ctx, helper.ProductCreateSQL, product.Id, product.Item, product.Price, product.Quantity)
	if err != nil {
		p.Log.WithFields(logrus.Fields{
			"layer": "repositories/product_repository/Create()",
		}).Errorf("error because : %v", err.Error())
		exceptions.ErrorInternal(err)
	}
	return product
}

func (p *productRepository) Update(ctx context.Context, tx *sql.Tx, product Product) Product {
	_, err := tx.ExecContext(ctx, helper.ProductUpdateSQL, product.Item, product.Price, product.Quantity, product.Id)
	if err != nil {
		p.Log.WithFields(logrus.Fields{
			"layer": "repositories/product_repository/Update()",
		}).Errorf("error because : %v", err.Error())
		exceptions.ErrorInternal(err)
	}
	return product
}

func (p *productRepository) Delete(ctx context.Context, tx *sql.Tx, productId string) error {
	result, err := tx.ExecContext(ctx, helper.ProductDeleteSQL, productId)
	if err != nil {
		p.Log.WithFields(logrus.Fields{
			"layer": "repositories/product_repository/Delete()",
		}).Errorf("error because : %v", err.Error())
		exceptions.ErrorInternal(err)
	}
	row, err2 := result.RowsAffected()
	if err2 != nil {
		p.Log.WithFields(logrus.Fields{
			"layer": "repositories/product_repository/Delete()",
		}).Errorf("error because : %v", err.Error())
		exceptions.ErrorInternal(err)
	}
	if row >= 1 {
		return nil
	}
	p.Log.WithFields(logrus.Fields{
		"layer": "repositories/product_repository/Delete()",
	}).Errorf("error because id (%s) not found", productId)
	return exceptions.NewNotFoundError("NOT FOUND")
}
