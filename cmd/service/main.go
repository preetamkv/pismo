package main

import (
	"fmt"
	"net/http"

	"github.com/preetamkv/pismo/internal/app/pismo"
	"github.com/preetamkv/pismo/internal/app/pismo/accounts"
	"github.com/preetamkv/pismo/internal/app/pismo/transactions"
	"github.com/preetamkv/pismo/internal/pkg/middleware"
	"github.com/preetamkv/pismo/internal/pkg/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"github.com/go-chi/chi/v5"
)

func main() {
	dsn := "host=c-postgrescluster.dlksek2fawx2u2.postgres.cosmos.azure.com port=5432 dbname=citus user=citus password=H@Sh1CoR3! sslmode=require"

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

	db.Exec("CREATE SCHEMA IF NOT EXISTS pismo;")
	db.Exec("SET search_path TO pismo;")
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
