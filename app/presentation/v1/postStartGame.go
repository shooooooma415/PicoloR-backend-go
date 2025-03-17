package v1

import (
	"encoding/json"
	"net/http"
	roomApp "picolor-backend/app/application/room"
	"picolor-backend/app/domain/auth"
	"picolor-backend/app/infrastructure/postgresql/utils"
	"time"
)

type PostStartGameParam struct {
	RoomID  auth.RoomID `json:"roomID"`
	StartAt time.Time	      `json:"startAt"`
}

type PostStartGame struct {
	service *roomApp.RoomServiceImpl
}

func NewPostStartGame(service *roomApp.RoomServiceImpl) *PostRoom {
	return &PostRoom{service: service}
}

func (pc *PostStartGame) PostStartGameHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var param PostStartGameParam
		if err := json.NewDecoder(r.Body).Decode(&param); err != nil {
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(utils.NewErrorResponse(err.Error()))
			return
		}

		_, err := pc.service.StartGame(roomApp.StartGame{
			RoomID:  param.RoomID,
			StartAt: param.StartAt,
		})
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode(utils.NewErrorResponse(err.Error()))
			return
		}


		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
	}
}