package utils

import (
	"database/sql"
)

var (
	ProductFindAllSQL  = "SELECT id, item, price, quantity FROM products"
	ProductFindByIdSQL = "SELECT id , item, price, quantity FROM products where id = ?"
	ProductCreateSQL   = "INSERT INTO products(id, item, price, quantity) VALUES (?,?,?,?)"
	ProductUpdateSQL   = "UPDATE products SET item = ?, price = ?, quantity = ? WHERE id = ?"
	ProductDeleteSQL   = "DELETE FROM products WHERE id = ?"
)

func PanicIfError(err error) {
	if err != nil {
		panic(err.Error())
	}
}

func RollbackIfPanic(tx *sql.Tx) {
	if errPanic := recover(); errPanic != nil {
		err := tx.Rollback()
		PanicIfError(err)
		panic(errPanic)
	} else {
		err := tx.Commit()
		PanicIfError(err)
	}
}
