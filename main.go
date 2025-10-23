package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"os"

	"TPEWeb.com/TPEWeb/internal/db"
	_ "github.com/lib/pq"
)

func main() {

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

	queries := db.New(dbconn)
	ctx := context.Background()

	createdProduct, err := queries.CreateProducto(ctx, db.CreateProductoParams{
		Titulo:      "Manteca",
		Descripcion: sql.NullString{String: "Lacteo", Valid: true},
		Cantidad:    1,
	})
	if err != nil {
		log.Fatalf("failed to create product: %v", err)
	}
	fmt.Printf("Created Product: %+v\n", createdProduct)

	product, err := queries.GetProducto(ctx, createdProduct.ID) // Read One
	if err != nil {
		log.Fatalf("failed to get product: %v", err)
	}
	fmt.Printf("Retrieved product: %+v\n", product)

	products, err := queries.ListProductos(ctx) // Read Many
	if err != nil {
		log.Fatalf("failed to list products: %v", err)
	}
	fmt.Printf("All products: %+v\n", products)

	err = queries.UpdateProducto(ctx, createdProduct.ID)
	if err != nil {
		log.Fatalf("failed to update product: %v", err)
	}
	fmt.Println("Product updated successfully")

	updatedProduct, err := queries.GetProducto(ctx, createdProduct.ID)
	if err != nil {
		log.Fatalf("failed to get updated product: %v", err)
	}
	fmt.Printf("Updated product: %+v\n", updatedProduct)
}
