package main

import (
	"database/sql"
	"log"
	"net/http"
	"os"
	"strconv"
	
	"github.com/valentinabulat/TPE3Web/pkg/views"


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

	queries := db.New(dbconn)
	
	http.HandleFunc("GET /", func(w http.ResponseWriter, r *http.Request) {
		//Obtenga todos los registros de la base de datos usando el método List... de sqlc
		productos, err := queries.ListProductos(r.Context())
		if err != nil {
			http.Error(w, "Error al obtener los productos", http.StatusInternalServerError)
			return
		}
		
		component := views.IndexPage(productos)

		// Renderice el componente completo en el http.ResponseWriter.
		err = component.Render(r.Context(),w)
		if err != nil {
			http.Error(w, "Error al renderizar la página", http.StatusInternalServerError)
			return
		}
	})

	http.HandleFunc("POST /products", func(w http.ResponseWriter, r *http.Request) {
		//Parsee los datos del formulario.
		if err := r.ParseForm(); err != nil {
			http.Error(w, "Error al parsear el formulario", http.StatusBadRequest)
			return
		}

		//Obtenga los valores del formulario.
		titulo := r.FormValue("titulo")
		descripcion := r.FormValue("descripcion")
		cantidadStr := r.FormValue("cantidad")
		
		// chequear valores vacios
		if titulo == "" || descripcion == "" || cantidadStr == "" {
			http.Error(w, "Todos los campos son obligatorios", http.StatusBadRequest)
			return
		}

		// formatear cantidad a int
		cantidad, err := strconv.Atoi(cantidadStr)
		if err != nil {
			http.Error(w, "Cantidad inválida", http.StatusBadRequest)
			return
		}

		// chequear cantidad valida
		if cantidad < 0 {
			http.Error(w, "La cantidad no puede ser negativa", http.StatusBadRequest)
			return
		}

		productoACrear := db.CreateProductoParams{
			Titulo:      titulo,                
			Descripcion: descripcion,           
			Cantidad:    int32(cantidad),       
		}

		//Inserte un nuevo registro en la base de datos usando el método Create... de sqlc.
		_, err = queries.CreateProducto(r.Context(), productoACrear)

		if err != nil {
			http.Error(w, "Error al crear el producto", http.StatusInternalServerError)
			return
		}

		// Redirija al usuario de vuelta a la página principal.
		http.Redirect(w, r, "/", http.StatusSeeOther)
	})
	
	// iniciar servidor
	log.Printf("Servidor escuchando en http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))

}
