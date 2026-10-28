package main

import (
	"log"
	"net/http"
	"time"

	"github.com/AhmadRK94/go-backend-lab/auth-web/internal/auth"
	"github.com/AhmadRK94/go-backend-lab/auth-web/internal/database"
	"github.com/AhmadRK94/go-backend-lab/auth-web/internal/user"
	"github.com/go-playground/validator/v10"
)

func main() {
	db, err := database.InitDB("test.db")
	if err != nil {
		log.Fatal(err)
	}
	validate := validator.New()
	userRepo := user.NewUserRepository(db)
	userService := user.NewService(userRepo)
	userHandler := user.NewHandler(userService, validate)

	jwtManager := auth.NewJWTManager(
		"test_access_secret",
		"test_refresh_secret",
		15*time.Minute,
		7*24*time.Hour,
	)
	authService := auth.NewService(userRepo, jwtManager)
	authHandler := auth.NewHandler(authService, validate)

	http.HandleFunc("POST /register", userHandler.Register)

}
