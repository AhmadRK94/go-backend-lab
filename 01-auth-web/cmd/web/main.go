package main

import (
	"log"

	"github.com/AhmadRK94/go-backend-lab/auth-web/internal/database"
	"github.com/AhmadRK94/go-backend-lab/auth-web/internal/user"
)

func main() {
	db, err := database.InitDB("test.db")
	if err != nil {
		log.Fatal(err)
	}
	newUserRepo := user.NewUserRepository(db)
	newUserService := user.NewUserService(newUserRepo)

}
