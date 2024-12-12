# **API Go: Gestión de Productos**

Esta es una API escrita en Go que proporciona detalles sobre los productos de pedidos. Ofrece endpoints para listar todos los productos y obtener información de un producto específico mediante su ID.

## **Características**
	•	Listar todos los productos disponibles.
	•	Consultar detalles de un producto por su ID.
	•	Desplegable fácilmente en Docker.

## **Requisitos**
	•	Go (v1.20 o superior)
	•	Docker
	•	Docker Compose

## **Estructura del Proyecto**
```bash
api-go-products/
├── main.go     
├── handlers/
│   ├── product_handler.go 
├── models/
│   ├── product.go
├── data/
│   ├── products.go   
├── Dockerfile          
└── docker-compose.yml    
```

## **Ejecutar con Docker**
1.	Construir y levantar los servicios:
```bash
docker-compose up --build
```

2.	La API estará disponible en http://localhost:8080.

## **Endpoints**

1. Obtener todos los productos
```bash
	•	URL: /products
	•	Método: GET
	•	Descripción: Devuelve una lista de todos los productos disponibles.
	•	Respuesta de ejemplo:
        [
        {
            "id": "1",
            "name": "iPhone 14",
            "description": "Teléfono inteligente de última generación de Apple.",
            "price": 999.99
        },
        {
            "id": "2",
            "name": "Samsung Galaxy S23",
            "description": "Teléfono de alta gama con pantalla AMOLED.",
            "price": 899.99
        }
        ]
```

2. Obtener un producto por ID
```bash
	•	URL: /products/{id}
	•	Método: GET
	•	Descripción: Devuelve los detalles de un producto específico.
	•	Ejemplo de solicitud: GET /products/1
	•	Respuesta de ejemplo:
        {
        "id": "1",
        "name": "iPhone 14",
        "description": "Teléfono inteligente de última generación de Apple.",
        "price": 999.99
        }
```