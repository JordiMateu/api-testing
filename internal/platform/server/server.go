package server

import (
	mmdb "api-testing/internal"
	"api-testing/internal/platform/server/handler/health"
	"api-testing/internal/platform/server/handler/movies"
	"github.com/labstack/echo/v4"
)

type Server struct {
	engine     *echo.Echo
	repository mmdb.MovieRepository
}

func New(r mmdb.MovieRepository) *echo.Echo {
	s := Server{
		repository: r,
		engine:     echo.New(),
	}

	s.registerRoutes()
	return s.engine
}

func (s *Server) registerRoutes() {
	s.engine.GET("/health", health.HealthHandler())
	s.engine.POST("/movies", movies.CreateHandler(s.repository))
}
