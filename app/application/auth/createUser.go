package auth

import "picolor-backend/app/domain/auth"

func (c *AuthServiceImpl) CreateUser(userName auth.UserName) (*auth.UserID, error) {
	createdUser, err := c.authRepo.CreateUser(userName)
	if err != nil {
		return nil, err
	}
	return &createdUser.ID, nil
}
