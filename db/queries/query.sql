-- name: GetProducto :one
SELECT p.ID, p.titulo, p.descripcion, l.cantidad, l.comprado
FROM lista_productos l 
JOIN producto p ON (l.ID_producto = p.ID)
WHERE l.ID = $1;

-- name: ListProductos :many
SELECT p.ID, p.titulo, p.descripcion, l.cantidad, l.comprado
FROM lista_productos l 
JOIN producto p ON (p.ID=l.ID_producto)
ORDER BY l.ID;
/*

-- name: CreateProducto :one
WITH nuevo_producto AS (
  INSERT INTO producto (titulo, descripcion)
  VALUES ($1, $2)
  RETURNING ID
)
INSERT INTO lista_productos (ID_producto, cantidad)
SELECT ID, $3
FROM nuevo_producto
RETURNING *;
*/

-- name: CreateProducto :one
WITH nuevo_producto AS (
  INSERT INTO producto (titulo, descripcion)
  VALUES ($1, $2)
  RETURNING id, titulo, descripcion -- 1. Asegúrate de retornar todo aquí
)
INSERT INTO lista_productos (ID_producto, cantidad)
SELECT id, $3
FROM nuevo_producto
RETURNING 
    (SELECT id FROM nuevo_producto) as id,          -- El ID del producto
    (SELECT titulo FROM nuevo_producto) as titulo,  -- El Título original
    (SELECT descripcion FROM nuevo_producto) as descripcion, -- La Descripción
    cantidad;                                       -- La cantidad insertada
    
-- name: UpdateProducto :one
UPDATE lista_productos
SET comprado = NOT comprado
FROM producto p
WHERE 
    lista_productos.ID_producto = p.ID AND 
    lista_productos.ID = $1
RETURNING p.ID, p.titulo, p.descripcion, lista_productos.cantidad, lista_productos.comprado;

-- name: DeleteProducto :execresult
DELETE FROM lista_productos
WHERE ID = $1;