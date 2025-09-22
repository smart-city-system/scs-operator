package http

import (
	"scs-operator/internal/app/premise/dto"
	services "scs-operator/internal/app/premise/service"
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

// CreatePremise creates a new premise
// @Summary Create a new premise
// @Description Create a new premise with the provided information
// @Tags premises
// @Accept json
// @Produce json
// @Param premise body dto.CreatePremiseDto true "Premise creation data"
// @Success 201 {object} models.Premise
// @Failure 400 {object} errors.ErrorResponse
// @Failure 500 {object} errors.ErrorResponse
// @Router /premises [post]
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

// GetPremises retrieves a paginated list of premises
// @Summary Get premises with pagination
// @Description Get a paginated list of all premises
// @Tags premises
// @Accept json
// @Produce json
// @Param page query int false "Page number" default(1)
// @Param limit query int false "Number of items per page" default(10)
// @Success 200 {object} types.PremiseListResponse
// @Failure 400 {object} errors.ErrorResponse
// @Failure 500 {object} errors.ErrorResponse
// @Router /premises [get]
func (h *Handler) GetPremises() echo.HandlerFunc {
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
		premises, err := h.svc.GetPremises(c.Request().Context(), pageInt, limitInt)
		if err != nil {
			return err
		}
		return c.JSON(200, premises)
	}
}

// GetPremise retrieves a specific premise by ID
// @Summary Get premise by ID
// @Description Get a specific premise by its ID
// @Tags premises
// @Accept json
// @Produce json
// @Param id path string true "Premise ID"
// @Success 200 {object} models.Premise
// @Failure 400 {object} errors.ErrorResponse
// @Failure 404 {object} errors.ErrorResponse
// @Failure 500 {object} errors.ErrorResponse
// @Router /premises/{id} [get]
func (h *Handler) GetPremise() echo.HandlerFunc {
	return func(c echo.Context) error {
		premiseID := c.Param("id")
		premise, err := h.svc.GetPremiseByID(c.Request().Context(), premiseID)
		if err != nil {
			return err
		}
		return c.JSON(200, premise)
	}
}

// GetAvailableUsers retrieves users assigned to a premise
// @Summary Get users assigned to premise
// @Description Get all users assigned to a specific premise
// @Tags premises
// @Accept json
// @Produce json
// @Param id path string true "Premise ID"
// @Success 200 {object} types.UserListResponse
// @Failure 400 {object} errors.ErrorResponse
// @Failure 404 {object} errors.ErrorResponse
// @Failure 500 {object} errors.ErrorResponse
// @Router /premises/{id}/users [get]
func (h *Handler) GetAvailableUsers() echo.HandlerFunc {
	return func(c echo.Context) error {
		premiseID := c.Param("id")
		guards, err := h.svc.GetAvailableUsers(c.Request().Context(), premiseID)
		if err != nil {
			return err
		}
		return c.JSON(200, guards)
	}
}

// UpdatePremise updates an existing premise
// @Summary Update premise
// @Description Update an existing premise with new information
// @Tags premises
// @Accept json
// @Produce json
// @Param id path string true "Premise ID"
// @Param premise body dto.UpdatePremiseDto true "Premise update data"
// @Success 200 {object} models.Premise
// @Failure 400 {object} errors.ErrorResponse
// @Failure 404 {object} errors.ErrorResponse
// @Failure 500 {object} errors.ErrorResponse
// @Router /premises/{id} [patch]
func (h *Handler) UpdatePremise() echo.HandlerFunc {
	return func(c echo.Context) error {
		premiseID := c.Param("id")
		updatePremiseDto := &dto.UpdatePremiseDto{}
		if err := c.Bind(updatePremiseDto); err != nil {
			return errors.NewBadRequestError("Invalid request body")
		}

		// Validate the DTO
		if err := validation.ValidateStruct(updatePremiseDto); err != nil {
			return err
		}
		updatedPremise, err := h.svc.UpdatePremise(c.Request().Context(), premiseID, updatePremiseDto)
		if err != nil {
			return err
		}
		return c.JSON(200, updatedPremise)

	}
}

// AssignUsers assigns or removes users from a premise
// @Summary Assign users to premise
// @Description Add or remove users from a premise
// @Tags premises
// @Accept json
// @Produce json
// @Param id path string true "Premise ID"
// @Param users body dto.UpdatePremiseUserDto true "User assignment data"
// @Success 200 {string} string "success"
// @Failure 400 {object} errors.ErrorResponse
// @Failure 404 {object} errors.ErrorResponse
// @Failure 500 {object} errors.ErrorResponse
// @Router /premises/{id}/assign-users [post]
func (h *Handler) AssignUsers() echo.HandlerFunc {
	return func(c echo.Context) error {
		premiseID := c.Param("id")
		updatePremiseUserDto := &dto.UpdatePremiseUserDto{}
		if err := c.Bind(updatePremiseUserDto); err != nil {
			return errors.NewBadRequestError("Invalid request body")
		}

		// Validate the DTO
		if err := validation.ValidateStruct(updatePremiseUserDto); err != nil {
			return err
		}
		err := h.svc.AssignUsers(c.Request().Context(), premiseID, updatePremiseUserDto)
		if err != nil {
			return err
		}
		return c.JSON(200, "success")

	}
}
