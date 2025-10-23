CREATE TABLE producto (
    ID SERIAL PRIMARY KEY,
    titulo VARCHAR(255) NOT NULL,
    descripcion VARCHAR(255) NOT NULL
);

CREATE TABLE lista_productos (
    ID SERIAL PRIMARY KEY,
    ID_producto INT NOT NULL REFERENCES producto(ID),
    cantidad INT NOT NULL, 
    comprado BOOLEAN DEFAULT FALSE
);