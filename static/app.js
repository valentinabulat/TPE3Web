const API_URL = 'http://localhost:8080/products';

//Espera a que el contenido del DOM esté completamente cargado
document.addEventListener('DOMContentLoaded', () => {
    
    const formProducto = document.getElementById('crear_lista_productos');
    const listProductos = document.getElementById('lista_productos');

    //Obtiene todos los productos de la API y los muestra en la lista.
    
    async function fetchAndRenderProductos() {
        try {
            //Pide los productos a la API
            const response = await fetch(API_URL);
            if (!response.ok) throw new Error(`Error HTTP: ${response.status}`);
            const productos = await response.json();

            //Limpia la lista actual antes de dibujar la nueva
            listProductos.innerHTML = '';

            //Si no hay productos, muestra un mensaje
            if (!productos || productos.length === 0) {
                listProductos.innerHTML = '<li>No se encontraron productos.</li>';
                return;
            }

            //Genera y agrega cada producto a la lista
            productos.forEach(producto => {
                const li = document.createElement('li');
                
                if (producto.Completed) { 
                    li.classList.add('completed');
                }

                //Crea el checkbox
                const checkbox = document.createElement('input');
                checkbox.type = 'checkbox';
                checkbox.className = 'update-checkbox'; // Clase para CSS y JS
                checkbox.checked = producto.Completed || false; // Marca el check si está completado
                checkbox.dataset.id = producto.ID; // Almacena el ID en el checkbox
                li.appendChild(checkbox);

                //Crea un span para el texto
                const textSpan = document.createElement('span');
                textSpan.innerHTML = `
                    <strong>${producto.Titulo || 'Sin nombre'}:</strong> 
                    <span>${producto.Descripcion || 'Sin descripción'}</span>
                    <span>(Cantidad: ${producto.Cantidad || 0})</span>
                `;
                li.appendChild(textSpan);

                //Crea el botón de eliminar
                const deleteButton = document.createElement('button');
                deleteButton.textContent = 'Eliminar';
                deleteButton.className = 'delete-btn'; // Clase para CSS y JS
                deleteButton.dataset.id = producto.ID; // Almacena el ID en el botón
                li.appendChild(deleteButton);
                
                listProductos.appendChild(li); //agrega el <li> completo al <ul>
            });

        } catch (error) {
            console.error('Error al obtener los productos:', error);
            listProductos.innerHTML = '<li>Error al cargar la lista.</li>';
        }
    }

    //Crear nuevo producto
    async function handleFormSubmit(event) {
         event.preventDefault(); 
        
        const tituloInput = document.getElementById('titulo');
        const descripcionInput = document.getElementById('descripcion');
        const cantidadInput = document.getElementById('cantidad');

        //Crea el objeto de datos para enviar
        const newProd = {
            Titulo: tituloInput.value,
            Descripcion: descripcionInput.value,
            Cantidad: parseInt(cantidadInput.value, 10) || 0,
            Completed: false 
        };

        try {
            //Realiza la petición POST con fetch
            const response = await fetch(API_URL, {
                method: 'POST',
                headers: {
                    'Content-Type': 'application/json'
                },
                body: JSON.stringify(newProd)
            });

            if (!response.ok) throw new Error(`Error HTTP: ${response.status}`);

            //Limpia los campos del formulario
            nameInput.value = '';
            descripcionInput.value = '';
            cantidadInput.value = '';

            //Refresca la lista para mostrar el nuevo ítem
            fetchAndRenderProductos();

        } catch (error) {
            console.error('Error al crear el producto:', error);
        }
    }

    
    //Manejo de clics
    async function handleListClick(event) {
        const target = event.target; // Elemento donde se hizo clic

        //Caso 1: Clic en el botón de ELIMINAR
        if (target.classList.contains('delete-btn')) {
            const id = target.dataset.id;
            
            if (!id || !confirm('¿Estás seguro de que quieres eliminar este producto?')) {
                return;
            }

            try {
                //Envía la petición DELETE
                const response = await fetch(`${API_URL}/${id}`, { 
                    method: 'DELETE' 
                });
                
                if (!response.ok) throw new Error(`Error HTTP: ${response.status}`);
                
                //Recarga la lista para mostrar que el ítem fue eliminado
                fetchAndRenderProductos(); 

            } catch (error) {
                console.error('Error al eliminar el producto:', error);
            }
        }

        //Caso 2: Click en el checkbox de actualizar
        if (target.classList.contains('update-checkbox')) {
            const id = target.dataset.id;
            const li = target.closest('li'); // Obtiene el <li> padre

            try {
                //Envía la petición PUT (Update)
                const response = await fetch(`${API_URL}/${id}`, {
                    method: 'PUT'
                });

                if (!response.ok) {
                    throw new Error(`Error HTTP: ${response.status}`);
                }
                
                //Actualiza la UI visualmente al instante (Optimistic Update)
                //'target.checked' nos da el *nuevo* estado del checkbox (true o false)
                li.classList.toggle('completed', target.checked);

            } catch (error) {
                console.error('Error al actualizar el producto:', error);
                //Si la API falla, revierte el checkbox a su estado anterior
                target.checked = !target.checked;
                alert("No se pudo actualizar el producto.");
            }
        }
    }
    
    //Añade el "event listener" al formulario
    formProducto.addEventListener('submit', handleFormSubmit);

    //Añade un solo listener a toda la lista para manejar los clics
    listProductos.addEventListener('click', handleListClick);

    //Carga la lista de productos al iniciar la página
    fetchAndRenderProductos();
});

