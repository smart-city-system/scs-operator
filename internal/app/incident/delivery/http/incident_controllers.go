package http

import (
	"scs-operator/internal/app/incident/dto"
	services "scs-operator/internal/app/incident/service"
	"scs-operator/pkg/errors"
	"scs-operator/pkg/validation"
	"strconv"

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

// CreateIncident creates a new incident
// @Summary Create a new incident
// @Description Create a new incident with the provided information
// @Tags incidents
// @Accept json
// @Produce json
// @Param incident body dto.CreateIncidentDto true "Incident creation data"
// @Success 201 {object} models.Incident
// @Failure 400 {object} errors.ErrorResponse
// @Failure 500 {object} errors.ErrorResponse
// @Security BearerAuth
// @Router /incidents [post]
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

// UpdateIncident updates an existing incident
// @Summary Update incident
// @Description Update an existing incident's status or other properties
// @Tags incidents
// @Accept json
// @Produce json
// @Param id path string true "Incident ID"
// @Param incident body dto.UpdateIncidentDto true "Incident update data"
// @Success 200 {object} models.Incident
// @Failure 400 {object} errors.ErrorResponse
// @Failure 404 {object} errors.ErrorResponse
// @Failure 500 {object} errors.ErrorResponse
// @Security BearerAuth
// @Router /incidents/{id} [patch]
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

// GetIncidents retrieves a paginated list of incidents
// @Summary Get incidents with pagination
// @Description Get a paginated list of all incidents
// @Tags incidents
// @Accept json
// @Produce json
// @Param page query int false "Page number" default(1)
// @Param limit query int false "Number of items per page" default(10)
// @Success 200 {object} types.IncidentListResponse
// @Failure 400 {object} errors.ErrorResponse
// @Failure 500 {object} errors.ErrorResponse
// @Security BearerAuth
// @Router /incidents [get]
func (h *Handler) GetIncidents() echo.HandlerFunc {
	return func(c echo.Context) error {
		page := c.QueryParam("page")
		limit := c.QueryParam("limit")
		if page == "" {
			page = "1"
		}
		if limit == "" {
			limit = "10"
		}
		pageInt, err := strconv.Atoi(page)
		if err != nil {
			return errors.NewBadRequestError("Invalid page number")
		}
		limitInt, err := strconv.Atoi(limit)
		if err != nil {
			return errors.NewBadRequestError("Invalid limit number")
		}

		incidents, err := h.svc.GetIncidents(c.Request().Context(), pageInt, limitInt)
		if err != nil {
			return err
		}
		return c.JSON(200, incidents)
	}
}

// GetIncident retrieves a specific incident by ID
// @Summary Get incident by ID
// @Description Get a specific incident by its ID
// @Tags incidents
// @Accept json
// @Produce json
// @Param id path string true "Incident ID"
// @Success 200 {object} models.Incident
// @Failure 400 {object} errors.ErrorResponse
// @Failure 404 {object} errors.ErrorResponse
// @Failure 500 {object} errors.ErrorResponse
// @Security BearerAuth
// @Router /incidents/{id} [get]
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

// AssignGuidance assigns guidance template to an incident
// @Summary Assign guidance to incident
// @Description Assign a guidance template to an incident with an assignee
// @Tags incidents
// @Accept json
// @Produce json
// @Param id path string true "Incident ID"
// @Param guidance body dto.AssignGuidance true "Guidance assignment data"
// @Success 201 {object} models.IncidentGuidance
// @Failure 400 {object} errors.ErrorResponse
// @Failure 404 {object} errors.ErrorResponse
// @Failure 500 {object} errors.ErrorResponse
// @Security BearerAuth
// @Router /incidents/{id}/assign-guidance [post]
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

// GetIncidentGuidance retrieves guidance for an incident
// @Summary Get incident guidance
// @Description Get the guidance assigned to a specific incident
// @Tags incidents
// @Accept json
// @Produce json
// @Param id path string true "Incident ID"
// @Success 200 {object} models.IncidentGuidance
// @Failure 400 {object} errors.ErrorResponse
// @Failure 404 {object} errors.ErrorResponse
// @Failure 500 {object} errors.ErrorResponse
// @Security BearerAuth
// @Router /incidents/{id}/guidance [get]
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

// CompleteIncident marks an incident as completed
// @Summary Complete incident
// @Description Mark an incident as resolved/completed
// @Tags incidents
// @Accept json
// @Produce json
// @Param id path string true "Incident ID"
// @Success 200 {string} string "success"
// @Failure 400 {object} errors.ErrorResponse
// @Failure 404 {object} errors.ErrorResponse
// @Failure 500 {object} errors.ErrorResponse
// @Security BearerAuth
// @Router /incidents/{id}/complete [patch]
func (h *Handler) CompleteIncident() echo.HandlerFunc {
	return func(c echo.Context) error {
		incidentID := c.Param("id")
		err := h.svc.CompleteIncident(c.Request().Context(), incidentID)
		if err != nil {
			return err
		}
		return c.JSON(200, "success")
	}
}
