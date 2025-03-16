package v1

import (
	"encoding/json"
	"net/http"
	roomApp "picolor-backend/app/application/room"
	"picolor-backend/app/domain/auth"
	"picolor-backend/app/infrastructure/postgresql/utils"
)

type PostRoomResponse struct {
	RoomID auth.RoomID `json:"roomID"`
}

type PostRoom struct {
	service *roomApp.RoomServiceImpl
}

func NewPostRoom(service *roomApp.RoomServiceImpl) *PostRoom {
	return &PostRoom{service: service}
}

func (pc *PostRoom) PostRoomHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		roomID, err := pc.service.CreateRoom()
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode(utils.NewErrorResponse(err.Error()))
			return
		}

		res := PostRoomResponse{
			RoomID: *roomID,
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(res)
	}
}
