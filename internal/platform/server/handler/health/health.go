package health

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

func HealthHandler() echo.HandlerFunc {
	return func(c echo.Context) error {
		//Here we can add database status, reddis status, queues status...
		return c.String(http.StatusOK, "*** All systems running ***")
	}
}
