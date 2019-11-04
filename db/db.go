package db

import (
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mssql"
)

type Config struct {
	Host         string
	Port         int
	Password     string
	User         string
	DBName       string
	MaxOpenConns int
	MaxIdleConns int
}

var CRMDB *gorm.DB

func Open(c *Config) (db *gorm.DB, err error) {
	connString := fmt.Sprintf("server=%s;user id=%s;password=%s;port=%d;database=%s;ApplicationIntent=ReadOnly", c.Host, c.User, c.Password, c.Port, c.DBName)

	dbConnection, err := gorm.Open("mssql", connString)
	if err != nil {
		return nil, fmt.Errorf("Error while connecting to db %v", err)
	}
	dbConnection.DB().SetMaxIdleConns(10)
	dbConnection.DB().SetMaxOpenConns(100)

	return dbConnection, nil
}
