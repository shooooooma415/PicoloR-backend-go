package v1

import (
	"encoding/json"
	"net/http"
	authApp "picolor-backend/app/application/auth"
	auth "picolor-backend/app/domain/auth"
	room "picolor-backend/app/domain/room"
	"picolor-backend/app/infrastructure/postgresql/utils"
)

type PostMemberParams struct {
	RoomID auth.RoomID `json:"roomID"`
	UserID auth.UserID `json:"userID"`
}

type PostMemberResponse struct {
	UserID auth.UserID `json:"userID"`
}

type PostMember struct {
	service *authApp.AuthServiceImpl
}

func NewPostMember(service *authApp.AuthServiceImpl) *PostMember {
	return &PostMember{service: service}
}

func (pc *PostMember) PostMemberHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req PostMemberParams
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			http.Error(w, "Invalid request payload", http.StatusBadRequest)
			return
		}

		joinRoom := room.RoomMember{
			RoomID: req.RoomID,
			UserID: req.UserID,
		}

		_, err := pc.service.RegisterMember(joinRoom)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode(utils.NewErrorResponse(err.Error()))
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
	}
}
