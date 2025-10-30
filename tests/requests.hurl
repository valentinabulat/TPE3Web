# tests para cada endpoint de nuestra API

# 1. GET todo los productos
GET http://localhost:8080/products
HTTP 200
[Captures]
initial_count: jsonpath "$[*]" count # cantidad incial de productos

# 2. Crear un nuevo producto
POST http://localhost:8080/products
{
    "titulo": "Producto de Test",
    "descripcion": "Este es un producto para testing",
    "cantidad": 1
}
HTTP 201
[Captures]
product_id: jsonpath "$.id_producto" # el id es autogenerado, es el unico momento que podemos capturarlo
[Asserts]
jsonpath "$.producto.titulo" == "Producto de Test"
jsonpath "$.producto.descripcion" == "Este es un producto para testing"

# 3. Verificar que el producto fue creado
GET http://localhost:8080/products/{{product_id}}
HTTP 200
[Asserts]
jsonpath "$.ID" == {{product_id}} # JSON sensible a mayus, arreglar desde el codigo
jsonpath "$.Titulo" == "Producto de Test"
jsonpath "$.Descripcion" == "Este es un producto para testing"

# 4. Verificar lista actualizada
GET http://localhost:8080/products
HTTP 200
[Captures]
current_count: jsonpath "$[*]" count # captura el nuevo conteo
[Asserts]
variable "current_count" > {{initial_count}} # chequear que sea mayor al inicial

# 5. Actualizar el producto
PUT http://localhost:8080/products/{{product_id}}
HTTP 200 # devuelve el producto actualizado

# 6. Verificar que el producto fue actualizado
GET http://localhost:8080/products/{{product_id}}
HTTP 200 # deberia mostrar comprado = true

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