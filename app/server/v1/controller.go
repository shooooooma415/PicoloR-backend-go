package v1

import (
	authApp "picolor-backend/app/application/auth"
	v1 "picolor-backend/app/presentation/v1"

	"github.com/gorilla/mux"
)

func ControllerRouter(router *mux.Router, service *authApp.AuthServiceImpl) {
	postController := v1.NewPostController(service)
	deleteController := v1.NewDeleteController(service)

	router.HandleFunc("/room", postController.PostControllerHandler()).Methods("POST")
	router.HandleFunc("/room", deleteController.DeleteControllerHandler()).Methods("DELETE")
}
