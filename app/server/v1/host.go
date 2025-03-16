package v1

import (
	roomApp "picolor-backend/app/application/room"
	v1 "picolor-backend/app/presentation/v1"

	"github.com/gorilla/mux"
)

func HostRouter(router *mux.Router, service *roomApp.RoomServiceImpl) {
	postRoom := v1.NewPostRoom(service)
	deleteRoom := v1.NewDeleteRoom(service)

	router.HandleFunc("/room", postRoom.PostRoomHandler()).Methods("POST")
	router.HandleFunc("/room", deleteRoom.DeleteRoomHandler()).Methods("DELETE")
}
