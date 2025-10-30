# Variables
APP_NAME = tpe3web-api
MAIN_PATH = cmd/api/main.go

# Target por defecto
all: build

# Comandos para la base de datos
start_db:
	@echo "Iniciando base de datos PostgreSQL..."
	docker compose up -d db

generate:
	@echo "Generando codigo con sqlc..."
	sqlc generate

# Desarrollo con hot-reload usando Air
air:
	@echo "Iniciando servidor con Air..."
	air

# Iniciar la API directamente
start_server: start_db
	@echo "Iniciando API..."
	go run $(MAIN_PATH) &

# Iniciar en modo producción
build: generate start_db
	@echo "Compilando e iniciando API..."
	go build -o $(APP_NAME) $(MAIN_PATH)
	./$(APP_NAME) &

# Ejecutar tests con Hurl
test: build
	@echo "Ejecutando tests con Hurl..."
	@echo "Asegurate de que la API este corriendo antes de ejecutar los tests"
	hurl --test tests/requests.hurl

# Detener todos los servicios. Eliminar contenedores y volúmenes (si no se elimina volúmenes, usa la imagen que ya existia)
stop:
	@echo "Deteniendo servicios..."
	@-pkill $(APP_NAME)
	@-docker compose down -v 
	@echo "Servicios detenidos"

# Limpiar archivos generados
clean:
	@echo "Limpiando archivos generados..."
	rm -f $(APP_NAME)
	rm -rf internal/db/*
	@echo "Limpieza completada"

# Ayuda
help:
	@echo "Comandos disponibles:"
	@echo "  make start_db     - Inicia la base de datos PostgreSQL"
	@echo "  make generate     - Genera código con sqlc"
	@echo "  make air         - Inicia el servidor con hot-reload (desarrollo)"
	@echo "  make start_server - Inicia la API directamente con go run"
	@echo "  make build       - Compila e inicia la API en modo producción"
	@echo "  make test        - Ejecuta los tests con Hurl"
	@echo "  make stop        - Detiene todos los servicios"
	@echo "  make clean       - Limpia archivos generados"
