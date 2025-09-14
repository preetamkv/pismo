package main

import (
	"fmt"
	"net/http"

	"github.com/preetamkv/pismo/internal/app/pismo"
	"github.com/preetamkv/pismo/internal/app/pismo/accounts"
	"github.com/preetamkv/pismo/internal/app/pismo/transactions"
	"github.com/preetamkv/pismo/internal/pkg/config"
	"github.com/preetamkv/pismo/internal/pkg/middleware"
	"github.com/preetamkv/pismo/internal/pkg/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"github.com/go-chi/chi/v5"
)

func main() {
	cfg, err := config.LoadConfig("settings.json")
	if err != nil {
		fmt.Println("failed to fetch settings:", err)
	}

	// Create Connection string
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=%s",
		cfg.Database.Host,
		cfg.Database.User,
		cfg.Database.Password,
		cfg.Database.DBName,
		cfg.Database.Port,
		cfg.Database.SSLMode,
	)

	//Establish connectivity with the DB
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println("failed to connect database:", err)
		return
	}

	// Ensuring to close the connections to DB while closing the app
	sqlDB, err := db.DB()
	if err != nil {
		fmt.Println("failed to connect database:", err)
		return
	}
	defer sqlDB.Close()

	// Select Schema and auto generate tables
	db.Exec(fmt.Sprintf("CREATE SCHEMA IF NOT EXISTS %s;", cfg.App.Schema))
	db.Exec(fmt.Sprintf("SET search_path TO %s;", cfg.App.Schema))
	db.AutoMigrate(&models.Account{}, &models.Transaction{})

	// Storing the DB Client in App so that we can reuse it again and again.
	app := &pismo.App{DB: db}

	// Create a router for the app with JSON only check
	r := chi.NewRouter()
	r.Use(middleware.JSONOnly)

	// Create sub routers for each namespace
	r.Mount("/accounts", accounts.Routes(app))
	r.Mount("/transactions", transactions.Routes(app))

	// Start on port 8080
	fmt.Println("Starting server on port 8080")
	err = http.ListenAndServe(":8080", r) // listen on port 8080
	if err != nil {
		fmt.Println("Error starting server:", err)
	}
}
