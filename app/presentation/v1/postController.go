package v1

import (
	"encoding/json"
	"net/http"
	authApp "picolor-backend/app/application/auth"
	auth "picolor-backend/app/domain/auth"
	"picolor-backend/app/infrastructure/postgresql/utils"
)

type PostControllerParams struct {
	RoomID   auth.RoomID   `json:"roomID"`
	UserName auth.UserName `json:"userName"`
}

type PostControllerResponse struct {
	UserID auth.UserID `json:"userID"`
}

type PostController struct {
	service *authApp.AuthServiceImpl
}

func NewPostController(service *authApp.AuthServiceImpl) *PostController {
	return &PostController{service: service}
}

func (pc *PostController) PostControllerHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req PostControllerParams
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			http.Error(w, "Invalid request payload", http.StatusBadRequest)
			return
		}

		createInfo := authApp.CreateUserAndJoinRoom{
			RoomID:   req.RoomID,
			UserName: req.UserName,
		}

		userID, err := pc.service.CreateUserAndJoinRoom(createInfo)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode(utils.NewErrorResponse(err.Error()))
			return
		}

		res := PostControllerResponse{
			UserID: *userID,
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(res)
	}
}
