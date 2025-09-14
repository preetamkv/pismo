package main

import (
	"fmt"
	"net/http"

	"github.com/preetamkv/pismo/internal/app/pismo"
	"github.com/preetamkv/pismo/internal/app/pismo/accounts"
	"github.com/preetamkv/pismo/internal/app/pismo/transactions"
	"github.com/preetamkv/pismo/internal/pkg/middleware"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"

	"github.com/go-chi/chi/v5"
)

func main() {
	db, err := gorm.Open(sqlite.Open("gorm.db"), &gorm.Config{})
	if err != nil {
		fmt.Println("failed to connect database:", err)
		return
	}

	sqlDB, err := db.DB()
	if err != nil {
		fmt.Println("failed to connect database:", err)
		return
	}
	defer sqlDB.Close()

	app := &pismo.App{DB: db}

	r := chi.NewRouter()
	r.Use(middleware.JSONOnly)

	r.Mount("/accounts", accounts.Routes(app))
	r.Mount("/transactions", transactions.Routes(app))

	fmt.Println("Starting server on port 8080")
	err = http.ListenAndServe(":8080", r) // listen on port 8080
	if err != nil {
		fmt.Println("Error starting server:", err)
	}
}
