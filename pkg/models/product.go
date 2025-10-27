package models

type Producto struct {
	ID          int32  `json:"id"`
	Titulo      string `json:"titulo"`
	Descripcion string `json:"descripcion"`
}

type ListaProducto struct {
	ID         int32     `json:"id"`
	IDProducto int32     `json:"id_producto"`
	Cantidad   int32     `json:"cantidad"`
	Comprado   bool      `json:"comprado"`
	Producto   *Producto `json:"producto,omitempty"` // Relación opcional
}
