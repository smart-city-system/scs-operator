package http

import (
	"scs-operator/internal/app/user/dto"
	services "scs-operator/internal/app/user/service"
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

func (h *Handler) CreateUser() echo.HandlerFunc {
	return func(c echo.Context) error {
		createUserDto := &dto.CreateUserDto{}
		if err := c.Bind(createUserDto); err != nil {
			return errors.NewBadRequestError("Invalid request body")
		}

		// Validate the DTO
		if err := validation.ValidateStruct(createUserDto); err != nil {
			return err
		}

		createdUser, err := h.svc.CreateUser(c.Request().Context(), createUserDto)
		if err != nil {
			return err
		}
		createdUser.Password = ""
		return c.JSON(201, createdUser)
	}
}

func (h *Handler) GetUsers() echo.HandlerFunc {
	return func(c echo.Context) error {
		guards, err := h.svc.GetUsers(c.Request().Context())
		if err != nil {
			return err
		}
		return c.JSON(200, guards)
	}
}

func (h *Handler) GetAssignments() echo.HandlerFunc {
	return func(c echo.Context) error {
		userID := "72b194cd-3cb1-4653-b7d5-ed2fc032ed62"
		assignments, err := h.svc.GetAssignments(c.Request().Context(), userID)
		if err != nil {
			return err
		}
		return c.JSON(200, assignments)
	}
}
func (h *Handler) CompleteStep() echo.HandlerFunc {
	return func(c echo.Context) error {
		assignmentId := c.Param("assignmentId")
		stepId := c.Param("stepId")
		err := h.svc.CompleteStep(c.Request().Context(), assignmentId, stepId)
		if err != nil {
			return err
		}
		return c.JSON(200, "success")
	}
}
