package auth

import (
	"fmt"
	"picolor-backend/app/domain/auth"
	"picolor-backend/app/domain/room"
)


type CreateUserAndJoinRoom struct {
	RoomID   auth.RoomID
	UserName auth.UserName
}

func (c *AuthServiceImpl) CreateUserAndJoinRoom(createInfo CreateUserAndJoinRoom) (*auth.UserID, error) {
	joinUser := auth.JoinUser{
		RoomID:   createInfo.RoomID,
		UserName: createInfo.UserName,
	}

	createdUser, err := c.authRepo.CreateUser(joinUser)

	if err != nil {
		return nil, fmt.Errorf("failed to create user:%w", err)
	}

	joinRoom := room.RoomMember{
		RoomID: createInfo.RoomID,
		UserID: createdUser.ID,
	}

	roomMember, err := c.roomRepo.CreateRoomMember(joinRoom)

	if err != nil {
		return nil, fmt.Errorf("failed to create room member: %w", err)
	}

	return &roomMember.UserID, nil
}
