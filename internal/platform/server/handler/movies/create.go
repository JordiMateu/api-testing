package movies

import (
	mmdb "api-testing/internal"
	"github.com/labstack/echo/v4"
	"net/http"
)

type createRequest struct {
	ID       string `json:"id" form:"id" query:"name"`
	Name     string `json:"name" form:"name" query:"name"`
	Genre    string `json:"genre" form:"genre" query:"genre"`
	Duration string `json:"duration" form:"duration" query:"duration"`
}

func CreateHandler(movieRepository mmdb.MovieRepository) echo.HandlerFunc {
	return func(c echo.Context) error {
		u := new(createRequest)
		if err := c.Bind(u); err != nil {
			return err
		}

		m, err := mmdb.NewMovie(u.ID, u.Name, u.Genre, u.Duration)

		if err != nil {
			return c.JSON(http.StatusInternalServerError, err.Error())
		}

		movieRepository.Save(c.Request().Context(), m)
		return c.JSON(http.StatusCreated, u)
	}
}
