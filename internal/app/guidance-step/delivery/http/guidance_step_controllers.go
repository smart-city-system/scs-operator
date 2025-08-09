package http

import (
	"scs-operator/internal/app/guidance-step/dto"
	services "scs-operator/internal/app/guidance-step/service"
	"scs-operator/pkg/errors"
	"scs-operator/pkg/validation"

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

func (h *Handler) CreateGuidanceStep() echo.HandlerFunc {
	return func(c echo.Context) error {
		createGuidanceStepDto := &dto.CreateGuidanceStepDto{}
		if err := c.Bind(createGuidanceStepDto); err != nil {
			return errors.NewBadRequestError("Invalid request body")
		}

		// Validate the DTO
		if err := validation.ValidateStruct(createGuidanceStepDto); err != nil {
			return err
		}

		createdGuidanceStep, err := h.svc.CreateGuidanceStep(c.Request().Context(), createGuidanceStepDto)
		if err != nil {
			return err
		}
		return c.JSON(201, createdGuidanceStep)
	}
}

func (h *Handler) GetGuidanceGuidanceSteps() echo.HandlerFunc {
	return func(c echo.Context) error {
		guidanceSteps, err := h.svc.GetGuidanceSteps(c.Request().Context())
		if err != nil {
			return err
		}
		return c.JSON(200, guidanceSteps)
	}
}

func (h *Handler) GetGuidanceGuidanceStep() echo.HandlerFunc {
	return func(c echo.Context) error {
		guidanceStepID := c.Param("id")
		guidanceStep, err := h.svc.GetGuidanceStepByID(c.Request().Context(), guidanceStepID)
		if err != nil {
			return err
		}
		return c.JSON(200, guidanceStep)
	}
}
