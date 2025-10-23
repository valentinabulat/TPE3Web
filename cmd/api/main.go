package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/valentinabulat/TPE3Web/pkg/handlers"

	_ "github.com/lib/pq"
	"github.com/valentinabulat/TPE3Web/internal/db"
)

func main() {

	// Conectar a la base de datos
	connStr := "user=admin password=admin dbname=midb sslmode=disable"
	dbconn, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatalf("failed to connect to DB: %v", err)
	}
	defer dbconn.Close()

	schemaSQL, err := os.ReadFile("db/schema/schema.sql")
	if err != nil {
		log.Fatalf("failed to read schema.sql: %v", err)
	}

	_, err = dbconn.Exec(string(schemaSQL))
	if err != nil {
		log.Fatalf("failed to execute schema: %v", err)
	}
	log.Println("Schema ejecutado correctamente")

	// crear insancias de queries
	queries := db.New(dbconn)
	ctx := context.Background()

	// crear instancia de api
	api := handlers.NewAPI(queries)

	mux := http.NewServeMux()
	// configurar endpoints
	mux.HandleFunc("/products", api.productsHandler)
	mux.HandleFunc("/products/", api.productHandler)

	// iniciar servidor
	fmt.Println("Servidor escuchando en el puerto 8080")
	log.Fatal(http.ListenAndServe(":8080", mux))
}
