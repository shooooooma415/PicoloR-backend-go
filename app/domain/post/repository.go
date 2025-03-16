package post

import "picolor-backend/app/domain/auth"

type PostRepository interface {
	GetPosts(roomID auth.RoomID) ([]Post, error)
	DeletePost(userID auth.UserID)(*auth.UserID,error)
}
