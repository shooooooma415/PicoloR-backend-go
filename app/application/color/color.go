package color

import "picolor-backend/app/domain/color"

type ColorServiceImpl struct {
	colorRepo color.ColorRepository
}

func NewColorService(colorRepo color.ColorRepository) *ColorServiceImpl {
	return &ColorServiceImpl{
		colorRepo: colorRepo,
	}
}
