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


-- name: UpdateProducto :exec
UPDATE lista_productos
SET comprado = true
WHERE ID = $1;

-- name: DeleteProducto :exec
DELETE FROM lista_productos
WHERE ID = $1;