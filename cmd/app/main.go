package main

import (
	"fmt"
	"ordent-test/config"
	database2 "ordent-test/internal/infrastructure/db"
	"ordent-test/internal/router"
)

// @title Ordent Test API
// @version 1.0
// @description This is a sample server for Ordent Test.

// @host localhost:3000
// @BasePath /api/v1

// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization
func main() {
	cfg := config.Get()

	db := database2.Get()
	err := database2.AutoMigrate(db)

	if err != nil {
		panic("failed to migrate db" + err.Error())
	}

	r := router.NewRouter(db)

	err = r.Run(fmt.Sprintf(":%s", cfg.ServerPort))

	if err != nil {
		panic(err)
	}
}
