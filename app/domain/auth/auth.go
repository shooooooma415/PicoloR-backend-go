package auth

type UserName string

type UserId int

type User struct {
	Id   UserId
	Name UserName
}
