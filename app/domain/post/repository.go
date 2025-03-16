package post

import "picolor-backend/app/domain/auth"

type PostRepository interface {
	GetPosts(roomID auth.RoomID) ([]Post, error)
	DeletePostByRoomID(roomID auth.RoomID)(*auth.RoomID,error)
}