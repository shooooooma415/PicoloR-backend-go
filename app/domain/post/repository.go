package post

import "picolor-backend/app/domain/room"

type PostRepository interface {
	GetPosts(roomId room.RoomId) ([]Post, error)
}