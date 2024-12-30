1. Inicia el servidor de PostgreSQL:

```bash
brew services start postgresql
```


2. Accede a PostgreSQL:
```bash
psql postgres
```


3. Aquí está el script completo para crear la base de datos y la tabla:

```sql
-- Crear la base de datos
CREATE DATABASE myapp;

-- Conectar a la base de datos
\c myapp

-- Crear la tabla de usuarios
CREATE TABLE users (
    id SERIAL PRIMARY KEY,
    name VARCHAR(100) NOT NULL,
    email VARCHAR(100) NOT NULL UNIQUE
);



4. Para acceder posteriormente:
```bash
# Acceder directamente a la base de datos myapp
psql myapp

# O si usas el usuario específico:
psql -U myapp_user -d myapp -h localhost
```


6. Para probar la conexión:
```sql
INSERT INTO users (name, email) VALUES 
    ('John Doe', 'john@example.com'),
    ('Jane Doe', 'jane@example.com');

-- Verificar los datos
SELECT * FROM users;
```