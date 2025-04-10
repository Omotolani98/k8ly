package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/Omotolani98/k8ly/services/logly/config"
	"github.com/Omotolani98/k8ly/services/logly/db"
	"github.com/Omotolani98/k8ly/services/logly/handler"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

func main() {
	// Load .env file
	_ = godotenv.Load()
	cfg := config.Load()
	dbConn := db.Init(cfg)
	db.AutoMigrate(dbConn)

	h := handler.NewLogHandler(dbConn)
	r := mux.NewRouter()
	r.HandleFunc("/logs", h.CreateLog).Methods("POST")
	r.HandleFunc("/logs", h.GetLogs).Methods("GET")

	fmt.Printf("🚀 Logly running on port %s\n", cfg.Port)
	log.Fatal(http.ListenAndServe(":"+cfg.Port, r))
}
