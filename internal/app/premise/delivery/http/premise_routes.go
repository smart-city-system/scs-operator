package http

import (
	"github.com/labstack/echo/v4"
)

func (h *Handler) RegisterRoutes(g *echo.Group) {
	g.POST("", h.CreatePremise())
	g.PATCH("/:id", h.UpdatePremise())
	g.POST("/:id/assign-users", h.AssignUsers())
	g.GET("", h.GetPremises())
	g.GET("/:id", h.GetPremise())
	g.GET("/:id/users", h.GetAvailableUsers())
}
