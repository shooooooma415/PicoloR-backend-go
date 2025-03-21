package main

import (
	"log"
	"net/http"

	authApp "picolor-backend/app/application/auth"
	colorApp "picolor-backend/app/application/color"
	postApp "picolor-backend/app/application/post"
	roomApp "picolor-backend/app/application/room"
	repo "picolor-backend/app/infrastructure/postgresql/repository"
	utils "picolor-backend/app/infrastructure/postgresql/utils"
	v1 "picolor-backend/app/server/v1"
	presentation "picolor-backend/app/presentation"

	"github.com/gorilla/mux"
)

func main() {
	db, err := utils.ConnectDB()
	if err != nil {
		log.Fatalf("Failed to connect to the database: %v", err)
	}
	defer db.Close()

	authRepo := repo.NewAuthRepository(db)
	roomRepo := repo.NewRoomRepository(db)
	postRepo := repo.NewPostRepository(db)
	colorRepo := repo.NewColorRepository(db)

	authService := authApp.NewAuthService(authRepo, roomRepo)
	roomService := roomApp.NewRoomService(authRepo, roomRepo, postRepo, colorRepo)
	colorService := colorApp.NewColorService(colorRepo)
	postService := postApp.NewPostService(postRepo, authRepo, colorRepo)

	router := mux.NewRouter()

	controllerRouter := router.PathPrefix("/controller").Subrouter()
	v1.ControllerRouter(controllerRouter, authService, colorService)

	hostRouter := router.PathPrefix("/host").Subrouter()
	v1.HostRouter(hostRouter, roomService, postService)

	corsHandler := presentation.CORSMiddleware(router)

	log.Println("Server is running on :8000")
	if err := http.ListenAndServe(":8000", corsHandler); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
