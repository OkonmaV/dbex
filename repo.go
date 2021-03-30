package dbex

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type MySqlConnection struct {
	DB *gorm.DB
}

func NewMySqlConnection(connectionString string) (*MySqlConnection, error) {
	db, err := gorm.Open(mysql.Open(connectionString), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Silent),
	})
	if err != nil {
		return nil, err
	}
	return &MySqlConnection{DB: db}, nil
}

func (conn *MySqlConnection) Close() error {
	db, err := conn.DB.DB()
	if err != nil {
		return err
	}
	return db.Close()
}
