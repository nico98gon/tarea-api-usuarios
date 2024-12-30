package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"tarea-api-usuarios/internal/config"
	"tarea-api-usuarios/internal/domain/user"
	handler "tarea-api-usuarios/internal/infrastructure/http"
	"tarea-api-usuarios/internal/infrastructure/repository"

	"github.com/joho/godotenv"
)

func main() {
	// Cargar las variables de entorno desde el archivo .env
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}
	
	portAsString := os.Getenv("DB_PORT")
	if portAsString == "" {
		portAsString = "5432"
	}
	port, _ := strconv.Atoi(portAsString)

	newDatabase := config.DatabaseConfig{
		Host:     os.Getenv("DB_HOST"),
		Port:     port,
		User:     os.Getenv("DB_USER"),
		Password: os.Getenv("DB_PASSWORD"),
		DBName:   os.Getenv("DB_NAME"),
	}
	// Inicializar conexión a la base de datos
	db, err := config.NewPostgresConnection(newDatabase)
	if err != nil {
		log.Fatalf("Error connecting to database: %v", err)
	}
	defer db.Close()

	// Inicializar dependencias - repository
	userPostgresRepo := repository.NewPostgresUserRepository(db)

	fmt.Println(userPostgresRepo)
	userService := user.NewService(userPostgresRepo)
	userHandler := handler.NewUserHandler(userService)

	// Configurar rutas
	http.HandleFunc("GET /users", userHandler.GetUsers)
	http.HandleFunc("POST /users", userHandler.CreateUser)
	http.HandleFunc("GET /users/{id}", userHandler.GetUserByID)
	http.HandleFunc("PUT /users/{id}", userHandler.UpdateUser)
	http.HandleFunc("DELETE /users/{id}", userHandler.DeleteUser)

	// Iniciar servidor
	log.Println("Server running on http://localhost:8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
