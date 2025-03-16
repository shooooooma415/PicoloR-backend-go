package v1

import (
	"encoding/json"
	"net/http"
	roomApp "picolor-backend/app/application/room"
	auth "picolor-backend/app/domain/auth"
	"picolor-backend/app/infrastructure/postgresql/utils"
)

type DeleteRoomInfoResponse struct {
	RoomID auth.RoomID `json:"roomID"`
}

type DeleteRoomInfoParams struct {
	RoomID auth.RoomID `json:"roomID"`
}

type DeleteRoomInfo struct {
	service *roomApp.RoomServiceImpl
}

func NewDeleteRoomInfo(service *roomApp.RoomServiceImpl) *DeleteRoomInfo {
	return &DeleteRoomInfo{service: service}
}

func (pc *DeleteRoomInfo) DeleteRoomInfoHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req DeleteRoomParams
		roomID, err := pc.service.ResetRoom(req.RoomID)
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
