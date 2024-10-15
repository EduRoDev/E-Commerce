package database

import (
	"fmt"
	"log"
	"os"

	models "github.com/EduRoDev/E-commerce/Models"
	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var Database *gorm.DB

func Init() {
	// Cargar las variables de entorno
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file:", err)
		panic(err)
	}

	// Obtener las credenciales de la base de datos
	user := os.Getenv("DATABASE_USER")
	pass := os.Getenv("DATABASE_PASS")
	host := os.Getenv("DATABASE_HOST")
	port := os.Getenv("DATABASE_PORT")
	name := os.Getenv("DATABASE_NAME")

	// Conectarse a la base de datos MySQL
	dns := fmt.Sprintf("%s:%s@tcp(%s:%s)/?charset=utf8mb4&parseTime=True&loc=Local", user, pass, host, port)
	db, err := gorm.Open(mysql.Open(dns), &gorm.Config{})
	if err != nil {
		log.Fatal("Error connecting to MySQL server:", err)
		panic(err)
	}

	// Crear la base de datos si no existe
	err = db.Exec("CREATE DATABASE IF NOT EXISTS " + name).Error
	if err != nil {
		log.Fatal("Error creating database:", err)
		panic(err)
	}

	// Cerrar la conexión anterior
	sqlDB, _ := db.DB()
	sqlDB.Close()

	// Conectarse a la base de datos recién creada
	dns = fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", user, pass, host, port, name)
	Database, err = gorm.Open(mysql.Open(dns), &gorm.Config{})
	if err != nil {
		log.Fatal("Error connecting to the database:", err)
		panic(err)
	}

	// Ahora realiza las migraciones
	createSchema(Database)
}

func createSchema(db *gorm.DB) {
	// Migrar las tablas de la base de datos
	err := db.AutoMigrate(&models.Cliente{}, &models.Producto{}, &models.Orden{})
	// Comprobar si hay errores
	if err != nil {
		log.Fatal("Error creating schema:", err)
		panic(err)
	}

	// Mostrar un mensaje de confirmación
	log.Println("Schema created successfully")
}
