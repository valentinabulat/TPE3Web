TPEWeb

Requisitos
	- Go 
	- sqlc
    - Docker

Estructura del proyecto:

	TPEWeb/
	| ├── db
	│   ├── queries
	│   │   └── query.sql
	│   └── schema
	│       └── schema.sql
	├── go.mod
	├── go.sum
	├── main.go
	├── README.TXT
	├── sqlc.yaml
	└── static
	    └── index.html

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