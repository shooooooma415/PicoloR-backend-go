package v1

import (
	"encoding/json"
	"net/http"
	authApp "picolor-backend/app/application/auth"
	auth "picolor-backend/app/domain/auth"
	"picolor-backend/app/infrastructure/postgresql/utils"
)

type DeleteControllerParams struct {
	UserID auth.UserID `json:"userID"`
}

type DeleteControllerResponse struct {
	UserID auth.UserID `json:"userID"`
}

type DeleteController struct {
	service *authApp.AuthServiceImpl
}

func NewDeleteController(service *authApp.AuthServiceImpl) *DeleteController {
	return &DeleteController{service: service}
}

func (pc *DeleteController) DeleteControllerHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req DeleteControllerParams
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			http.Error(w, "Invalid request payload", http.StatusBadRequest)
			return
		}


		userID, err := pc.service.DeleteUserByUserID(req.UserID)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode(utils.NewErrorResponse(err.Error()))
			return
		}

		res := DeleteControllerResponse{
			UserID: *userID,
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(res)
	}
}
