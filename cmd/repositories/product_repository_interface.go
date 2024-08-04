package repositories

import (
	"context"
	"database/sql"
	. "github.com/mrizkisaputra/golang-restfull-starter/cmd/model/entities"
)

type ProductRepositoryInterface interface {
	FindAll(ctx context.Context, tx *sql.Tx) []Product
	FindById(ctx context.Context, tx *sql.Tx, productId string) (Product, error)
	Create(ctx context.Context, tx *sql.Tx, product Product) Product
	Update(ctx context.Context, tx *sql.Tx, product Product) Product
	Delete(ctx context.Context, tx *sql.Tx, productId string) error
}
