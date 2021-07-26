package server

import (
	"api-testing/internal/creating"
	"api-testing/internal/platform/server/handler/health"
	"api-testing/internal/platform/server/handler/movies"
	"github.com/labstack/echo/v4"
)

type Server struct {
	engine       *echo.Echo
	movieService creating.MovieService
}

func New(movieService creating.MovieService) *echo.Echo {
	s := Server{
		engine:       echo.New(),
		movieService: movieService,
	}

	s.registerRoutes()
	return s.engine
}

func (s *Server) registerRoutes() {
	s.engine.GET("/health", health.HealthHandler())
	s.engine.POST("/movies", movies.CreateHandler(s.movieService))
}
