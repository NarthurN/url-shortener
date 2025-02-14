package main

import (
	"fmt"

	"github.com/NarthurN/url-shortener/internal/config"
)


func main() {
	// config / cleanenv
	cfg := config.MustLoad()

	fmt.Println(cfg)

	// loger / slog

	// storage / sqlite

	// router / chi, chi_render

	// server / go
}