package handlers

func main () {
	// Handler
	http.HandleFunc("/products", productsHandler)
	http.HandleFunc("/products/", productHandler)

	// Iniciar servidor


	// 
	func productsHandler(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
			case http.MethodGet:
				GetProducts(w, r)
			case http.MethodPost:
				CreateProduct(w, r)
			default:http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	}

	func productHandler(w http.ResponseWriter, r *http.Request) {
		// Extraer ID del path
		parts := strings.Split(r.URL.Path, "/")

		if len(parts) != 3 {
			http.Error(w, "Invalid URL", http.StatusBadRequest)
			return
		}
		id, err := strconv.Atoi(parts[2])
		if err != nil {
			http.Error(w, "Invalid product ID",http.StatusBadRequest)
			return
		}
		switch r.Method {
			case http.MethodGet:
				getProduct(w, r, id)
			case http.MethodPut:
				updateProduct(w, r, id)
			case http.MethodDelete:
				deleteProduct(w, r, id)
			default:
				http.Error(w, "Method not allowed",http.StatusMethodNotAllowed)
		}
	}

	func getProducts(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		products, err := ListProductos()
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		json.NewEncoder(w).Encode(products)
	}

	func createProduct(w http.ResponseWriter, r *http.Request) {

		var req db.CreateProductoParams
		err := json.NewDecoder(r.Body).Decode(&req)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		// validar que titulo no sea vacio ??

		producto, error := queries.CreateProducto(r.Context(), req) //titulo, descripcion,cantidad 

		w.Header().Set("Content-Type", "application/json")
		WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(producto)
	}

	func getProduct(w http.ResponseWriter, r *http.Request, id int) {
		product, err := GetProducto(id) // llama al metodo de sqlc
		if err != nil {http.Error(w, err.Error(), http.StatusNotFound)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(product)
	}



}