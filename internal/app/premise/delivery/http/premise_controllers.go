package http

import (
	"scs-operator/internal/app/premise/dto"
	services "scs-operator/internal/app/premise/service"
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

func (h *Handler) CreatePremise() echo.HandlerFunc {
	return func(c echo.Context) error {
		createPremiseDto := &dto.CreatePremiseDto{}
		if err := c.Bind(createPremiseDto); err != nil {
			return errors.NewBadRequestError("Invalid request body")
		}

		// Validate the DTO
		if err := validation.ValidateStruct(createPremiseDto); err != nil {
			return err
		}

		createdPremise, err := h.svc.CreatePremise(c.Request().Context(), createPremiseDto)
		if err != nil {
			return err
		}
		return c.JSON(201, createdPremise)
	}
}

func (h *Handler) GetPremises() echo.HandlerFunc {
	return func(c echo.Context) error {
		premises, err := h.svc.GetPremises(c.Request().Context())
		if err != nil {
			return err
		}
		return c.JSON(200, premises)
	}
}
