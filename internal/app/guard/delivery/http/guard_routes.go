package http

import (
	"github.com/labstack/echo/v4"
)

func (h *Handler) RegisterRoutes(g *echo.Group) {
	g.POST("/guards", h.Create())
	g.GET("/guards", h.GetGuard())
}
