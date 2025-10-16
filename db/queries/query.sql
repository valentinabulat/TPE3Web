-- name: GetProducto :one
SELECT id, titulo, descripcion, cantidad, comprado
FROM productos
WHERE id = $1;

-- name: ListProductos :many
SELECT id, titulo, descripcion, cantidad, comprado
FROM productos
ORDER BY comprado, id;

-- name: CreateProducto :one
INSERT INTO productos (titulo, descripcion, cantidad)
VALUES ($1, $2, $3)
RETURNING id, titulo, descripcion, cantidad, comprado;

-- name: UpdateProducto :exec
UPDATE productos
SET comprado = true
WHERE id = $1;

-- name: DeleteProducto :exec
DELETE FROM productos
WHERE id = $1;