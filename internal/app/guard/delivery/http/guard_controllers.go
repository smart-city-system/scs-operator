package http

import (
	"scs-operator/internal/app/guard/dto"
	services "scs-operator/internal/app/guard/service"

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
		guard := &dto.CreateGuardDto{}
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
func (h *Handler) AssignPremises() echo.HandlerFunc {
	return func(c echo.Context) error {
		assignPremisesDto := &dto.AssignPremisesDto{}
		if err := c.Bind(assignPremisesDto); err != nil {
			return err
		}
		err := h.svc.AssignPremises(c.Request().Context(), assignPremisesDto.GuardID, assignPremisesDto.PremiseID)
		if err != nil {
			return err
		}
		return c.JSON(200, "success")
	}
}
