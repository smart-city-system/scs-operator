package http

import (
	"scs-operator/internal/app/alarm/dto"
	services "scs-operator/internal/app/alarm/service"
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

func (h *Handler) GetAlarms() echo.HandlerFunc {
	return func(c echo.Context) error {
		status := c.QueryParam("status")
		premises, err := h.svc.GetAlarms(c.Request().Context(), status)
		if err != nil {
			return err
		}
		return c.JSON(200, premises)
	}
}

func (h *Handler) UpdateAlarm() echo.HandlerFunc {
	return func(c echo.Context) error {
		alarmID := c.Param("id")
		updateAlarmDto := &dto.UpdateAlarmDto{}
		if err := c.Bind(updateAlarmDto); err != nil {
			return errors.NewBadRequestError("Invalid request body")
		}

		// Validate the DTO
		if err := validation.ValidateStruct(updateAlarmDto); err != nil {
			return err
		}
		alarm, err := h.svc.UpdateAlarm(c.Request().Context(), alarmID, updateAlarmDto)
		if err != nil {
			return err
		}
		return c.JSON(200, alarm)

	}
}
