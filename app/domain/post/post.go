package post

import (
	"picolor-backend/app/domain/auth"
	"picolor-backend/app/domain/color"
	"time"
)

type Image string

type Post struct {
	UserId   auth.UserId
	ColorId  color.ColorId
	Rank     int
	Image    Image
	PostedAt time.Time
}
