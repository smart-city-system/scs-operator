package server

import (
	"net/http"
	assetsHttp "smart-city/internal/app/asset/delivery/http"
	assetsRepository "smart-city/internal/app/asset/repository"
	assetsService "smart-city/internal/app/asset/service"

	premisesHttp "smart-city/internal/app/premise/delivery/http"
	premisesRepository "smart-city/internal/app/premise/repository"
	premisesService "smart-city/internal/app/premise/service"

	usersHttp "smart-city/internal/app/user/delivery/http"
	usersRepository "smart-city/internal/app/user/repository"
	usersService "smart-city/internal/app/user/service"

	incidentsHttp "smart-city/internal/app/incident/delivery/http"
	incidentsRepository "smart-city/internal/app/incident/repository"
	incidentsService "smart-city/internal/app/incident/service"

	guidanceTemplatesHttp "smart-city/internal/app/guidance-template/delivery/http"
	guidanceTemplatesRepository "smart-city/internal/app/guidance-template/repository"
	guidanceTemplatesService "smart-city/internal/app/guidance-template/service"

	guidanceStepsHttp "smart-city/internal/app/guidance-step/delivery/http"
	guidanceStepsRepository "smart-city/internal/app/guidance-step/repository"
	guidanceStepsService "smart-city/internal/app/guidance-step/service"

	middleware "smart-city/internal/middlewares"

	"github.com/labstack/echo/v4"
)

func (s *Server) MapHandlers(e *echo.Echo) error {
	// Init repositories
	usersRepo := usersRepository.NewUserRepository(s.db)
	premisesRepo := premisesRepository.NewPremiseRepository(s.db)
	assetsRepo := assetsRepository.NewAssetRepository(s.db)
	snapshotsRepo := assetsRepository.NewSnapshotRepository(s.db)
	incidentsRepo := incidentsRepository.NewIncidentRepository(s.db)
	guidanceTemplatesRepo := guidanceTemplatesRepository.NewGuidanceTemplateRepository(s.db)
	guidanceStepsRepo := guidanceStepsRepository.NewGuidanceStepRepository(s.db)
	// Init service
	usersSvc := usersService.NewUserService(*usersRepo)
	premisesSvc := premisesService.NewPremiseService(*premisesRepo)
	assetsSvc := assetsService.NewAssetService(*assetsRepo, *premisesRepo, *snapshotsRepo)
	incidentsSvc := incidentsService.NewIncidentService(*incidentsRepo)
	guidanceTemplatesSvc := guidanceTemplatesService.NewGuidanceTemplateService(*guidanceTemplatesRepo, *guidanceStepsRepo)
	guidanceStepsSvc := guidanceStepsService.NewGuidanceStepService(*guidanceStepsRepo)
	// Init handlers
	usersHandlers := usersHttp.NewHandler(*usersSvc)
	premisesHandlers := premisesHttp.NewHandler(*premisesSvc)
	assetsHandlers := assetsHttp.NewHandler(*assetsSvc)
	incidentsHandlers := incidentsHttp.NewHandler(*incidentsSvc)
	guidanceTemplatesHandlers := guidanceTemplatesHttp.NewHandler(*guidanceTemplatesSvc)
	guidanceStepsHandlers := guidanceStepsHttp.NewHandler(*guidanceStepsSvc)

	mw := middleware.NewMiddlewareManager(s.cfg, []string{"*"}, s.logger)
	e.Use(mw.RequestLoggerMiddleware)
	e.Use(mw.ErrorHandlerMiddleware)
	v1 := e.Group("/api/v1")

	health := v1.Group("/health")
	usersGroup := v1.Group("/users")
	premisesGroup := v1.Group("/premises")
	assetsGroup := v1.Group("/assets")
	incidentsGroup := v1.Group("/incidents")
	guidanceTemplatesGroup := v1.Group("/guidance-templates")
	guidanceStepsGroup := v1.Group("/guidance-steps")

	health.GET("", func(c echo.Context) error {
		return c.JSON(http.StatusOK, map[string]string{"status": "OK"})
	})
	usersHandlers.RegisterRoutes(usersGroup)
	premisesHandlers.RegisterRoutes(premisesGroup)
	assetsHandlers.RegisterRoutes(assetsGroup)
	incidentsHandlers.RegisterRoutes(incidentsGroup)
	guidanceTemplatesHandlers.RegisterRoutes(guidanceTemplatesGroup)
	guidanceStepsHandlers.RegisterRoutes(guidanceStepsGroup)
	return nil

}
