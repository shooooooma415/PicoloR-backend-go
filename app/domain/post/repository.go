package post

import "picolor-backend/app/domain/room"

type PostRepository interface {
	GetPosts(roomID room.RoomID) ([]Post, error)
}