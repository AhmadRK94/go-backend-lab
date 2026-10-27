package main

import (
	"log"
	"net/http"
	"time"

	"github.com/AhmadRK94/go-backend-lab/auth-web/internal/auth"
	"github.com/AhmadRK94/go-backend-lab/auth-web/internal/database"
	"github.com/AhmadRK94/go-backend-lab/auth-web/internal/user"
)

func main() {
	db, err := database.InitDB("test.db")
	if err != nil {
		log.Fatal(err)
	}

	userRepo := user.NewUserRepository(db)
	userService := user.NewUserService(userRepo)
	userHandler := user.NewUserHandler(userService)

	jwtManager := auth.NewJWTManager(
		"test_access_secret",
		"test_refresh_secret",
		15*time.Minute,
		24*time.Hour,
	)
	authService := auth.NewAuthService(userRepo, jwtManager)

	http.HandleFunc("POST /register", userHandler.Register)
	
}
