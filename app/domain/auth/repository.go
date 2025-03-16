package auth

type AuthRepository interface {
	CreateUser(JoinUser JoinUser) (*User, error)
	DeleteUserByUserID(UserID UserID) (*User, error)
}
