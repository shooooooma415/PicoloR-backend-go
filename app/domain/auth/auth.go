package auth

import (
	"picolor-backend/app/domain/room"
)

type UserName string

type UserID int

type User struct {
	Id   UserID
	Name UserName
}

type JoinUser struct {
	RoomID room.RoomID
	UserName UserName
}

type DeleteUser struct {
	RoomID room.RoomID
	UserID UserID
}