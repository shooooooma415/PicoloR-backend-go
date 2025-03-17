package v1

import (
	"encoding/json"
	"net/http"
	authApp "picolor-backend/app/application/auth"
	auth "picolor-backend/app/domain/auth"
	"picolor-backend/app/infrastructure/postgresql/utils"
)

type PostUserParams struct {
	UserName auth.UserName `json:"userName"`
}

type PostUserResponse struct {
	UserID auth.UserID `json:"userID"`
}

type PostUser struct {
	service *authApp.AuthServiceImpl
}

func NewUserController(service *authApp.AuthServiceImpl) *PostUser {
	return &PostUser{service: service}
}

func (pc *PostUser) PostUserHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req PostUserParams
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			http.Error(w, "Invalid request payload", http.StatusBadRequest)
			return
		}


		userID, err := pc.service.CreateUser(req.UserName)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode(utils.NewErrorResponse(err.Error()))
			return
		}

		res := PostUserResponse{
			UserID: *userID,
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(res)
	}
}
