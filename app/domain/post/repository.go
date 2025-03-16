package post

import "picolor-backend/app/domain/auth"

type PostRepository interface {
	GetPosts(roomID auth.RoomID) ([]Post, error)
}
