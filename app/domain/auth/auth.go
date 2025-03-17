package auth

type UserName string

type UserID int
type RoomID int
type ColorCode string
type ColorID int

type User struct {
	ID   UserID
	Name UserName
}

type DeleteUser struct {
	RoomID RoomID
	UserID UserID
}
