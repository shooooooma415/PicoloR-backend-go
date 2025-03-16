package v1

import (
	"encoding/json"
	"net/http"
	colorApp "picolor-backend/app/application/color"
	auth "picolor-backend/app/domain/auth"
	"picolor-backend/app/domain/color"
)

type GetThemeColorsResponse struct {
	ThemeColors []color.ColorCode `json:"themeColors"`
}

type GetThemeColorsParams struct {
	RoomID auth.RoomID `json:"roomID"`
}

type GetThemeColors struct {
	service *colorApp.ColorServiceImpl
}

func NewGetThemeColor(service *colorApp.ColorServiceImpl) *GetThemeColors {
	return &GetThemeColors{service: service}
}

func (g *GetThemeColors) GetThemeColor(params GetThemeColorsParams) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		themeColors, err := g.service.GetThemeColor(params.RoomID)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		res := GetThemeColorsResponse{
			ThemeColors: themeColors,
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(res)
	}
}
