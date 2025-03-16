package auth

import (
	"fmt"
	"picolor-backend/app/domain/auth"
)

func (c *AuthServiceImpl) DeleteMember(deleteInfo auth.DeleteUser) (*auth.UserID, error) {

	_, err := c.roomRepo.DeleteRoomMember(deleteInfo.UserID)

	if err != nil {
		return nil, fmt.Errorf("failed to delete user:%w", err)
	}

	deletedUser, err := c.authRepo.DeleteUser(deleteInfo)

	if err != nil {
		return nil, fmt.Errorf("failed to delete room member:%w", err)
	}
	return &deletedUser.ID, nil
}
