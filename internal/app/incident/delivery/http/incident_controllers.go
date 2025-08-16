package http

import (
	"scs-operator/internal/app/incident/dto"
	services "scs-operator/internal/app/incident/service"
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
func (h *Handler) UpdateIncident() echo.HandlerFunc {
	return func(c echo.Context) error {
		updateIncidentDto := &dto.UpdateIncidentDto{}
		incidentID := c.Param("id")
		if err := c.Bind(updateIncidentDto); err != nil {
			return errors.NewBadRequestError("Invalid request body")
		}

		// Validate the DTO
		if err := validation.ValidateStruct(updateIncidentDto); err != nil {
			return err
		}
		incident, err := h.svc.UpdateIncident(c.Request().Context(), incidentID, updateIncidentDto)
		if err != nil {
			return err
		}
		return c.JSON(200, incident)
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

func (h *Handler) AssignGuidance() echo.HandlerFunc {
	return func(c echo.Context) error {
		incidentID := c.Param("id")
		assignGuidance := &dto.AssignGuidance{}
		if err := c.Bind(assignGuidance); err != nil {
			return errors.NewBadRequestError("Invalid request body")
		}

		if err := validation.ValidateStruct(assignGuidance); err != nil {
			return err
		}
		createdIncidentGuidance, err := h.svc.AssignGuidance(c.Request().Context(), incidentID, assignGuidance)
		if err != nil {
			return err
		}
		return c.JSON(201, createdIncidentGuidance)
	}
}
func (h *Handler) GetIncidentGuidance() echo.HandlerFunc {
	return func(c echo.Context) error {
		incidentID := c.Param("id")
		incidentGuidance, err := h.svc.GetIncidentGuidance(c.Request().Context(), incidentID)
		if err != nil {
			return err
		}
		return c.JSON(200, incidentGuidance)
	}
}
