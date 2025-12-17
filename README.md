# Caso 3: CRUD con Arquitectura Hexagonal

## Descripción

En este caso, deberás implementar un sistema que gestione tres entidades relacionadas (`Order`, `Customer`, `Product`) utilizando la Arquitectura Hexagonal y aplicando principios de DDD (Domain-Driven Design) . El objetivo es evaluar tu capacidad para estructurar el código en contextos delimitados (Bounded Contexts ) y separar las capas de dominio, aplicación e infraestructura.

Este caso está diseñado para evaluar habilidades avanzadas en Express, incluyendo la implementación de autenticación, operaciones CRUD, filtros, ordenamiento, entre otras, todo organizado bajo los principios de DDD.

## Historias de Usuario

1. **HU-01** : Como usuario, quiero poder registrarme en el sistema proporcionando mi nombre, correo electrónico y contraseña para acceder a las funcionalidades protegidas.
2. **HU-02** : Como usuario registrado, quiero poder iniciar sesión en el sistema utilizando mi correo electrónico y contraseña para acceder a mis datos y realizar operaciones.
3. **HU-03** : Como usuario autenticado, quiero poder crear nuevos clientes para gestionar sus órdenes posteriormente.
4. **HU-04** : Como usuario autenticado, quiero poder listar todos los clientes registrados en el sistema, con opciones de paginación, filtrado por nombre y ordenamiento por fecha de creación.
5. **HU-05** : Como usuario autenticado, quiero poder actualizar la información de un cliente existente para mantener los datos actualizados.
6. **HU-06** : Como usuario autenticado, quiero poder eliminar un cliente del sistema si ya no es necesario.
7. **HU-07** : Como usuario autenticado, quiero poder crear nuevos productos para asociarlos a las órdenes de los clientes.
8. **HU-08** : Como usuario autenticado, quiero poder listar todos los productos disponibles en el sistema, con opciones de filtrado por categoría y ordenamiento por precio.
9. **HU-09** : Como usuario autenticado, quiero poder crear órdenes asociando productos específicos a un cliente determinado.
10. **HU-10** : Como usuario autenticado, quiero poder ver el historial de órdenes de un cliente específico, mostrando los detalles de cada orden (productos asociados, cantidades y precios).
11. **HU-11** : Como usuario autenticado, quiero poder calcular el total gastado por un cliente en todas sus órdenes, considerando el precio unitario de los productos y las cantidades asociadas.
12. **HU-12** : Como usuario autenticado, quiero que cada producto tenga un código QR asociado que al escanearlo me muestre los datos del producto.

## Especificaciones

### Autenticación

- Implementa registro e inicio de sesión usando **JWT**.
- Asegúrate de que todas las rutas estén protegidas mediante middleware de autenticación.

### Código QR

- Implementa endpoints para generar código QR.
- Asocia el código QR a información del producto.

### Operaciones CRUD

- Implementa operaciones CRUD para las entidades `Customer`, `Product` y `Order`.
- Usa la **Arquitectura Hexagonal** para separar las capas de dominio, aplicación e infraestructura.

### Filtros y Ordenamiento

- En el listado de clientes, permite:

    - Filtrar por nombre.
    - Ordenar por fecha de creación (Ascendente/Descendente).

- En el listado de productos, permite:

    - Filtrar por categoría.
    - Ordenar por precio (Ascendente/Descendente).

### Relaciones

- Implementa una función que calcule el total gastado por un cliente en todas sus órdenes. Esta función debe:

    - Recorrer las órdenes del cliente.
    - Multiplicar el precio unitario de cada producto por su cantidad en la orden.
    - Sumar los totales de todas las órdenes.

## Bounded Contexts

### 1. Contexto de Identidad y Acceso (Identity & Access Management - IAM)

- **Responsabilidad** : Gestionar la autenicación y autorización de usuarios.
- **Entidades** :
    - `User`: Representa a los usuarios del sistema.
- **Casos de Uso** :
    - Registro de usuarios.
    - Inicio de sesión.
- **Directorio** :
```
src/IdentityAndAccess/User/
├── Domain/
│   ├── Entities/User.ts
│   ├── Contract/UserContract.ts
│   └── ValueObjects/*
├── Application/
│   ├── RegisterUseCase.ts
│   └── LoginUseCase.ts
└── Infrastructure/
    ├── Controllers/AuthController.ts
    ├── Validators/*
    ├── Routes/Router.ts
    └── Repositories/SequelizeUserRepository.ts
```

### 2. Contexto de Gestión de Clientes (Customer Management)

- **Responsabilidad** : Gestionar la creación, actualización, eliminación y consulta de clientes.
- **Entidades** :
    - `Customer` : Representa a los clientes.
- **Casos de Uso** :
    - Crear cliente.
    - Listar clientes (con filtros y ordenamiento).
    - Actualizar cliente.
    - Eliminar cliente.
- **Directorio** :
```
src/CustomerManagement/Customer/
├── Domain/
│   ├── Entities/Customer.ts
│   ├── Contract/CustomerContract.ts
│   └── ValueObjects/*
├── Application/
│   ├── CreateCustomerUseCase.ts
│   ├── ListCustomersUseCase.ts
│   ├── UpdateCustomerUseCase.ts
│   └── DeleteCustomerUseCase.ts
└── Infrastructure/
    ├── Controllers/CustomerController.ts
    ├── Validators/*
    ├── Routes/Router.ts
    └── Repositories/SequelizeCustomerRepository.ts
```

### 3. Contexto de Gestión de Productos (Product Management)

- **Responsabilidad** : Gestionar la creación, actualización, eliminación y consulta de productos.
- **Entidades** :
    - `Product`: Representa a los productos del sistema.
- **Casos de Uso** :
    - Crear producto.
    - Listar productos (con filtros y ordenamiento).
    - Actualizar producto.
    - Eliminar producto.
- **Directorio** :
```
src/ProductManagement/Product/
├── Domain/
│   ├── Entities/Product.ts
│   ├── Contract/ProductContract.ts
│   └── ValueObjects/*
├── Application/
│   ├── CreateProductUseCase.ts
│   ├── ListProductsUseCase.ts
│   ├── UpdateProductUseCase.ts
│   ├── DeleteProductUseCase.ts
│   └── GenerateQrUseCase.ts
└── Infrastructure/
    ├── Controllers/ProductController.ts
    ├── Controllers/QrController.ts
    ├── Validators/*
    ├── Routes/api.ts
    └── Repositories/EloquentProductRepository.ts
```

### 4. Contexto de Gestión de Órdenes (Order Management)

- **Responsabilidad** : Gestionar la creación, consulta y cálculo de órdenes.
- **Entidades** :
    - `Order`: Representa una orden.
    - `OrderItem`: Representa un producto asociado a una orden con su cantidad.
- **Casos de Uso** :
    - Crear orden.
    - Consultar historial de órdenes de un cliente.
    - Calcular el total gastado por un cliente.
- **Directorio** :
```
src/OrderManagement/Order/
├── Domain/
│   ├── Entities/Order.ts
│   ├── Entities/OrderItem.ts
│   ├── Contract/OrderContract.ts
│   └── ValueObjects/*
├── Application/
│   ├── CreateOrder.ts
│   ├── ListOrdersByCustomer.ts
│   └── CalculateTotalSpentByCustomer.ts
└── Infrastructure/
    ├── Controllers/OrderController.ts
    ├── Validators/*
    ├── Routes/Router.ts
    └── Repositories/SequelizeOrderRepository.ts
```

## Bases de datos

**Usuario**

- id: Autoincremental, Primary Key, Integer.
- name: string, not null.
- email: string, not null, email válida y único.
- password: string, not null, mínimo 8 carácteres, mínimo 1 mayúscula, 1 número y un carácter especial.

**Cliente**

- id: Autoincremental, Primary Key, Integer.
- name: string, not null.
- email: string, not null, email válida y único.

**Producto**

- id: Autoincremental, Primary Key, Integer.
- name: string, not null.
- category: string, not null, las categorías son: Electronics, Clothing, Books.
- price: float, not null.

**Orden**

- id: Autoincremental, Primary Key, Integer.
- status: string, not null, valor por defecto "Pending", los status son: pending, processing, completed, declined.
- total: float, nullable.
- shipping_address: text, nullable.
- shipped_at: timestamp, nullable.
- customer_id: Foreign Key de id (de la entidad Cliente), Implementar eliminación en cascada.

**Orden-Producto (Tabla Intermedia)**

- id: Autoincremental, Primary Key, Integer.
- order_id: Foreign Key de id (de la entidad Orden), Implementar eliminación en cascada.
- product_id: Foreign Key de id (de la entidad Producto), Implementar eliminación en cascada.
- quantity: integer, not null, valor por defecto 1

## Consejos

1. **Organización del Código** : Asegúrate de que cada contexto tenga su propia carpeta y siga la estructura de Dominio , Aplicación e Infraestructura.
2. **Interfaces** : Define interfaces en el nivel de Dominio y proporciona implementaciones concretas en Infraestructura.

## Tutoriales

Puedes guiarte buscando información acerca de arquitectura hexagonal con Typescript y node js en los siguientes videos que se adjuntan donde se explica teoría y práctica acerca de esto.

1. Arquitectura Hexagonal - Dominio: https://youtu.be/H0Sbna9rxog
2. Arquitectura Hexagonal - Aplicación: https://youtu.be/EcmRo9LUfyE
3. Arquitectura Hexagonal - Infraestructura: https://youtu.be/np7nOWWZHtU
4. Arquitectura Hexagonal - Testeando el API: https://youtu.be/7PKmEXHbGWI