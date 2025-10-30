package handlers

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"strings"

	"github.com/valentinabulat/TPE3Web/internal/db"
	"github.com/valentinabulat/TPE3Web/pkg/models"
)

// API es la struct que contendrá las dependencias, como la conexión a la DB.
type API struct {
	queries *db.Queries
}

// constructor para la struct API
func NewAPI(q *db.Queries) *API {
	return &API{queries: q}
}

// manejar peticiones a /products
func (a *API) ProductsHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		a.getProductos(w, r)
	case http.MethodPost:
		a.createProduct(w, r)
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

// manejar peticiones a /products/{id}
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

// getProductos maneja las peticiones GET para listar todos los productos
func (a *API) getProductos(w http.ResponseWriter, r *http.Request) {
	productos, err := a.queries.ListProductos(r.Context())
	if err != nil {
		log.Printf("Error en ListProductos: %v", err)
		http.Error(w, "Error al obtener los productos", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(productos)
}

// createProduct maneja las peticiones POST para crear un nuevo producto.
func (a *API) createProduct(w http.ResponseWriter, r *http.Request) {
	var params db.CreateProductoParams

	err := json.NewDecoder(r.Body).Decode(&params) // decodifica el body de la petición, aca json es sensible a mayusculas y minusculas
	if err != nil {
		log.Printf("Error en Decode: %v", err)
		http.Error(w, "Cuerpo de la petición inválido", http.StatusBadRequest)
		return
	}

	if params.Titulo == "" || params.Descripcion == "" || params.Cantidad <= 0 { // chequear que sean datos validos
		http.Error(w, "Faltan campos obligatorios", http.StatusBadRequest)
		return
	}

	listaProductoFlat, err := a.queries.CreateProducto(r.Context(), params)
	if err != nil {
		http.Error(w, "Error al crear el producto", http.StatusInternalServerError)
		return
	}

	// hay que crear estructuras auxiliares para poder enviar la respuesta completa
	productoDetalles := models.Producto{
		ID:          listaProductoFlat.IDProducto, // De Fuente 2
		Titulo:      params.Titulo,                // De Fuente 1
		Descripcion: params.Descripcion,           // De Fuente 1
	}

	// Creamos la respuesta JSON completa
	respuestaCompleta := models.ListaProducto{
		ID:         listaProductoFlat.ID,            // De Fuente 2
		IDProducto: listaProductoFlat.IDProducto,    // De Fuente 2
		Cantidad:   listaProductoFlat.Cantidad,      // De Fuente 2
		Comprado:   listaProductoFlat.Comprado.Bool, // De Fuente 2
		Producto:   &productoDetalles,               // El objeto "cosido"
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated) // devolver un 201 Created.
	err = json.NewEncoder(w).Encode(respuestaCompleta)
	if err != nil {
		log.Printf("Error en Encode: %v", err)
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
// definimos 'actualizar' como cambiar el estado de 'comprado'
func (a *API) updateProducto(w http.ResponseWriter, r *http.Request, id int32) {
	var prodAct db.UpdateProductoRow
	prodAct, err := a.queries.UpdateProducto(r.Context(), id)

	if err == sql.ErrNoRows {
		http.Error(w, "Product not found to update", http.StatusNotFound) //404
		return
	}
	if err != nil {
		http.Error(w, "Database error", http.StatusInternalServerError) //500
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK) // 200 OK
	json.NewEncoder(w).Encode(prodAct)
}

// deleteProducto maneja las peticiones DELETE para eliminar un producto.
func (a *API) deleteProducto(w http.ResponseWriter, r *http.Request, id int32) {
	result, err := a.queries.DeleteProducto(r.Context(), id) // result nos sirve para ver cuantas filas fueron afectadas
	if err != nil {
		http.Error(w, "Database error", http.StatusInternalServerError)
		return
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		http.Error(w, "Database error", http.StatusInternalServerError)
		return
	}
	if rowsAffected == 0 { // si no se afectó ninguna fila, el producto no existía
		http.Error(w, "Product not found to delete", http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusNoContent) // 204 No Content
}
