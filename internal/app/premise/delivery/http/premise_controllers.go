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
