package handlers

import (
	"api-go-products/data"
	"encoding/json"
	"net/http"
	"strings"
)

// GetProducts retorna todos los productos
func GetProducts(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(data.Products)
}

// GetProduct retorna un producto específico por ID
func GetProduct(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	// Extraer el ID del producto desde la URL
	id := strings.TrimPrefix(r.URL.Path, "/products/")
	if id == "" {
		http.Error(w, "ID del producto no proporcionado", http.StatusBadRequest)
		return
	}

	// Buscar el producto
	product := data.GetProductByID(id)
	if product == nil {
		http.Error(w, "Producto no encontrado", http.StatusNotFound)
		return
	}

	// Retornar el producto en formato JSON
	json.NewEncoder(w).Encode(product)
}
