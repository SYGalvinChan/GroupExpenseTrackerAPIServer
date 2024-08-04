package storage

import (
	"GroupExpenseTracker/pkg/config"
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Storage struct {
	db *gorm.DB
}

func New() Storage {
	databaseConfig := config.GetConfig().DatabaseConfig
	user := databaseConfig.User
	password := databaseConfig.Password
	address := databaseConfig.Address
	databaseName := databaseConfig.DatabaseName

	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", user, password, address, databaseName)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("unable to connect to db")
	}

	return Storage{
		db: db,
	}
}
