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

func (h *Handler) GetGuidanceGuidanceTemplates() echo.HandlerFunc {
	return func(c echo.Context) error {
		guidanceTemplates, err := h.svc.GetGuidanceTemplates(c.Request().Context())
		if err != nil {
			return err
		}
		return c.JSON(200, guidanceTemplates)
	}
}

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
