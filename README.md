# TP 4

### Requisitos
	- Go 
	- sqlc
    - Docker 
	- Make


### Estructura del proyecto:

	TPEWeb/
	| ├── cmd/
	│ │   └── api/
	│ │       └── main.go
	| ├── db/
	│ │  ├── queries/
	│ │  │   └── query.sql
	│ │  └── schema/
	│ │      └── schema.sql
	│ ├── pkg/
	│ │  ├── handlers/
	│ │  │   └── products.go
	│ │  └── models/
	│ │      └── product.go
	│ ├── static/
	│ │  ├── app.js
	│ │  ├── index.html
	│ │  └── styles.css
	│ └── tests/
	│	   └── requests.hurl
	├── go.mod
	├── go.sum
	├── docker-compose.yml
	├── dockerfile
	├── Makefile
	├── README.TXT
	└── sqlc.yaml


### Comandos de ejecucion:
make start_db     - Inicia la base de datos PostgreSQL
make generate     - Genera código con sqlc
make air         - Inicia el servidor con air
make start_server - Inicia la API directamente con go run
make build       - Compila e inicia la API
make test        - Ejecuta los tests con Hurl
make stop        - Detiene todos los servicios
make clean       - Limpia archivos generados
make help		 - Ver todos los comandos disponibles

### Orden de ejecucion recomendado
Iniciar base, generar sqlc e iniciar servidor:
- make start_db
- make generate
- make start_server
Abrir en el navegador la web app (localhost:8080)
Luego para eliminar archivos autogenerados, y cerrar la base y el servidor:
- make stop
- make clean


### Autores:
- Cordoba Pablo Javier
- Bulat Maria Valentina
- Juarez Abril Valentina