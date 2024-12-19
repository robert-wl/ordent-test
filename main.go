package main

import (
	"fmt"
	"ordent-test/internal/config"
	"ordent-test/internal/router"
)

func main() {
	cfg := config.Load()

	r := router.NewRouter()

	err := r.Run(fmt.Sprintf(":%s", cfg.ServerPort))

	if err != nil {
		panic(err)
	}
}
