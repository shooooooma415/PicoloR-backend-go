package v1

import (
	"encoding/json"
	"net/http"
	colorApp "picolor-backend/app/application/color"
	"picolor-backend/app/domain/auth"
	"strconv"
)

type GetThemeColorsResponse struct {
	ThemeColors []auth.ColorCode `json:"themeColors"`
}

type GetThemeColors struct {
	service *colorApp.ColorServiceImpl
}

func NewGetThemeColor(service *colorApp.ColorServiceImpl) *GetThemeColors {
	return &GetThemeColors{service: service}
}

func (g *GetThemeColors) GetThemeColorHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		roomIDStr := r.URL.Query().Get("roomID")
		roomID, err := strconv.Atoi(roomIDStr)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		themeColors, err := g.service.GetThemeColor(auth.RoomID(roomID))
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
