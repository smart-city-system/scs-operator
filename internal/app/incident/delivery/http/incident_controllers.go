package http

import (
	"smart-city/internal/app/incident/dto"
	services "smart-city/internal/app/incident/service"
	"smart-city/pkg/errors"
	"smart-city/pkg/validation"

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

func (h *Handler) CreateIncident() echo.HandlerFunc {
	return func(c echo.Context) error {
		createIncidentDto := &dto.CreateIncidentDto{}
		if err := c.Bind(createIncidentDto); err != nil {
			return errors.NewBadRequestError("Invalid request body")
		}

		// Validate the DTO
		if err := validation.ValidateStruct(createIncidentDto); err != nil {
			return err
		}

		createdIncident, err := h.svc.CreateIncident(c.Request().Context(), createIncidentDto)
		if err != nil {
			return err
		}
		return c.JSON(201, createdIncident)
	}
}

func (h *Handler) GetIncidents() echo.HandlerFunc {
	return func(c echo.Context) error {
		incidents, err := h.svc.GetIncidents(c.Request().Context())
		if err != nil {
			return err
		}
		return c.JSON(200, incidents)
	}
}

func (h *Handler) GetIncident() echo.HandlerFunc {
	return func(c echo.Context) error {
		incidentID := c.Param("id")
		incident, err := h.svc.GetIncidentByID(c.Request().Context(), incidentID)
		if err != nil {
			return err
		}
		return c.JSON(200, incident)
	}
}
