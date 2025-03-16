package v1

import (
	"encoding/json"
	"net/http"
	postApp "picolor-backend/app/application/post"
	auth "picolor-backend/app/domain/auth"
)

type DeletePostResponse struct {
	RoomID auth.RoomID `json:"roomID"`
}

type DeletePostParams struct {
	RoomID auth.RoomID `json:"roomID"`
}

type DeletePost struct {
	service *postApp.PostServiceImpl
}

func NewDeletePostResponse(roomID auth.RoomID) *DeletePostResponse {
	return &DeletePostResponse{
		RoomID: roomID,
	}
}

func (r *DeletePost) DeletePosts(roomID auth.RoomID) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		roomID, err := r.service.DeletePostsByRoomID(roomID)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		res := NewDeletePostResponse(*roomID)

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(res)
	}
}
