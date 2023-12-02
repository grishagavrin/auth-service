package main

import (
	"fmt"
	"sso/cmd/internal/config"
)

func main() {
	cfg := config.MustLoad()

	fmt.Println(cfg)

	// TODO: инициализировать логгер

	// TODO: инициализировать приложение (app)
}
