package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/Omotolani98/k8ly/services/authly/config"
	"github.com/Omotolani98/k8ly/services/authly/db"
	"github.com/Omotolani98/k8ly/services/authly/routes"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

func main() {
	// Load .env file
	_ = godotenv.Load()
	cfg := config.Load()
	dbConn := db.Init(cfg)
	db.AutoMigrate(dbConn)

	r := mux.NewRouter()

	routes.RegisterAuthRoutes(r, dbConn)

	fmt.Printf("🚀 Authly running on port %s\n", cfg.Port)
	log.Fatal(http.ListenAndServe(":"+cfg.Port, r))
}
