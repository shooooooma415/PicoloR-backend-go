package v1

import (
	"encoding/json"
	"net/http"
	roomApp "picolor-backend/app/application/room"
	"picolor-backend/app/domain/auth"
	"picolor-backend/app/infrastructure/postgresql/utils"
)

type DeleteRoomResponse struct {
	RoomID auth.RoomID `json:"roomID"`
}

type DeleteRoomParams struct {
	RoomID auth.RoomID `json:"roomID"`
}

type DeleteRoom struct {
	service *roomApp.RoomServiceImpl
}

func NewDeleteRoom(service *roomApp.RoomServiceImpl) *DeleteRoom {
	return &DeleteRoom{service: service}
}

func (pc *PostRoom) DeleteRoomHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req DeleteRoomParams
		roomID, err := pc.service.DeleteRoom(req.RoomID)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode(utils.NewErrorResponse(err.Error()))
			return
		}

		res := DeleteRoomResponse{
			RoomID: *roomID,
		}
		

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(res)
	}
}
