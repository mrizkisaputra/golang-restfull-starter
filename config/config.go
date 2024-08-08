package config

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/spf13/viper"
)

var v = viper.New()

func init() {
	v.SetConfigName("config")
	v.SetConfigType("env")
	v.AddConfigPath("./config")
}

func GetConnectDB() (*sql.DB, error) {
	err := v.ReadInConfig()
	if err != nil {
		return nil, err
	}
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s",
		v.GetString("DB_USERNAME"),
		v.GetString("DB_PASSWORD"),
		v.GetString("DB_HOST"),
		v.GetString("DB_PORT"),
		v.GetString("DB_NAME"))

	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, err
	}
	return db, nil
}
