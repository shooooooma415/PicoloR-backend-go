package post

import (
	"picolor-backend/app/domain/auth"
	"picolor-backend/app/domain/color"
)

type Image string

type Post struct {
	UserID     auth.UserID
	RoomID     auth.RoomID
	ColorID    color.ColorID
	Rank       int
	Image      Image
	PostedTime string
}

type GetPost struct {
	UserName   auth.UserName
	Rank       int
	RoomID     auth.RoomID
	Image      Image
	PostedTime string
}
