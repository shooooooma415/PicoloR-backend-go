package auth

import (
	"picolor-backend/app/domain/auth"
)

func (c *AuthServiceImpl) DeleteUserByUserID(userID auth.UserID) (*auth.UserID, error) {
	deletedUser, _ := c.authRepo.DeleteUserByUserID(userID)
	return &deletedUser.ID, nil
}
