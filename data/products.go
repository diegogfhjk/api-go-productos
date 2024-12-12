package data

import "api-go-products/models"

var Products = []models.Product{
	{ID: "1", Name: "iPhone 14", Description: "Teléfono inteligente de última generación de Apple.", Price: 999.99},
	{ID: "2", Name: "Samsung Galaxy S23", Description: "Teléfono de alta gama con pantalla AMOLED.", Price: 899.99},
	{ID: "3", Name: "MacBook Pro 16\"", Description: "Portátil potente con chip M1 Pro de Apple.", Price: 2499.00},
	{ID: "4", Name: "Dell XPS 13", Description: "Ultrabook con pantalla táctil 4K y procesador Intel i7.", Price: 1499.99},
	{ID: "5", Name: "Sony WH-1000XM5", Description: "Auriculares inalámbricos con cancelación de ruido líder en su clase.", Price: 349.99},
	{ID: "6", Name: "Logitech MX Master 3", Description: "Ratón ergonómico inalámbrico con funciones avanzadas.", Price: 99.99},
	{ID: "7", Name: "GoPro Hero 11", Description: "Cámara de acción resistente al agua con grabación 5K.", Price: 499.99},
	{ID: "8", Name: "Nintendo Switch OLED", Description: "Consola de videojuegos híbrida con pantalla OLED de 7 pulgadas.", Price: 349.99},
	{ID: "9", Name: "Kindle Paperwhite", Description: "Lector de libros electrónicos con pantalla retroiluminada.", Price: 139.99},
	{ID: "10", Name: "Dyson V15 Detect", Description: "Aspiradora inalámbrica con sensor láser para suciedad.", Price: 699.99},
}

// GetProductByID busca un producto por su ID
func GetProductByID(id string) *models.Product {
	for _, product := range Products {
		if product.ID == id {
			return &product
		}
	}
	return nil
}
