package main

import (
	"log"
	"net/http"

	"github.com/username/auth-service/internal/auth/handlers"
)

func main() {
	// Inisialisasi router
	router := handlers.NewRouter()

	// Menetapkan handler dan rute HTTP
	http.Handle("/", router)

	// Menjalankan server HTTP pada port tertentu
	log.Fatal(http.ListenAndServe(":8080", nil))
}
