package dbex

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type MySqlConnection struct {
	DB *gorm.DB
}

func NewMySqlConnection(connectionString string) (*MySqlConnection, error) {
	db, err := gorm.Open(mysql.Open(connectionString), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	return &MySqlConnection{DB: db}, nil
}
