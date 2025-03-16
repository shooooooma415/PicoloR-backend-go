package v1

import (
	"encoding/json"
	"net/http"
	postApp "picolor-backend/app/application/post"
	auth "picolor-backend/app/domain/auth"
	"picolor-backend/app/domain/post"
	"picolor-backend/app/infrastructure/postgresql/utils"
	"strconv"
)

type GetResultResponse struct {
	Results []post.GetPost `json:"results"`
}

type GetResult struct {
	service *postApp.PostServiceImpl
}

func NewGetResult(service *postApp.PostServiceImpl) *GetResult {
	return &GetResult{service: service}
}

func (pc *GetResult) GetResultHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		roomIDStr := r.URL.Query().Get("roomID")
		roomID, err := strconv.Atoi(roomIDStr)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		posts, err := pc.service.GetPostsByRoomID(auth.RoomID(roomID))
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode(utils.NewErrorResponse(err.Error()))
			return
		}

		res := GetResultResponse{
			Results: posts,
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(res)
	}
}
