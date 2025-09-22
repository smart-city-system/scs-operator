package http

import (
	"scs-operator/internal/app/guidance-template/dto"
	services "scs-operator/internal/app/guidance-template/service"
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

// CreateGuidanceTemplate creates a new guidance template
// @Summary Create a new guidance template
// @Description Create a new guidance template with steps
// @Tags guidance-templates
// @Accept json
// @Produce json
// @Param template body dto.CreateGuidanceTemplateDto true "Guidance template creation data"
// @Success 201 {object} models.GuidanceTemplate
// @Failure 400 {object} errors.ErrorResponse
// @Failure 500 {object} errors.ErrorResponse
// @Security BearerAuth
// @Router /guidance-templates [post]
func (h *Handler) CreateGuidanceTemplate() echo.HandlerFunc {
	return func(c echo.Context) error {
		createGuidanceTemplateDto := &dto.CreateGuidanceTemplateDto{}
		if err := c.Bind(createGuidanceTemplateDto); err != nil {
			return errors.NewBadRequestError("Invalid request body")
		}

		// Validate the DTO
		if err := validation.ValidateStruct(createGuidanceTemplateDto); err != nil {
			return err
		}

		createdGuidanceTemplate, err := h.svc.CreateGuidanceTemplate(c.Request().Context(), createGuidanceTemplateDto)
		if err != nil {
			return err
		}
		return c.JSON(201, createdGuidanceTemplate)
	}
}

// GetGuidanceGuidanceTemplates retrieves all guidance templates
// @Summary Get all guidance templates
// @Description Get a list of all guidance templates
// @Tags guidance-templates
// @Accept json
// @Produce json
// @Success 200 {object} types.GuidanceTemplateListResponse
// @Failure 500 {object} errors.ErrorResponse
// @Security BearerAuth
// @Router /guidance-templates [get]
func (h *Handler) GetGuidanceGuidanceTemplates() echo.HandlerFunc {
	return func(c echo.Context) error {
		guidanceTemplates, err := h.svc.GetGuidanceTemplates(c.Request().Context())
		if err != nil {
			return err
		}
		return c.JSON(200, guidanceTemplates)
	}
}

// GetGuidanceGuidanceTemplate retrieves a specific guidance template by ID
// @Summary Get guidance template by ID
// @Description Get a specific guidance template by its ID
// @Tags guidance-templates
// @Accept json
// @Produce json
// @Param id path string true "Guidance Template ID"
// @Success 200 {object} models.GuidanceTemplate
// @Failure 400 {object} errors.ErrorResponse
// @Failure 404 {object} errors.ErrorResponse
// @Failure 500 {object} errors.ErrorResponse
// @Security BearerAuth
// @Router /guidance-templates/{id} [get]
func (h *Handler) GetGuidanceGuidanceTemplate() echo.HandlerFunc {
	return func(c echo.Context) error {
		guidanceTemplateID := c.Param("id")
		guidanceTemplate, err := h.svc.GetGuidanceTemplateByID(c.Request().Context(), guidanceTemplateID)
		if err != nil {
			return err
		}
		return c.JSON(200, guidanceTemplate)
	}
}

// UpdateGuidanceTemplate updates an existing guidance template
// @Summary Update guidance template
// @Description Update an existing guidance template and its steps
// @Tags guidance-templates
// @Accept json
// @Produce json
// @Param id path string true "Guidance Template ID"
// @Param template body dto.UpdateGuidanceTemplateDto true "Guidance template update data"
// @Success 200 {object} models.GuidanceTemplate
// @Failure 400 {object} errors.ErrorResponse
// @Failure 404 {object} errors.ErrorResponse
// @Failure 500 {object} errors.ErrorResponse
// @Security BearerAuth
// @Router /guidance-templates/{id} [put]
func (h *Handler) UpdateGuidanceTemplate() echo.HandlerFunc {
	return func(c echo.Context) error {
		guidanceTemplateID := c.Param("id")
		updateGuidanceTemplateDto := &dto.UpdateGuidanceTemplateDto{}
		if err := c.Bind(updateGuidanceTemplateDto); err != nil {
			return errors.NewBadRequestError("Invalid request body")
		}

		// Validate the DTO
		if err := validation.ValidateStruct(updateGuidanceTemplateDto); err != nil {
			return err
		}
		updatedGuidanceTemplate, err := h.svc.UpdateGuidanceTemplate(c.Request().Context(), guidanceTemplateID, updateGuidanceTemplateDto)
		if err != nil {
			return err
		}
		return c.JSON(200, updatedGuidanceTemplate)

	}
}
