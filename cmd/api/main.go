package main

import (
	"database/sql"
	"log"
	"net/http"
	"os"

	_ "github.com/lib/pq"
	"github.com/valentinabulat/TPE3Web/internal/db"
	"github.com/valentinabulat/TPE3Web/internal/handlers"
)

func main() {

	// Conectar a la base de datos
	connStr := "user=admin password=admin dbname=midb sslmode=disable"
	dbconn, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatalf("failed to connect to DB: %v", err)
	}
	defer dbconn.Close()

	// desde aca

	schemaSQL, err := os.ReadFile("db/schema/schema.sql")
	if err != nil {
		log.Fatalf("failed to read schema.sql: %v", err)
	}

	_, err = dbconn.Exec(string(schemaSQL))
	if err != nil {
		log.Fatalf("failed to execute schema: %v", err)
	}
	log.Println("Schema ejecutado correctamente")

	queries := db.New(dbconn)

	// Crear la instancia del Handler con las dependencias (queries)
    h := handlers.NewHandler(queries)

    // Definición de Handlers
    http.HandleFunc("GET /", h.GetIndex)
    http.HandleFunc("POST /products", h.PostProduct)
    http.HandleFunc("DELETE /products/{id}", h.DeleteProduct)

	// iniciar servidor
	log.Printf("Servidor escuchando en http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))

}
