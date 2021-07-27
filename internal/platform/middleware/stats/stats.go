package stats

import (
	"fmt"
	"github.com/labstack/echo/v4"
)

type (
	Stats struct {
		RequestCount uint64 `json:"requestCount"`
	}
)

func NewMiddleStats() *Stats {
	return &Stats{
		RequestCount: 0,
	}
}

//Based on echo example https://echo.labstack.com/cookbook/middleware/
func (s *Stats) MiddleStats(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		if err := next(c); err != nil {
			c.Error(err)
		}
		s.RequestCount++
		fmt.Printf("Number of requests: %d \n", s.RequestCount)
		return nil
	}
}
