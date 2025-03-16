package auth

import (
	"picolor-backend/app/domain/auth"
)

func (c *AuthServiceImpl) DeleteUserByUserID(deleteInfo auth.DeleteUser) (*auth.UserID, error) {
	deletedUser, _ := c.authRepo.DeleteUserByUserID(deleteInfo)
	return &deletedUser.ID, nil
}
