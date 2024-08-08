package helper

import (
	"database/sql"
	"github.com/sirupsen/logrus"
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
		panic(err)
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

func CloseDB(db *sql.DB, Log *logrus.Logger) {
	if err := db.Close(); err != nil {
		Log.Errorf("close database error : %v", err)
	}
}

func CloseQuery(rows *sql.Rows, log *logrus.Logger) {
	if err := rows.Close(); err != nil {
		log.Errorf("Error closing rows: %s", err)
	}
}
