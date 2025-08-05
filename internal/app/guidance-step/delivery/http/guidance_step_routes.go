package http

import (
	"github.com/labstack/echo/v4"
)

func (h *Handler) RegisterRoutes(g *echo.Group) {
	g.POST("", h.CreateGuidanceStep())
	g.GET("", h.GetGuidanceGuidanceSteps())
	g.GET("/:id", h.GetGuidanceGuidanceStep())
}
