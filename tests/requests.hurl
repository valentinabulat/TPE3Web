
# 1. GET inicial para ver productos existentes
GET http://localhost:8080/products
HTTP 200
[Captures]
initial_count: jsonpath "$[*]" count

# 2. Crear un nuevo producto
POST http://localhost:8080/products
{
    "titulo": "Producto de Test",
    "descripcion": "Este es un producto para testing",
    "cantidad": 1
}
HTTP 201
[Captures]
product_id: jsonpath "$.id_producto"
#[Asserts]
#jsonpath "$.producto.titulo" == "Producto de Test"
#jsonpath "$.producto.descripcion" == "Este es un producto para testing"

# 3. Verificar que el producto fue creado
GET http://localhost:8080/products/{{product_id}}
HTTP 200
[Asserts]
jsonpath "$.id" == {{product_id}}
jsonpath "$.titulo" == "Producto de Test"
jsonpath "$.descripcion" == "Este es un producto para testing"

# 4. Verificar lista actualizada
GET http://localhost:8080/products
HTTP 200
[Captures]
# Captura el nuevo conteo
current_count: jsonpath "$[*]" count
[Asserts]
variable "current_count" > {{initial_count}}

# 5. Actualizar el producto
PUT http://localhost:8080/products/{{product_id}}
HTTP 204

# 6. Verificar que el producto fue actualizado
GET http://localhost:8080/products/{{product_id}}
HTTP 200

# 7. Eliminar el producto creado para tests
DELETE http://localhost:8080/products/{{product_id}}
HTTP 204

# 8. Verificar que el producto fue eliminado
GET http://localhost:8080/products/{{product_id}}
HTTP 404

# 9. Verificar que la lista volvió a su estado original
GET http://localhost:8080/products
HTTP 200
[Asserts]
jsonpath "$[*]" count == {{initial_count}}

# Tests de error
# 10. Intentar obtener un producto que no existe
GET http://localhost:8080/products/999999
HTTP 404

# 11. Intentar crear un producto con datos inválidos
POST http://localhost:8080/products
{
    "titulo_invalido": "Test"
}
HTTP 400

# 12. Intentar actualizar un producto que no existe
PUT http://localhost:8080/products/999999
HTTP 404

# 13. Intentar eliminar un producto que no existe
DELETE http://localhost:8080/products/999999
HTTP 404