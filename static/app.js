const API_URL = 'http://localhost:8080/products';

//Espera a que el contenido del DOM esté completamente cargado
document.addEventListener('DOMContentLoaded', () => {
    
    const entityForm = document.getElementById('crear_lista_productos');
    const entityList = document.getElementById('lista_productos');

    //Obtiene todos los productos de la API y los muestra en la lista.
    
    async function fetchAndRenderEntities() {
        try {
            //Pide los productos a la API
            const response = await fetch(API_URL);
            if (!response.ok) throw new Error(`Error HTTP: ${response.status}`);
            const entities = await response.json();

            //Limpia la lista actual antes de dibujar la nueva
            entityList.innerHTML = '';

            //Si no hay entidades, muestra un mensaje
            if (!entities || entities.length === 0) {
                entityList.innerHTML = '<li>No se encontraron productos.</li>';
                return;
            }

            //Genera y agrega cada entidad a la lista
            entities.forEach(entity => {
                const li = document.createElement('li');
                
                if (entity.Completed) { 
                    li.classList.add('completed');
                }

                //Crea el Checkbox
                const checkbox = document.createElement('input');
                checkbox.type = 'checkbox';
                checkbox.className = 'update-checkbox'; // Clase para CSS y JS
                checkbox.checked = entity.Completed || false; // Marca el check si está completado
                checkbox.dataset.id = entity.ID; // Almacena el ID en el checkbox
                li.appendChild(checkbox);

                //Crea un span para el texto
                const textSpan = document.createElement('span');
                textSpan.innerHTML = `
                    <strong>${entity.Titulo || 'Sin nombre'}:</strong> 
                    <span>${entity.Descripcion || 'Sin descripción'}</span>
                    <span>(Cantidad: ${entity.Cantidad || 0})</span>
                `;
                li.appendChild(textSpan);

                //Crea el botón de eliminar
                const deleteButton = document.createElement('button');
                deleteButton.textContent = 'Eliminar';
                deleteButton.className = 'delete-btn'; // Clase para CSS y JS
                deleteButton.dataset.id = entity.ID; // Almacena el ID en el botón
                li.appendChild(deleteButton);
                
                entityList.appendChild(li); //agrega el <li> completo al <ul>
            });

        } catch (error) {
            console.error('Error al obtener las entidades:', error);
            entityList.innerHTML = '<li>Error al cargar la lista.</li>';
        }
    }

    //Crear nuevo producto
    async function handleFormSubmit(event) {
         event.preventDefault(); 
        
        const nameInput = document.getElementById('titulo');
        const descriptionInput = document.getElementById('descripcion');
        const cantidadInput = document.getElementById('cantidad');

        //Crea el objeto de datos para enviar
        const newData = {
            Titulo: nameInput.value,
            Descripcion: descriptionInput.value,
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
                body: JSON.stringify(newData)
            });

            if (!response.ok) throw new Error(`Error HTTP: ${response.status}`);

            //Limpia los campos del formulario
            nameInput.value = '';
            descriptionInput.value = '';
            cantidadInput.value = '';

            //Refresca la lista para mostrar el nuevo ítem
            fetchAndRenderEntities();

        } catch (error) {
            console.error('Error al crear la entidad:', error);
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
                fetchAndRenderEntities(); 

            } catch (error) {
                console.error('Error al eliminar la entidad:', error);
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
                console.error('Error al actualizar la entidad:', error);
                //Si la API falla, revierte el checkbox a su estado anterior
                target.checked = !target.checked;
                alert("No se pudo actualizar el producto.");
            }
        }
    }

    //Configuración Inicial
    
    //Añade el "event listener" al formulario
    entityForm.addEventListener('submit', handleFormSubmit);

    //Añade un solo listener a toda la lista para manejar los clics
    entityList.addEventListener('click', handleListClick);

    //Carga la lista de entidades al iniciar la página
    fetchAndRenderEntities();
});

