# Caso 2: CRUD con Autenticación

## Descripción

En este caso, deberás implementar un CRUD para una entidad con autenticación basada en JWT. Además, deberás incluir paginación, filtrado y ordenamiento en las consultas.

## Historias de Usuario

1. **HU-01** : Como usuario nuevo, quiero tener la posibilidad de registrarme en la aplicación.
2. **HU-02** : Como usuario registrado, quiero tener la posibilidad de autenticarme en la aplicación.
3. **HU-03** : Como usuario autenticado, quiero poder crear registros.
4. **HU-04** : Como usuario autenticado, quiero poder listar registros con paginación.
5. **HU-05** : Como usuario autenticado, quiero poder filtrar registros por usuario que lo creó.
5. **HU-06** : Como usuario autenticado, quiero poder filtrar por registros creados por mí.
6. **HU-07** : Como usuario autenticado, quiero poder ordenar registros por fecha de creación.
7. **HU-08** : Como usuario autenticado, quiero poder actualizar y eliminar registros.

## Especificaciones

- Entidad sugerida: Post.
- Implementa autenticación con JWT.
- Usa paginación para listar registros.
- Añade endpoints para filtrar y ordenar.
- Construir una base de datos en MySQL y realizar migraciones con datos de ejemplo para la revisión.
- Realizar las validaciones respectivas en las entidades de Base de Datos.

## Detalles de las entidades

**Usuario**

- id: Autoincremental, Primary Key, Integer.
- name: string, not null.
- email: string, not null, email válida y único.
- password: string, not null, mínimo 8 carácteres, mínimo 1 mayúscula, 1 número y un carácter especial.

**Post**

- id: Autoincremental, Primary Key, Integer.
- title: string, not null.
- content: string, not null.
- user_id: Foreign Key de id (de la entidad Usuario), Implementar eliminación en cascada.