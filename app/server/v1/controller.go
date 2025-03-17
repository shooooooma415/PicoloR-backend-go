package v1

import (
	authApp "picolor-backend/app/application/auth"
	colorApp "picolor-backend/app/application/color"
	v1 "picolor-backend/app/presentation/v1"

	"github.com/gorilla/mux"
)

func ControllerRouter(router *mux.Router, authService *authApp.AuthServiceImpl, colorService *colorApp.ColorServiceImpl) {
	postController := v1.NewUserController(authService)
	deleteController := v1.NewDeleteController(authService)
	getThemeColors := v1.NewGetThemeColor(colorService)

	router.HandleFunc("/user", postController.PostUserHandler()).Methods("POST")
	router.HandleFunc("/room", deleteController.DeleteControllerHandler()).Methods("DELETE")
	router.HandleFunc("/colors", getThemeColors.GetThemeColorHandler()).Methods("GET")
}
