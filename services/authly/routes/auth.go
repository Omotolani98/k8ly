// services/authly/routes/auth_routes.go
package routes

import (
	"github.com/gorilla/mux"
	"gorm.io/gorm"
	"net/http"

	handler "github.com/Omotolani98/k8ly/services/authly/handler"
)

func RegisterAuthRoutes(router *mux.Router, db *gorm.DB) {
	authHandler := handler.NewAuthHandler(db)

	router.HandleFunc("/auth/register", authHandler.Register).Methods(http.MethodPost)
	router.HandleFunc("/auth/login", authHandler.Login).Methods(http.MethodPost)
}

