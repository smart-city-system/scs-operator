package http

import (
	"github.com/labstack/echo/v4"
)

func (h *Handler) RegisterRoutes(g *echo.Group) {
	g.POST("", h.CreateAsset())
	g.GET("", h.GetAssets())
	g.POST("/publish", h.StartPublishing())
	g.POST("/snapshot", h.CreateSnapshot())

}
