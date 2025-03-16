package main

import (
	"log"
	"net/http"
	v1 "picolor-backend/app/server/v1"
	authApp "picolor-backend/app/application/auth"
	repo "picolor-backend/app/infrastructure/postgresql/repository"
	utils "picolor-backend/app/infrastructure/postgresql/utils"
)

func main() {

	db, err := utils.ConnectDB()
	if err != nil {
		log.Fatalf("Failed to connect to the database: %v", err)
	}
	defer db.Close()

	authRepo := repo.NewAuthRepository(db)
	roomRepo := repo.NewRoomRepository(db)


	authService := authApp.NewAuthService(authRepo, roomRepo)

	router := v1.ControllerRouter(authService)

	log.Println("Server is running on :8000")
	http.ListenAndServe(":8000", router)
}