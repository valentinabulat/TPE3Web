package main

import (
	"database/sql"
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

	// hasta aca
	// hacerlo con migraciones? para no crear la tabla cada vez que se corre el programa

	// crear insancias de queries
	queries := db.New(dbconn)
	//ctx := context.Background()

	// crear instancia de api
	api := handlers.NewAPI(queries)

	mux := http.NewServeMux()
	fs := http.FileServer(http.Dir("./static"))

	// configurar endpoints
	mux.HandleFunc("/products", api.ProductsHandler)
	mux.HandleFunc("/products/", api.ProductHandler)
	mux.Handle("/", fs)

	// iniciar servidor
	log.Println("Servidor escuchando en el puerto 8080")
	log.Fatal(http.ListenAndServe(":8080", mux))
}
