package auth

type AuthRepository interface {
	CreateUser(JoinUser JoinUser) (*User, error)
	DeleteUser(deleteUser DeleteUser) (*User, error)
}
