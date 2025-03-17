package v1

import (
	authApp "picolor-backend/app/application/auth"
	colorApp "picolor-backend/app/application/color"
	v1 "picolor-backend/app/presentation/v1"

	"github.com/gorilla/mux"
)

func ControllerRouter(router *mux.Router, authService *authApp.AuthServiceImpl, colorService *colorApp.ColorServiceImpl) {
	postUser := v1.NewPostUser(authService)
	postMember := v1.NewPostMember(authService)
	deleteController := v1.NewDeleteController(authService)
	getThemeColors := v1.NewGetThemeColor(colorService)

	router.HandleFunc("/user", postUser.PostUserHandler()).Methods("POST")
	router.HandleFunc("/room", postMember.PostMemberHandler()).Methods("POST")
	router.HandleFunc("/user", deleteController.DeleteControllerHandler()).Methods("DELETE")
	router.HandleFunc("/colors", getThemeColors.GetThemeColorHandler()).Methods("GET")
}
