package http

import (
	"github.com/labstack/echo/v4"
)

func (h *Handler) RegisterRoutes(g *echo.Group) {
	g.POST("", h.CreatePremise())
	g.GET("", h.GetPremises())
	g.GET("/:id/guards", h.GetAvailableGuards())
}
