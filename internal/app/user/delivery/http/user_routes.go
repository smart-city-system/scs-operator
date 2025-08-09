package http

import (
	"github.com/labstack/echo/v4"
)

func (h *Handler) RegisterRoutes(g *echo.Group) {
	g.POST("", h.CreateUser())
	g.GET("", h.GetUsers())
	g.GET("/me/assignments", h.GetAssignments())
	g.PATCH("/me/:assignmentId/steps/:stepId", h.CompleteStep())
}
