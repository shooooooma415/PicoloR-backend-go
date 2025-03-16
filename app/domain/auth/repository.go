package auth

type AuthRepository interface {
	CreateUser(JoinUser JoinUser) (*User, error)
	DeleteUserByUserID(deleteUser DeleteUser) (*User, error)
}
