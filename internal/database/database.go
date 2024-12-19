package database

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"ordent-test/internal/config"
)

var db *gorm.DB

func newDatabase() *gorm.DB {
	cfg := config.Get()

	user := cfg.PostgresUsername
	password := cfg.PostgresPassword
	host := cfg.PostgresHost
	port := cfg.PostgresPort
	db_name := cfg.PostgresDB

	url := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable", user, password, host, port, db_name)

	db, err := gorm.Open(
		postgres.Open(url),
		&gorm.Config{},
	)

	if err != nil {
		panic("failed to connect database" + err.Error())
	}

	return db
}

func Get() *gorm.DB {
	if db == nil {
		db = newDatabase()
	}

	return db
}
