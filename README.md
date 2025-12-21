# Prueba Técnica: CRUD Básico

Este proyecto implementa un API RESTful para gestionar libros y autores, cumpliendo con los requisitos del **Caso 1: CRUD Básico**.

## Resumen

- **API REST** sin autenticación.
- Entidades: **Libro** (principal) y **Autor** (relacionada).
- Relación **1:N**: un autor puede tener muchos libros; un libro pertenece a un solo autor.
- **Base de datos**: MySQL.
- **Eliminación en cascada**: al eliminar un autor, se eliminan sus libros.
- **Validaciones**: campos obligatorios, email único, manejo de errores HTTP adecuado.
- **Datos de prueba**: se generan automáticamente con `gofakeit` al iniciar la app (solo si la base está vacía).

---

## Tecnologías y Frameworks

| Capa | Tecnología |
|------|------------|
| **Lenguaje** | Go (Golang) |
| **Framework Web** | [Gin](https://gin-gonic.com/) |
| **ORM** | [GORM](https://gorm.io/) |
| **Generación de datos** | [gofakeit](https://github.com/brianvoe/gofakeit) |
| **Gestión de variables** | [godotenv](https://github.com/joho/godotenv) |
| **Base de datos** | MySQL |

---

## Estructura del Proyecto

```bash
.
├── .env.example                # Template para variables de entorno
├── .gitignore
├── go.mod
├── go.sum
├── main.go                     # Punto de entrada, configuración de DB y rutas
└── internal/
    ├── models/                 # Definición de entidades (Autor, Libro)
    ├── service/                # Lógica de negocio (CRUD, validaciones)
    └── transport/              # Manejo de HTTP (handlers para Gin)
```

### Descripción por capa

- **`models/`**:  
  Contiene los structs `Autor` y `Libro` con sus tags de GORM para mapeo a la base de datos, incluyendo relaciones y restricciones (clave foránea, cascadas, unicidad).

- **`service/`**:  
  Implementa la lógica del CRUD: crear, listar, obtener, actualizar y eliminar libros. Incluye validaciones básicas y carga relaciones con `Preload`. También contiene el seeder de datos de prueba.

- **`transport/`**:  
  Define los handlers para Gin. Mapea las rutas HTTP a los métodos del servicio, maneja la serialización/deserialización JSON y los códigos de estado HTTP.

- **`main.go`**:  
  Configura la conexión a MySQL (usando `.env`), ejecuta migraciones, siembra datos de prueba y registra las rutas.

---

## Instrucciones de Uso

### Requisitos
- Go 1.20+
- MySQL 5.7+ o 8.0
- Una base de datos llamada `prueba_tecnica_bd` (debe existir, pero puede estar vacía)

### Pasos

1. **Clonar el repositorio**:
   ```bash
   git clone <tu-repo>
   cd prueba-tecnica-back
   ```

2. **Configurar variables de entorno**:
   ```bash
   cp .env.example .env
   ```
   Edita `.env` con tus credenciales de MySQL.

3. **Instalar dependencias**:
   ```bash
   go mod tidy
   ```

4. **Ejecutar la API**:
   ```bash
   go run main.go
   ```
   La API estará disponible en: `http://localhost:8080`

5. **Probar con Postman**:
   - `GET /libros` → Listar todos los libros (con autores).
   - `POST /libros` → Crear un libro.
   - `GET /libros/1` → Obtener un libro por ID.
   - `PUT /libros/1` → Actualizar un libro.
   - `DELETE /libros/1` → Eliminar un libro.

---