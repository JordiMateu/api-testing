package movies

import (
	mmdb "api-testing/internal"
	"api-testing/internal/creating"
	"errors"
	"github.com/labstack/echo/v4"
	"net/http"
)

type createRequest struct {
	ID       string `json:"id" form:"id" query:"name"`
	Name     string `json:"name" form:"name" query:"name"`
	Genre    string `json:"genre" form:"genre" query:"genre"`
	Duration string `json:"duration" form:"duration" query:"duration"`
}

func CreateHandler(s creating.MovieService) echo.HandlerFunc {
	return func(c echo.Context) error {
		u := new(createRequest)
		if err := c.Bind(u); err != nil {
			return err
		}

		err := s.CreateMovie(c.Request().Context(), u.ID, u.Name, u.Genre, u.Duration)

		if err != nil {
			switch {
			case errors.Is(err, mmdb.ErrEmptyID),
				errors.Is(err, mmdb.ErrEmptyName), errors.Is(err, mmdb.ErrEmptyGenre):
				return c.JSON(http.StatusBadRequest, err.Error())
			default:
				return c.JSON(http.StatusInternalServerError, err.Error())
			}
		}

		return c.JSON(http.StatusCreated, u)
	}
}
