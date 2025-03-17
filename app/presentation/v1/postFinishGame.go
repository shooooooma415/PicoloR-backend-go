package v1

import (
	"encoding/json"
	"net/http"
	roomApp "picolor-backend/app/application/room"
	auth "picolor-backend/app/domain/auth"
	"picolor-backend/app/infrastructure/postgresql/utils"
)

type PostFinishGameParam struct {
	RoomID auth.RoomID `json:"roomID"`
}

type PostFinishGame struct {
	service *roomApp.RoomServiceImpl
}

func NewPostFinishGame(service *roomApp.RoomServiceImpl) *PostFinishGame {
	return &PostFinishGame{service: service}
}

func (pc *PostFinishGame) PostFinishGameHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var param PostFinishGameParam
		if err := json.NewDecoder(r.Body).Decode(&param); err != nil {
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(utils.NewErrorResponse(err.Error()))
			return
		}

		_, err := pc.service.FinishGame(param.RoomID)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode(utils.NewErrorResponse(err.Error()))
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
	}
}
