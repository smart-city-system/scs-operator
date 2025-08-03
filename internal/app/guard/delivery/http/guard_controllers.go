package http

import (
	services "smart-city/internal/app/guard/service"
	"smart-city/internal/models"

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

func (h *Handler) Create() echo.HandlerFunc {
	return func(c echo.Context) error {
		guard := &models.User{Role: "guard"}
		if err := c.Bind(guard); err != nil {
			return err
		}
		createdGuard, err := h.svc.Create(c.Request().Context(), guard)
		if err != nil {
			return err
		}
		return c.JSON(201, createdGuard)
	}
}

func (h *Handler) GetGuard() echo.HandlerFunc {
	return func(c echo.Context) error {
		guards, err := h.svc.GetGuards(c.Request().Context())
		if err != nil {
			return err
		}
		return c.JSON(200, guards)
	}
}
