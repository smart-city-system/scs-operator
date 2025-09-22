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

// GetAlarms retrieves alarms with optional status filtering
// @Summary Get alarms
// @Description Get all alarms with optional status filtering
// @Tags alarms
// @Accept json
// @Produce json
// @Param status query string false "Filter by alarm status (new, ignored, dispatched)"
// @Success 200 {object} types.AlarmListResponse
// @Failure 400 {object} errors.ErrorResponse
// @Failure 500 {object} errors.ErrorResponse
// @Security BearerAuth
// @Router /alarms [get]
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

// UpdateAlarm updates an existing alarm
// @Summary Update alarm
// @Description Update an existing alarm's status or other properties
// @Tags alarms
// @Accept json
// @Produce json
// @Param id path string true "Alarm ID"
// @Param alarm body dto.UpdateAlarmDto true "Alarm update data"
// @Success 200 {object} models.Alarm
// @Failure 400 {object} errors.ErrorResponse
// @Failure 404 {object} errors.ErrorResponse
// @Failure 500 {object} errors.ErrorResponse
// @Security BearerAuth
// @Router /alarms/{id} [patch]
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
