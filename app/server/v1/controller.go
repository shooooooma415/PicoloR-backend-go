package v1

import (
	authApp "picolor-backend/app/application/auth"
	v1 "picolor-backend/app/presentation/v1"

	"github.com/gorilla/mux"
)

func ControllerRouter(router *mux.Router, service *authApp.AuthServiceImpl) {
	postController := v1.NewPostController(service)
	router.HandleFunc("/room", postController.PostControllerHandler()).Methods("POST")
}
