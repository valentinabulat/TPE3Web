FROM golang:1.21-alpine AS builder

WORKDIR /app

# Copiamos los archivos de dependencias PRIMERO.
COPY go.mod go.sum ./

# Descargamos todas las dependencias.
RUN go mod download

# Ahora, copiamos el resto del código fuente de tu proyecto.
COPY . .

# Compilamos tu aplicación.
RUN CGO_ENABLED=0 go build -o /app/server ./cmd/api/


FROM alpine:latest

# Establecemos el directorio de trabajo.
WORKDIR /app


COPY --from=builder /app/server .

EXPOSE 8080

CMD [ "./server" ]