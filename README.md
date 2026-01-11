# Caso 2: CRUD con Autenticación

API RESTful en Go que implementa autenticación JWT, paginación, filtrado y ordenamiento para la gestión de posts.

## Resumen

- **Autenticación segura** con JWT y validación de contraseñas.
- **CRUD de posts** con control de acceso: solo los dueños pueden editar/eliminar sus posts.
- **Listado avanzado**: paginación, filtrado por usuario y ordenamiento por fecha.
- **Arquitectura limpia**: separación en capas con DTOs para no exponer la base de datos.
- **Validaciones**: contraseña fuerte, tokens expirados, accesos no autorizados.

---

## Tecnologías

| Capa | Tecnología |
|------|-----------|
| Lenguaje | Go 1.22+ |
| Framework Web | [Gin](https://gin-gonic.com/) |
| ORM | [GORM](https://gorm.io/) |
| Autenticación | [golang-jwt/jwt](https://github.com/golang-jwt/jwt) |
| Gestión de entorno | [godotenv](https://github.com/joho/godotenv) |
| Base de datos | MySQL |

---

## Estructura del Proyecto

```
internal/
├── dto/             # DTOs para requests/responses (sin exposición de BD)
├── models/          # Modelos de base de datos (User, Post)
├── service/         # Lógica de negocio (auth, posts, JWT)
├── transport/       # Handlers HTTP (Gin)
└── middleware/      # Middleware de autenticación JWT
```

---

## Instrucciones de Uso

### Requisitos
- Go 1.22+
- MySQL
- Base de datos creada (ej. `prueba_caso2`)

### Configuración
1. Clona el proyecto:
   ```bash
   git clone <tu-repo>
   cd prueba-tecnica-back
   ```

2. Crea `.env`:
   
   Edita con tus credenciales de MySQL y una clave JWT segura.

3. Instala dependencias:
   ```bash
   go mod tidy
   ```

4. Ejecuta:
   ```bash
   go run main.go
   ```
   La API estará en `http://localhost:8080`.

---

## Endpoints

### Autenticación (públicos)
| Método | Ruta | Descripción |
|--------|------|------------|
| `POST` | `/register` | Registrar usuario |
| `POST` | `/login`    | Iniciar sesión |

### Posts (protegidos con JWT)
| Método | Ruta | Descripción |
|--------|------|------------|
| `POST`   | `/posts`          | Crear post |
| `GET`    | `/posts`          | Listar posts (mis posts por defecto) |
| `GET`    | `/posts/:id`      | Ver post específico |
| `PUT`    | `/posts/:id`      | Actualizar post (solo dueño) |
| `DELETE` | `/posts/:id`      | Eliminar post (solo dueño) |

### Parámetros de consulta (en `GET /posts`)
- `page=1` → número de página (por defecto: 1)
- `limit=10` → posts por página (máx: 100)
- `user_id=2` → filtrar por usuario (opcional)
- `sort_order=asc|desc` → ordenar por fecha (por defecto: `desc`)

> Todas las rutas de posts requieren header:  
> `Authorization: Bearer <tu_token>`

---

## Ejemplo de flujo en Postman

1. **Registro**:
   ```json
   POST /register
   { "name": "José", "email": "jose@example.com", "password": "Passw0rd!" }
   ```

2. **Login**:
   ```json
   POST /login
   { "email": "jose@example.com", "password": "Passw0rd!" }
   ```
   → Guarda el `token`.

3. **Crear post**:
   ```json
   POST /posts
   Authorization: Bearer <token>
   { "title": "Hola", "content": "Primer post" }
   ```

4. **Listar mis posts**:
   ```http
   GET /posts
   Authorization: Bearer <token>
   ```

---

## Seguridad

- Contraseñas hasheadas con `bcrypt`.
- Tokens JWT firmados con clave secreta.
- Validación de permisos en cada operación sensible.
- DTOs evitan exposición de campos sensibles (como `password`).

---

## Notas

- Al iniciar, la app crea las tablas `users` y `posts` si no existen.
- El campo `user_id` en posts es obligatorio y se relaciona con el usuario autenticado.
- Si no se envía `user_id` en `GET /posts`, se filtra automáticamente por el usuario del token.

---