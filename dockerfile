# -------------------------------------------------------------------- #
# ---- ETAPA 1: El Constructor (Builder) ----
# -------------------------------------------------------------------- #
# Usamos una imagen oficial de Go con Alpine Linux, que es ligera.
FROM golang:1.21-alpine AS builder

# Establecemos el directorio de trabajo dentro del contenedor.
WORKDIR /app

# Copiamos los archivos de dependencias PRIMERO.
COPY go.mod go.sum ./

# Descargamos todas las dependencias.
RUN go mod download

# Ahora, copiamos el resto del código fuente de tu proyecto.
COPY . .

# Compilamos tu aplicación.
# - CGO_ENABLED=0 crea un binario estático, ideal para contenedores.
# -o /app/server es el archivo de salida (nuestro ejecutable).
# - ./cmd/api/ es la ruta a la carpeta que contiene tu main.go. ¡Ajústala si es diferente!
RUN CGO_ENABLED=0 go build -o /app/server ./cmd/api/


# -------------------------------------------------------------------- #
# ---- ETAPA 2: La Imagen Final ----
# -------------------------------------------------------------------- #
# Empezamos desde una imagen base súper pequeña y segura.
FROM alpine:latest

# Establecemos el directorio de trabajo.
WORKDIR /app

# Copiamos SOLAMENTE el archivo ejecutable compilado desde la etapa "builder".
# Esta es la magia del multi-stage build.
COPY --from=builder /app/server .

# Le informamos a Docker que nuestra aplicación escucha en el puerto 8080.
# Esto es más que nada informativo, el mapeo real se hace en docker-compose.yml.
EXPOSE 8080

# Este es el comando que se ejecutará cuando el contenedor inicie.
# Simplemente corre el servidor que compilamos.
CMD [ "./server" ]