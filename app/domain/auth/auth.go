package auth

type UserName string

type UserID int
type RoomID int

type User struct {
	ID   UserID
	Name UserName
}

type JoinUser struct {
	RoomID   RoomID
	UserName UserName
}

type DeleteUser struct {
	RoomID RoomID
	UserID UserID
}
