package router

import (
	"net/http"
	"user-api/handlers"
	"user-api/middleware"
)

func SetupRouter() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/register", middleware.CORSMiddleware(handlers.RegisterUser))
	mux.HandleFunc("/login", middleware.CORSMiddleware(handlers.LoginUser))
	mux.HandleFunc("/users", middleware.CORSMiddleware(middleware.AuthMiddleware(handlers.GetAllUsers)))
	return mux
}
