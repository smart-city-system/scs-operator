package server

import (
	"net/http"
	camerasHttp "smart-city/internal/app/camera/delivery/http"
	camerasRepository "smart-city/internal/app/camera/repository"
	camerasService "smart-city/internal/app/camera/service"

	premisesHttp "smart-city/internal/app/premise/delivery/http"
	premisesRepository "smart-city/internal/app/premise/repository"
	premisesService "smart-city/internal/app/premise/service"

	usersHttp "smart-city/internal/app/user/delivery/http"
	usersRepository "smart-city/internal/app/user/repository"
	usersService "smart-city/internal/app/user/service"

	middleware "smart-city/internal/middlewares"

	"github.com/labstack/echo/v4"
)

func (s *Server) MapHandlers(e *echo.Echo) error {
	// Init repositories
	usersRepo := usersRepository.NewUserRepository(s.db)
	premisesRepo := premisesRepository.NewPremiseRepository(s.db)
	camerasRepo := camerasRepository.NewCameraRepository(s.db)
	// Init service
	usersSvc := usersService.NewUserService(*usersRepo)
	premisesSvc := premisesService.NewPremiseService(*premisesRepo)
	camerasSvc := camerasService.NewCameraService(*camerasRepo, *premisesRepo)
	// Init handlers
	usersHandlers := usersHttp.NewHandler(*usersSvc)
	premisesHandlers := premisesHttp.NewHandler(*premisesSvc)
	camerasHandlers := camerasHttp.NewHandler(*camerasSvc)

	mw := middleware.NewMiddlewareManager(s.cfg, []string{"*"}, s.logger)
	e.Use(mw.RequestLoggerMiddleware)
	v1 := e.Group("/api/v1")

	health := v1.Group("/health")
	usersGroup := v1.Group("/users")
	premisesGroup := v1.Group("/premises")
	camerasGroup := v1.Group("/cameras")

	health.GET("", func(c echo.Context) error {
		return c.JSON(http.StatusOK, map[string]string{"status": "OK"})
	})
	usersHandlers.RegisterRoutes(usersGroup)
	premisesHandlers.RegisterRoutes(premisesGroup)
	camerasHandlers.RegisterRoutes(camerasGroup)
	return nil

}
