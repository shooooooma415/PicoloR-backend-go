package auth

type AuthRepository interface {
	CreateUser(JoinRoom JoinUser) (*User, error)
	DeleteUser(deleteUser DeleteUser) (*User, error)
}
