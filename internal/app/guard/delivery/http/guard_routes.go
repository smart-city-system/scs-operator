package http

import (
	"github.com/labstack/echo/v4"
)

func (h *Handler) RegisterRoutes(g *echo.Group) {
	g.POST("", h.Create())
	g.GET("", h.GetGuard())
	g.POST("/assign-premises", h.AssignPremises())
}
