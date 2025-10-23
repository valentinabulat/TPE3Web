package handlers

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"strconv"
	"strings"

	"github.com/valentinabulat/TPE3Web/internal/db"
)

// API es la struct que contendrá las dependencias, como la conexión a la DB.
type API struct {
	queries *db.Queries
}

// NewAPI es el constructor para nuestra API.
func NewAPI(q *db.Queries) *API {
	return &API{queries: q}
}

func (a *API) ProductsHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		productos, err := a.queries.ListProductos(r.Context())
		if err != nil {
			http.Error(w, "Error al obtener los productos", http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(productos)
	case http.MethodPost:
		var params db.CreateProductoParams
		err := json.NewDecoder(r.Body).Decode(&params)
		if err != nil {
			http.Error(w, "Cuerpo de la petición inválido", http.StatusBadRequest)
			return
		}

		productoCreado, err := a.queries.CreateProducto(r.Context(), params)
		if err != nil {
			http.Error(w, "Error al crear el producto", http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated) // devolver un 201 Created.
		json.NewEncoder(w).Encode(productoCreado)
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func (a *API) ProductHandler(w http.ResponseWriter, r *http.Request) {
	// Extraer ID del path
	parts := strings.Split(r.URL.Path, "/")

	if len(parts) != 3 {
		http.Error(w, "Invalid URL", http.StatusBadRequest)
		return
	}
	id, err := strconv.Atoi(parts[2])
	if err != nil {
		http.Error(w, "Invalid product ID", http.StatusBadRequest)
		return
	}

	switch r.Method {
	case http.MethodGet:
		a.getProducto(w, r, int32(id)) // Llama a la función getProducto
	case http.MethodPut:
		a.updateProducto(w, r, int32(id)) // Llama a la función updateProducto
	case http.MethodDelete:
		a.deleteProducto(w, r, int32(id)) // Llama a la función deleteProducto
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

// getProducto maneja las peticiones GET para un solo producto.
func (a *API) getProducto(w http.ResponseWriter, r *http.Request, id int32) {
	producto, err := a.queries.GetProducto(r.Context(), id)

	if err == sql.ErrNoRows {
		http.Error(w, "Product not found", http.StatusNotFound)
		return
	}
	if err != nil {
		http.Error(w, "Database error", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(producto)
}

// updateProducto maneja las peticiones PUT para actualizar un producto.
func (a *API) updateProducto(w http.ResponseWriter, r *http.Request, id int32) {
	err := a.queries.UpdateProducto(r.Context(), id)

	if err == sql.ErrNoRows {
		http.Error(w, "Product not found to update", http.StatusNotFound)
		return
	}
	if err != nil {
		http.Error(w, "Database error", http.StatusInternalServerError)
		return
	}

	//w.Header().Set("Content-Type", "application/json")
	//json.NewEncoder(w).Encode() // No es necesario devolver el producto actualizado????
	w.WriteHeader(http.StatusNoContent)
}

// deleteProducto maneja las peticiones DELETE para eliminar un producto.
func (a *API) deleteProducto(w http.ResponseWriter, r *http.Request, id int32) {
	err := a.queries.DeleteProducto(r.Context(), id)

	if err == sql.ErrNoRows {
		http.Error(w, "Product not found to delete", http.StatusNotFound)
		return
	}
	if err != nil {
		http.Error(w, "Database error", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent) // 204 No Content
}
