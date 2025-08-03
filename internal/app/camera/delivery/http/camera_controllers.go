package http

import (
	"smart-city/internal/app/camera/dto"
	services "smart-city/internal/app/camera/service"

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

func (h *Handler) CreateCamera() echo.HandlerFunc {
	return func(c echo.Context) error {
		createCameraDto := &dto.CreateCameraDto{}
		if err := c.Bind(createCameraDto); err != nil {
			return err
		}
		createdCamera, err := h.svc.CreateCamera(c.Request().Context(), createCameraDto)
		if err != nil {
			return err
		}
		return c.JSON(201, createdCamera)
	}
}

func (h *Handler) GetCameras() echo.HandlerFunc {
	return func(c echo.Context) error {
		cameras, err := h.svc.GetCameras(c.Request().Context())
		if err != nil {
			return err
		}
		return c.JSON(200, cameras)
	}
}
