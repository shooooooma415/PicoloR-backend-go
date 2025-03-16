package main

import (
	"log"
	"net/http"

	authApp "picolor-backend/app/application/auth"
	roomApp "picolor-backend/app/application/room"
	repo "picolor-backend/app/infrastructure/postgresql/repository"
	utils "picolor-backend/app/infrastructure/postgresql/utils"
	v1 "picolor-backend/app/server/v1"

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

	authService := authApp.NewAuthService(authRepo, roomRepo)
	roomService := roomApp.NewRoomService(authRepo, roomRepo)

	// ここで router を定義
	router := mux.NewRouter()

	// サブルーターに登録
	controllerRouter := router.PathPrefix("/controller").Subrouter()
	v1.ControllerRouter(controllerRouter, authService)

	hostRouter := router.PathPrefix("/host").Subrouter()
	v1.HostRouter(hostRouter, roomService)

	// NotFoundHandlerの追加
	router.NotFoundHandler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf("NotFound: %s %s", r.Method, r.URL.Path)
		http.Error(w, "Custom 404 - Not Found", http.StatusNotFound)
	})

	log.Println("Server is running on :8000")
	if err := http.ListenAndServe(":8000", router); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}