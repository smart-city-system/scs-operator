package server

import (
	"net/http"

	premisesHttp "scs-operator/internal/app/premise/delivery/http"
	premisesRepository "scs-operator/internal/app/premise/repository"
	premisesService "scs-operator/internal/app/premise/service"

	usersRepository "scs-operator/internal/app/user/repository"

	incidentsHttp "scs-operator/internal/app/incident/delivery/http"
	incidentsRepository "scs-operator/internal/app/incident/repository"
	incidentsService "scs-operator/internal/app/incident/service"

	guidanceTemplatesHttp "scs-operator/internal/app/guidance-template/delivery/http"
	guidanceTemplatesRepository "scs-operator/internal/app/guidance-template/repository"
	guidanceTemplatesService "scs-operator/internal/app/guidance-template/service"

	guidanceStepsHttp "scs-operator/internal/app/guidance-step/delivery/http"
	guidanceStepsRepository "scs-operator/internal/app/guidance-step/repository"
	guidanceStepsService "scs-operator/internal/app/guidance-step/service"

	alarmsHttp "scs-operator/internal/app/alarm/delivery/http"
	alarmsRepository "scs-operator/internal/app/alarm/repository"
	alarmsService "scs-operator/internal/app/alarm/service"

	middleware "scs-operator/internal/middlewares"

	"github.com/labstack/echo/v4"
)

func (s *Server) MapHandlers(e *echo.Echo) error {
	// Init repositories
	usersRepo := usersRepository.NewUserRepository(s.db)
	premisesRepo := premisesRepository.NewPremiseRepository(s.db)
	incidentsRepo := incidentsRepository.NewIncidentRepository(s.db)
	guidanceTemplatesRepo := guidanceTemplatesRepository.NewGuidanceTemplateRepository(s.db)
	guidanceStepsRepo := guidanceStepsRepository.NewGuidanceStepRepository(s.db)
	incidentGuidanceRepo := incidentsRepository.NewIncidentGuidanceRepository(s.db)
	incidentGuidanceStepRepo := incidentsRepository.NewIncidentGuidanceStepRepository(s.db)
	alarmsRepo := alarmsRepository.NewAlarmRepository(s.db)
	// Init service
	premisesSvc := premisesService.NewPremiseService(*premisesRepo)
	incidentsSvc := incidentsService.NewIncidentService(*incidentsRepo, *incidentGuidanceRepo, *usersRepo, *guidanceTemplatesRepo, *incidentGuidanceStepRepo)
	guidanceTemplatesSvc := guidanceTemplatesService.NewGuidanceTemplateService(*guidanceTemplatesRepo, *guidanceStepsRepo)
	guidanceStepsSvc := guidanceStepsService.NewGuidanceStepService(*guidanceStepsRepo)
	alarmsSvc := alarmsService.NewAlarmService(*alarmsRepo, *premisesRepo)
	// Init handlers
	premisesHandlers := premisesHttp.NewHandler(*premisesSvc)
	incidentsHandlers := incidentsHttp.NewHandler(*incidentsSvc)
	guidanceTemplatesHandlers := guidanceTemplatesHttp.NewHandler(*guidanceTemplatesSvc)
	guidanceStepsHandlers := guidanceStepsHttp.NewHandler(*guidanceStepsSvc)
	alarmsHandlers := alarmsHttp.NewHandler(*alarmsSvc)

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
