package http

import (
	"smart-city/internal/app/user/dto"
	services "smart-city/internal/app/user/service"
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
