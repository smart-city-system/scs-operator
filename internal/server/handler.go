package server

import (
	"net/http"

	premisesHttp "scs-operator/internal/app/premise/delivery/http"

	incidentsHttp "scs-operator/internal/app/incident/delivery/http"

	guidanceTemplatesHttp "scs-operator/internal/app/guidance-template/delivery/http"

	guidanceStepsHttp "scs-operator/internal/app/guidance-step/delivery/http"

	alarmsHttp "scs-operator/internal/app/alarm/delivery/http"

	middleware "scs-operator/internal/middlewares"

	"github.com/labstack/echo/v4"
)

func (s *Server) MapHandlers(e *echo.Echo) error {
	// Use shared services from container instead of creating new instances
	// Init handlers using shared services
	premisesHandlers := premisesHttp.NewHandler(*s.container.PremiseService)
	incidentsHandlers := incidentsHttp.NewHandler(*s.container.IncidentService)
	guidanceTemplatesHandlers := guidanceTemplatesHttp.NewHandler(*s.container.GuidanceTemplateService)
	guidanceStepsHandlers := guidanceStepsHttp.NewHandler(*s.container.GuidanceStepService)
	alarmsHandlers := alarmsHttp.NewHandler(*s.container.AlarmService)

	mw := middleware.NewMiddlewareManager(s.cfg, []string{"*"}, s.logger)
	e.Use(mw.RequestLoggerMiddleware)
	e.Use(mw.ErrorHandlerMiddleware)
	v1 := e.Group("/api/v1")

	health := v1.Group("/health")
	premisesGroup := v1.Group("/premises", mw.JWTAuth)
	incidentsGroup := v1.Group("/incidents", mw.JWTAuth)
	guidanceTemplatesGroup := v1.Group("/guidance-templates", mw.JWTAuth)
	guidanceStepsGroup := v1.Group("/guidance-steps", mw.JWTAuth)
	alarmsGroup := v1.Group("/alarms", mw.JWTAuth)

	health.GET("", func(c echo.Context) error {
		return c.JSON(http.StatusOK, map[string]string{"status": "OK"})
	})
	premisesHandlers.RegisterRoutes(premisesGroup)
	incidentsHandlers.RegisterRoutes(incidentsGroup)
	guidanceTemplatesHandlers.RegisterRoutes(guidanceTemplatesGroup)
	guidanceStepsHandlers.RegisterRoutes(guidanceStepsGroup)
	alarmsHandlers.RegisterRoutes(alarmsGroup)
	return nil

}
