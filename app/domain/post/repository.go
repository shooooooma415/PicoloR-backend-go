package post

import "picolor-backend/app/domain/auth"

type PostRepository interface {
	GetPosts(roomID auth.RoomID) ([]Post, error)
	DeletePostByUserID(userID auth.UserID)(*auth.UserID,error)
	DeletePostByRoomID(roomID auth.RoomID)(*auth.RoomID,error)
}
