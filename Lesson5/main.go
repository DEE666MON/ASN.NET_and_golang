package main

import (
	"log"
	"net/http"
	"user-api/db"
	"user-api/router"
)

func main() {
	db.InitDB()
	mux := router.SetupRouter()
	log.Println("Сервер запущен на: 8080")
	log.Fatal(http.ListenAndServe(":8080", mux))
}
