package http

import (
	"github.com/labstack/echo/v4"
)

func (h *Handler) RegisterRoutes(g *echo.Group) {
	g.POST("", h.CreateGuidanceTemplate())
	g.GET("", h.GetGuidanceGuidanceTemplates())
	g.GET("/:id", h.GetGuidanceGuidanceTemplate())
}
