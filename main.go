package main

import (
	"fmt"
	"ordent-test/internal/config"
	"ordent-test/internal/database"
	"ordent-test/internal/router"
)

func main() {
	cfg := config.Get()

	db := database.Get()
	err := database.AutoMigrate(db)

	if err != nil {
		panic("failed to migrate database" + err.Error())
	}

	r := router.NewRouter()

	err = r.Run(fmt.Sprintf(":%s", cfg.ServerPort))

	if err != nil {
		panic(err)
	}
}
