package main

import (
	"fmt"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	"prueba-tecnica-back/internal/middleware"
	"prueba-tecnica-back/internal/models"
	"prueba-tecnica-back/internal/service"
	"prueba-tecnica-back/internal/transport"
)

func main() {
	// Cargar variables de entorno
	if err := godotenv.Load(); err != nil {
		log.Println("No se encontró .env, usando variables del sistema")
	}

	// Inicializar JWT
	jwtSecret := os.Getenv("JWT_SECRET")
	if jwtSecret == "" {
		log.Fatal("JWT_SECRET no configurado en .env")
	}
	service.InitJWT(jwtSecret)

	// Configuración de MySQL
	user := os.Getenv("MYSQL_USER")
	password := os.Getenv("MYSQL_PASSWORD")
	host := os.Getenv("MYSQL_HOST")
	port := os.Getenv("MYSQL_PORT")
	database := os.Getenv("MYSQL_DATABASE")

	if user == "" {
		user = "root"
	}
	if password == "" {
		password = ""
	}
	if host == "" {
		host = "localhost"
	}
	if port == "" {
		port = "3306"
	}
	if database == "" {
		database = "prueba_tecnica_ej2"
	}

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		user, password, host, port, database)

	// Conectar a la base de datos
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("No se pudo conectar a MySQL:", err)
	}

	// Ejecutar migraciones
	if err = db.AutoMigrate(&models.User{}, &models.Post{}); err != nil {
		log.Fatal("Error en migración:", err)
	}

	// Inicializar servicios
	authService := service.NewAuthService(db)
	postService := service.NewPostService(db)

	// Inicializar handlers
	authHandler := transport.NewAuthHandler(authService)
	postHandler := transport.NewPostHandler(postService)

	// Configurar router
	r := gin.Default()

	// Rutas públicas
	r.POST("/register", authHandler.Register)
	r.POST("/login", authHandler.Login)

	// Rutas protegidas
	posts := r.Group("/posts")
	posts.Use(middleware.AuthMiddleware())
	{
		posts.POST("", postHandler.Create)
		posts.GET("", postHandler.List)
		posts.GET("/:id", postHandler.GetByID)
		posts.PUT("/:id", postHandler.Update)
		posts.DELETE("/:id", postHandler.Delete)
	}

	log.Println("Servidor corriendo en http://localhost:8080")
	r.Run(":8080")
}
