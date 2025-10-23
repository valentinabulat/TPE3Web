package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"
	"strings"

	"TPE3WEB.com/TPE3WEB/internal/db"
)

// API es la struct que contendrá las dependencias, como la conexión a la DB.
type API struct {
	queries *db.Queries
}

// NewAPI es el constructor para nuestra API.
func NewAPI(q *db.Queries) *API {
	return &API{queries: q}
}

func (a *API) productsHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		a.ListProductos(w, r)
	case http.MethodPost:
		a.CreateProducto(w, r)
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func (a *API) productHandler(w http.ResponseWriter, r *http.Request) {
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
		a.getProducto(w, r, id)
	case http.MethodPut:
		a.updateProducto(w, r, id)
	case http.MethodDelete:
		a.deleteProducto(w, r, id)
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func (a *API) getProducts(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	products, err := db.ListProductos()
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	json.NewEncoder(w).Encode(products)
}

func (a *API) createProduct(w http.ResponseWriter, r *http.Request) {

	var req db.CreateProductoParams
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	// validar que titulo no sea vacio ??

	producto, error := queries.CreateProducto(r.Context(), req) //titulo, descripcion,cantidad

	w.Header().Set("Content-Type", "application/json")
	WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(producto)
}

func (a *API) getProduct(w http.ResponseWriter, r *http.Request, id int) {
	product, err := a.GetProducto(id) // llama al metodo de sqlc
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(product)
}
