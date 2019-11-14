package db

import (
	"database/sql"
	"fmt"
)
import _ "github.com/go-sql-driver/mysql"

type Config struct {
	Host         string
	Port         int
	Password     string
	User         string
	DBName       string
	MaxOpenConns int
	MaxIdleConns int
}

/*var DB *memdb.MemDB

func Open() (*memdb.MemDB, error) {

	// Create the DB schema
	schema := &memdb.DBSchema{
		Tables: map[string]*memdb.TableSchema{
			"user": &memdb.TableSchema{
				Name: "user",
				Indexes: map[string]*memdb.IndexSchema{
					"id": &memdb.IndexSchema{
						Name:    "id",
						Unique:  true,
						Indexer: &memdb.UintFieldIndex{Field: "Id"},
					},
					"age": &memdb.IndexSchema{
						Name:    "age",
						Unique:  false,
						Indexer: &memdb.IntFieldIndex{Field: "Age"},
					},
					"gender": &memdb.IndexSchema{
						Name:    "gender",
						Unique:  false,
						Indexer: &memdb.StringFieldIndex{Field: "Gender"},
					},
				},
			},
			"location": &memdb.TableSchema{
				Name: "location",
				Indexes: map[string]*memdb.IndexSchema{
					"id": &memdb.IndexSchema{
						Name:    "id",
						Unique:  true,
						Indexer: &memdb.UintFieldIndex{Field: "Id"},
					},
					"city": &memdb.IndexSchema{
						Name:    "city",
						Unique:  false,
						Indexer: &memdb.StringFieldIndex{Field: "City"},
					},
					"place": &memdb.IndexSchema{
						Name:    "place",
						Unique:  false,
						Indexer: &memdb.StringFieldIndex{Field: "Place"},
					},
				},
			},
			"visit": &memdb.TableSchema{
				Name: "visit",
				Indexes: map[string]*memdb.IndexSchema{
					"id": &memdb.IndexSchema{
						Name:    "id",
						Unique:  true,
						Indexer: &memdb.UintFieldIndex{Field: "Id"},
					},
					"user": &memdb.IndexSchema{
						Name:    "user",
						Unique:  false,
						Indexer: &memdb.IntFieldIndex{Field: "User"},
					},
					"location": &memdb.IndexSchema{
						Name:    "location",
						Unique:  false,
						Indexer: &memdb.IntFieldIndex{Field: "Location"},
					},
				},
			},
		},
	}

	dbConnection, err := memdb.NewMemDB(schema)
	return dbConnection, err

}*/

var DB *sql.DB

func Open(c *Config) (dbConnection *sql.DB, err error) {
	dataSourceName := fmt.Sprint(c.Host, ":", c.Password, "@tcp(", c.Host, ":", c.Port, ")/", c.DBName) //"username:password@tcp(127.0.0.1:3306)/test"
	dbConnection, err = sql.Open("mysql", dataSourceName)
	if err != nil {
		return nil, fmt.Errorf("Error while connecting to db %v", err)
	}

	return dbConnection, nil
}

//var CRMDB *gorm.DB

//func Open(c *Config) (db *gorm.DB, err error) {
//	connString := fmt.Sprintf("server=%s;user id=%s;password=%s;port=%d;database=%s", c.Host, c.User, c.Password, c.Port, c.DBName)
//
//	dbConnection, err := gorm.Open("mssql", connString)
//	if err != nil {
//		return nil, fmt.Errorf("Error while connecting to db %v", err)
//	}
//	dbConnection.DB().SetMaxIdleConns(10)
//	dbConnection.DB().SetMaxOpenConns(100)
//
//	return dbConnection, nil
//}
