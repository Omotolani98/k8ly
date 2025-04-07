package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"

	"github.com/Omotolani98/k8ly/services/reqly/config"
	"github.com/Omotolani98/k8ly/services/reqly/db"
	reqlyHandlers "github.com/Omotolani98/k8ly/services/reqly/handler"
)

func main() {
	godotenv.Load()
	cfg := config.Load()
	dbConn := db.Init(cfg)
	db.AutoMigrate(dbConn)

	r := mux.NewRouter()

	r.HandleFunc("/", reqlyHandlers.HandleWebhook(dbConn)).Methods("POST")
	r.HandleFunc("/inspect", reqlyHandlers.HandleInspect(dbConn)).Methods("GET")
	r.HandleFunc("/clear", reqlyHandlers.HandleClear(dbConn)).Methods("DELETE")

	fmt.Println("🚀 Reqly is running on :" + cfg.Port)
	log.Fatal(http.ListenAndServe(":"+cfg.Port, r))
}
