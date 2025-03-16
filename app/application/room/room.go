package room

import (
	"picolor-backend/app/domain/auth"
	"picolor-backend/app/domain/post"
	"picolor-backend/app/domain/room"
)

type RoomServiceImpl struct {
	authRepo auth.AuthRepository
	roomRepo room.RoomRepository
	postRepo post.PostRepository
}

func NewRoomService(authRepo auth.AuthRepository, roomRepo room.RoomRepository, postRepo post.PostRepository) *RoomServiceImpl {
	return &RoomServiceImpl{
		authRepo: authRepo,
		roomRepo: roomRepo,
		postRepo: postRepo,
	}
}
