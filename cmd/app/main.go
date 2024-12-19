package main

import (
	"fmt"
	"ordent-test/config"
	database2 "ordent-test/internal/infrastructure/db"
	"ordent-test/internal/router"
)

func main() {
	cfg := config.Get()

	db := database2.Get()
	err := database2.AutoMigrate(db)

	if err != nil {
		panic("failed to migrate db" + err.Error())
	}

	r := router.NewRouter()

	err = r.Run(fmt.Sprintf(":%s", cfg.ServerPort))

	if err != nil {
		panic(err)
	}
}
