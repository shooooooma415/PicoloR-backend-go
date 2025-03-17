package room

import (
	"picolor-backend/app/domain/auth"
	"picolor-backend/app/domain/post"
	"picolor-backend/app/domain/room"
	"picolor-backend/app/domain/color"
)

type RoomServiceImpl struct {
	authRepo auth.AuthRepository
	roomRepo room.RoomRepository
	postRepo post.PostRepository
	colorRepo color.ColorRepository
}

func NewRoomService(authRepo auth.AuthRepository, roomRepo room.RoomRepository, postRepo post.PostRepository, colorRepo color.ColorRepository) *RoomServiceImpl {
	return &RoomServiceImpl{
		authRepo: authRepo,
		roomRepo: roomRepo,
		postRepo: postRepo,
		colorRepo: colorRepo,
	}
}
