package server

import (
	"api-testing/internal/platform/server/handler/health"
	"api-testing/internal/platform/server/handler/movies"
	"github.com/labstack/echo/v4"
)


type Server struct {
	httpAddr string
	engine *echo.Echo
}

func New(address string) *echo.Echo{
		s := Server{
			httpAddr: address,
			engine: echo.New(),
		}

		s.registerRoutes()
		return s.engine
}

func (s *Server) registerRoutes() {
	s.engine.GET("/health", health.HealthHandler())
	s.engine.POST("/movies", movies.CreateHandler())
}