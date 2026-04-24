package main

import (
	"log"
	"net/http"

	"github.com/agdaniel10/Go-BasicAPI/internal/config"
	"github.com/agdaniel10/Go-BasicAPI/internal/db"
	"github.com/agdaniel10/Go-BasicAPI/internal/handler"
	"github.com/agdaniel10/Go-BasicAPI/internal/repository"
	"github.com/agdaniel10/Go-BasicAPI/internal/service"
	"github.com/joho/godotenv"
)

func main() {

	if err := godotenv.Load(); err != nil {
		log.Println("no .env file found, reading from environment")
	}

	cfg := config.Load()

	client := db.Connect(cfg.MongoURI)
	database := client.Database(cfg.MongoDB)

	userRepo := repository.NewUserRepository(database)
	userService := service.NewUserService(userRepo)
	userHandler := handler.NewUserHandler(userService)

	mux := http.NewServeMux()
	mux.HandleFunc("GET /users", userHandler.GetAll)
	mux.HandleFunc("GET /users/{id}", userHandler.GetByID)
	mux.HandleFunc("POST /users", userHandler.Create)

	log.Printf("Server running on port %s", cfg.Port)
	log.Fatal(http.ListenAndServe(":"+cfg.Port, mux))
}
