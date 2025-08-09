package http

import (
	services "scs-operator/internal/app/alarm/service"

	"github.com/labstack/echo/v4"
)

// Handler
type Handler struct {
	svc services.Service
}

// NewHandler constructor
func NewHandler(svc services.Service) *Handler {
	return &Handler{svc: svc}
}

func (h *Handler) GetAlarms() echo.HandlerFunc {
	return func(c echo.Context) error {
		premises, err := h.svc.GetAlarms(c.Request().Context())
		if err != nil {
			return err
		}
		return c.JSON(200, premises)
	}
}
