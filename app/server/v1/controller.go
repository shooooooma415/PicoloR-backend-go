package v1

import (
	authApp "picolor-backend/app/application/auth"
	v1 "picolor-backend/app/presentation/v1"

	"github.com/gorilla/mux"
)

func ControllerRouter(service *authApp.AuthServiceImpl) *mux.Router {
	router := mux.NewRouter()
	controllerRouter := router.PathPrefix("/controller").Subrouter()
	
	postController := v1.NewPostController(service)

	controllerRouter.HandleFunc("/room", postController.PostHandler()).Methods("POST")

	return router
}