package v1

import (
	v1 "picolor-backend/app/presentation/v1"

	"github.com/gorilla/mux"
)

func NewRouter() *mux.Router {
	router := mux.NewRouter()

	router.HandleFunc("/hello", v1.HelloHandler).Methods("GET")

	return router
}
