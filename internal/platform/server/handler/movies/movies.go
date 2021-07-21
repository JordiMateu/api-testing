package movies

import (
	"api-testing/internal/platform/server/storage/psql"
	"github.com/labstack/echo/v4"
	"net/http"
)

type createRequest struct {
	ID string `json:"id"" form:"id" query:"name"`
	Name string `json:"name"" form:"name" query:"name"`
	Genre string `json:"genre"" form:"genre" query:"genre"`
	Duration string `json:"duration"" form:"duration" query:"duration"`
}

func CreateHandler() echo.HandlerFunc {
	return func(c echo.Context) error {
		u := new(createRequest)
		if err := c.Bind(u); err != nil {
			return err
		}

		//This function returns a Movie struct. This can be marshalled into a JSON
		//and be returned.
		psql.AddMovie(u.ID, u.Name, u.Genre, u.Duration)
		return c.JSON(http.StatusCreated, u)
	}
}