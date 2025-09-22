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

// CreateGuidanceStep creates a new guidance step
// @Summary Create a new guidance step
// @Description Create a new guidance step for a guidance template
// @Tags guidance-steps
// @Accept json
// @Produce json
// @Param step body dto.CreateGuidanceStepDto true "Guidance step creation data"
// @Success 201 {object} models.GuidanceStep
// @Failure 400 {object} errors.ErrorResponse
// @Failure 500 {object} errors.ErrorResponse
// @Security BearerAuth
// @Router /guidance-steps [post]
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

// GetGuidanceGuidanceSteps retrieves all guidance steps
// @Summary Get all guidance steps
// @Description Get a list of all guidance steps
// @Tags guidance-steps
// @Accept json
// @Produce json
// @Success 200 {object} types.GuidanceStepListResponse
// @Failure 500 {object} errors.ErrorResponse
// @Security BearerAuth
// @Router /guidance-steps [get]
func (h *Handler) GetGuidanceGuidanceSteps() echo.HandlerFunc {
	return func(c echo.Context) error {
		guidanceSteps, err := h.svc.GetGuidanceSteps(c.Request().Context())
		if err != nil {
			return err
		}
		return c.JSON(200, guidanceSteps)
	}
}

// GetGuidanceGuidanceStep retrieves a specific guidance step by ID
// @Summary Get guidance step by ID
// @Description Get a specific guidance step by its ID
// @Tags guidance-steps
// @Accept json
// @Produce json
// @Param id path string true "Guidance Step ID"
// @Success 200 {object} models.GuidanceStep
// @Failure 400 {object} errors.ErrorResponse
// @Failure 404 {object} errors.ErrorResponse
// @Failure 500 {object} errors.ErrorResponse
// @Security BearerAuth
// @Router /guidance-steps/{id} [get]
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
