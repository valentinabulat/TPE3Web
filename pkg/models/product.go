package models

type Producto struct {
	ID          int    `json:"id"`
	Titulo      string `json:"titulo"`
	Descripcion string `json:"descripcion"`
}

type ListaProducto struct {
	ID         int       `json:"id"`
	IDProducto int       `json:"id_producto"`
	Cantidad   int       `json:"cantidad"`
	Comprado   bool      `json:"comprado"`
	Producto   *Producto `json:"producto,omitempty"` // Relación opcional
}
