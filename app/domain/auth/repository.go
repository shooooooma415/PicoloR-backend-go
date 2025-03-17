package auth

type AuthRepository interface {
	CreateUser(UserName UserName) (*User, error)
	DeleteUserByUserID(UserID UserID) (*User, error)
	FindUserByUserID(UserID UserID) (*User, error)
}
