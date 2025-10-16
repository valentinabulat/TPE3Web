CREATE TABLE productos (
    ID SERIAL PRIMARY KEY,
    titulo VARCHAR(255) NOT NULL,
    descripcion VARCHAR(255) NOT NULL,
    cantidad INT NOT NULL, 
    comprado BOOLEAN DEFAULT FALSE
);