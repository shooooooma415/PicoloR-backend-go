package v1

import (
	"encoding/json"
	"net/http"
	colorApp "picolor-backend/app/application/color"
	"picolor-backend/app/domain/auth"
	"strconv"
)

type ColorRes struct{
	ColorId   auth.ColorID
	ColorCode auth.ColorCode
}

type GetThemeColorsResponse struct {
	ThemeColors []ColorRes `json:"themeColors"`
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

		themeColors, err := g.service.GetThemeColors(auth.RoomID(roomID))
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		var res GetThemeColorsResponse
		for _, themeColor := range themeColors {
			res.ThemeColors = append(res.ThemeColors, ColorRes{
				ColorId:   themeColor.ColorId,
				ColorCode: themeColor.ColorCode,
			})
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(res)
	}
}
