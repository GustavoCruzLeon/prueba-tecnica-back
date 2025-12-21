package main

import (
	"fmt"
	"log"
	"os"

	"github.com/brianvoe/gofakeit/v6"
	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	"prueba-tecnica-back/internal/models"
	"prueba-tecnica-back/internal/service"
	"prueba-tecnica-back/internal/transport"

	"github.com/gin-gonic/gin"
)

func seedDatabase(db *gorm.DB) {
	var autorCount int64
	db.Model(&models.Autor{}).Count(&autorCount)
	if autorCount > 0 {
		return
	}

	gofakeit.Seed(1)

	autores := make([]models.Autor, 5)
	for i := range autores {
		autores[i] = models.Autor{
			Name:  gofakeit.Name(),
			Email: gofakeit.Email(),
		}
		if err := db.Create(&autores[i]).Error; err != nil {
			log.Printf("Error al crear autor: %v", err)
		}
	}

	libros := make([]models.Libro, 10)
	for i := range libros {
		libros[i] = models.Libro{
			Title:       gofakeit.BookTitle(),
			Description: gofakeit.Sentence(8), // 8 palabras
			Price:       float64(gofakeit.IntRange(10, 100)) + float64(gofakeit.Float32Range(0, 0.99)),
			AuthorID:    uint(gofakeit.IntRange(1, 5)), // Coincide con los 5 autores
		}
		if err := db.Create(&libros[i]).Error; err != nil {
			log.Printf("Error al crear libro: %v", err)
		}
	}

	log.Println("Base de datos sembrada con datos realistas")
}

func main() {

	if err := godotenv.Load(); err != nil {
		log.Println("No se encontró .env, usando variables del sistema")
	}

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
		database = "prueba_tecnica_bd"
	}

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		user, password, host, port, database)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("No se pudo conectar a MySQL:", err)
	}

	// Migrar modelos
	if err = db.AutoMigrate(&models.Autor{}, &models.Libro{}); err != nil {
		log.Fatal("Error en migración:", err)
	}

	// Sembrar datos de prueba
	seedDatabase(db)

	// Iniciar servicio y rutas
	svc := service.New(db)
	handler := transport.New(svc)

	r := gin.Default()

	libros := r.Group("/libros")
	{
		libros.GET("", handler.GetAll)
		libros.POST("", handler.Create)
		libros.GET("/:id", handler.GetByID)
		libros.PUT("/:id", handler.Update)
		libros.DELETE("/:id", handler.Delete)
	}

	log.Println("Servidor corriendo en http://localhost:8080")
	r.Run(":8080")
}
