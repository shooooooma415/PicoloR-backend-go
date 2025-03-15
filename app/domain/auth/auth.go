package auth

import (
	"picolor-backend/app/domain/room"
)

type UserName string

type UserId int

type User struct {
	Id   UserId
	Name UserName
}

type JoinUser struct {
	RoomId room.RoomId
	UserName UserName
}

type DeleteUser struct {
	RoomId room.RoomId
	UserId UserId
}