package handlers

import (
    "database/sql"
    "net/http"
    "strconv"

    "github.com/valentinabulat/TPE3Web/internal/db"
    "github.com/valentinabulat/TPE3Web/pkg/views"
)

// Handler es una estructura que contiene las dependencias necesarias para los handlers HTTP.
type Handler struct {
    Queries *db.Queries
}

// NewHandler crea e inicializa un nuevo Handler.
func NewHandler(queries *db.Queries) *Handler {
    return &Handler{
        Queries: queries,
    }
}

// GetIndex maneja la ruta GET /
func (h *Handler) GetIndex(w http.ResponseWriter, r *http.Request) {
    //Obtiene todos los registros de la base de datos usando el método List de sqlc
    productos, err := h.Queries.ListProductos(r.Context())
    if err != nil {
        http.Error(w, "Error al obtener los productos", http.StatusInternalServerError)
        return
    }

    component := views.IndexPage(productos)

    // Renderiza el componente completo en el http.ResponseWriter.
    err = component.Render(r.Context(), w)
    if err != nil {
        http.Error(w, "Error al renderizar la página", http.StatusInternalServerError)
        return
    }
}

// PostProduct maneja la ruta POST /products
func (h *Handler) PostProduct(w http.ResponseWriter, r *http.Request) {
    // Parsea los datos del formulario.
    if err := r.ParseForm(); err != nil {
        http.Error(w, "Error al parsear el formulario", http.StatusBadRequest)
        return
    }

    // Obtiene los valores del formulario.
    titulo := r.FormValue("titulo")
    descripcion := r.FormValue("descripcion")
    cantidadStr := r.FormValue("cantidad")

    // Chequea valores vacíos
    if titulo == "" || descripcion == "" || cantidadStr == "" {
        http.Error(w, "Todos los campos son obligatorios", http.StatusBadRequest)
        return
    }

    // Formatea cantidad a int
    cantidad, err := strconv.Atoi(cantidadStr)
    if err != nil {
        http.Error(w, "Cantidad inválida", http.StatusBadRequest)
        return
    }

    // Chequea cantidad válida
    if cantidad < 0 {
        http.Error(w, "La cantidad no puede ser negativa", http.StatusBadRequest)
        return
    }

    productoACrear := db.CreateProductoParams{
        Titulo:      titulo,
        Descripcion: descripcion,
        Cantidad:    int32(cantidad),
    }

    // Inserta un nuevo registro en la base de datos usando el método Create de sqlc.
    productoCreado, err := h.Queries.CreateProducto(r.Context(), productoACrear)
    if err != nil {
        http.Error(w, "Error al crear el producto", http.StatusInternalServerError)
        return
    }
    productoAMostrar := db.ListProductosRow{
        ID:          productoCreado.ID, // El ID nuevo
        Titulo:      productoCreado.Titulo,
        Descripcion: productoCreado.Descripcion,
        Cantidad:    productoCreado.Cantidad,
    }

    // Check for HTMX request header
    if r.Header.Get("HX-Request") == "true" {
        // Renderiza SOLO LA FILA (ProductRow), no la lista entera
        component := views.ProductRow(productoAMostrar)

        // HTMX toma este <tr> y lo pone al final del <tbody>
        component.Render(r.Context(), w)
        return
    }

    // Si no es HTMX, volvemos a consultar la lista completa y renderizarla.
    productos, err := h.Queries.ListProductos(r.Context())
    if err != nil {
        http.Error(w, "Error al obtener los productos", http.StatusInternalServerError)
        return
    }
    views.ProductList(productos).Render(r.Context(), w)
}

// DeleteProduct maneja la ruta DELETE /products/{id}
func (h *Handler) DeleteProduct(w http.ResponseWriter, r *http.Request) {
    // Obtiene el ID del producto de la URL.
    idStr := r.PathValue("id")
    id, err := strconv.Atoi(idStr)
    if err != nil {
        http.Error(w, "ID inválido", http.StatusBadRequest)
        return
    }

    // Elimina el registro de la base de datos usando el método Delete de sqlc.
    _, err = h.Queries.DeleteProducto(r.Context(), int32(id))
    if err != sql.ErrNoRows && err != nil { // Ignoramos ErrNoRows si se intenta borrar algo que no existe
        http.Error(w, "Error al eliminar el producto", http.StatusInternalServerError)
        return
    }

    w.WriteHeader(http.StatusOK) // 200 OK
}