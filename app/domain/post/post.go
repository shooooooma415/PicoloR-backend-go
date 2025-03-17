package post

import (
	"picolor-backend/app/domain/auth"
)

type Image string

type Post struct {
	UserID     auth.UserID
	RoomID     auth.RoomID
	ColorID    auth.ColorID
	Rank       int
	Image      Image
	PostedTime string
}

type GetPost struct {
	UserName   auth.UserName
	Rank       int
	RoomID     auth.RoomID
	Color      auth.ColorCode
	Image      Image
	PostedTime string
}
