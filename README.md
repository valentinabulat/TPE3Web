#TPEWeb

Requisitos
	- Go 
	- sqlc
    - Docker compose
	- Make
	- Air
	- Hurl

Estructura del proyecto:

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
	│ └── tests/
	│	   └── requests.hurl
	├── go.mod
	├── go.sum
	├── docker-compose.yml
	├── dockerfile
	├── Makefile
	├── README.TXT
	└── sqlc.yaml


Ejecucion:
make start_db     - Inicia la base de datos PostgreSQL
make generate     - Genera código con sqlc
make air         - Inicia el servidor con air
make start_server - Inicia la API directamente con go run
make build       - Compila e inicia la API
make test        - Ejecuta los tests con Hurl
make stop        - Detiene todos los servicios
make clean       - Limpia archivos generados


Ejecutar:
Configurar la base de datos:
docker run --name some-postgres -e POSTGRES_PASSWORD=XYZ -p 5432:5432 -d docker.io/postgres

Conectarse a la base de datos:
docker exec -it some-postgres psql -h localhost -U postgres

Dentro de psql, crear la base:
CREATE DATABASE midb;
CREATE USER admin WITH PASSWORD 'admin';
GRANT ALL PRIVILEGES ON DATABASE midb TO admin;
\c midb
GRANT ALL ON SCHEMA public TO admin;
\q

Generar codigo en go:
sqlc generate

Instalar dependencias de go:
go get github.com/lib/pq

Ejecutar archivo main.go:
go run main



Autores:
- Cordoba Pablo Javier
- Bulat Maria Valentina
- Juarez Abril Valentina