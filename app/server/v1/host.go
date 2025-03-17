package v1

import (
	roomApp "picolor-backend/app/application/room"
	postApp "picolor-backend/app/application/post"
	v1 "picolor-backend/app/presentation/v1"

	"github.com/gorilla/mux"
)

func HostRouter(router *mux.Router, roomService *roomApp.RoomServiceImpl, postService *postApp.PostServiceImpl) {
	postRoom := v1.NewPostRoom(roomService)
	deleteRoom := v1.NewDeleteRoom(roomService)
	deleteRoomInfo := v1.NewDeleteRoomInfo(roomService)
	getResult := v1.NewGetResult(postService)
	postStartGame := v1.NewPostStartGame(roomService)

	router.HandleFunc("/room", postRoom.PostRoomHandler()).Methods("POST")
	router.HandleFunc("/room", deleteRoom.DeleteRoomHandler()).Methods("DELETE")
	router.HandleFunc("/room/reset", deleteRoomInfo.DeleteRoomInfoHandler()).Methods("DELETE")
	router.HandleFunc("/result", getResult.GetResultHandler()).Methods("GET")
	router.HandleFunc("/room/start", postStartGame.PostStartGameHandler()).Methods("POST")
}
