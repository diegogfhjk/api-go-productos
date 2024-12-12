package main

import (
	"api-go-products/handlers"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/products", handlers.GetProducts)
	http.HandleFunc("/products/", handlers.GetProduct)

	// Iniciar el servidor
	log.Println("Servidor iniciado en http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
