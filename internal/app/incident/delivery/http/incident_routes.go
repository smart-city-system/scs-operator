package http

import (
	"github.com/labstack/echo/v4"
)

func (h *Handler) RegisterRoutes(g *echo.Group) {
	g.POST("", h.CreateIncident())
	g.GET("", h.GetIncidents())
	g.GET("/:id", h.GetIncident())
	g.POST("/:id/assign-guidance", h.AssignGuidance())
	g.GET("/:id/guidance", h.GetIncidentGuidance())
}
