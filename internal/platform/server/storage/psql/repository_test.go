package psql

import (
	mmdb "api-testing/internal"
	mmdbMock "api-testing/internal/internalmocks"
	"context"
	"errors"
	"fmt"
	. "github.com/golang/mock/gomock"
	"github.com/pashagolub/pgxmock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

func Test_MovieRepository_Save_Error(t *testing.T) {

	id, name, genre, duration := "10", "Foo Movie", "Western", "120"
	movie, err := mmdb.NewMovie(id, name, genre, duration)
	require.NoError(t, err)

	dm, err := pgxmock.NewConn()
	require.NoError(t, err)
	controller := NewController(t)
	repo := mmdbMock.NewMockMovieRepository(controller)

	dm.ExpectExec("INSERT INTO public.movies (id, movie_name, genre, duration) VALUES ($1, $2, $3, $4)").
		WithArgs(id, name, genre, duration).
		WillReturnError(errors.New("something-failed"))

	repo.EXPECT().Save(context.Background(), movie).Return(errors.New("something-failed"))

	err = repo.Save(context.Background(), movie)
	fmt.Println(dm.ExpectationsWereMet())
	assert.Error(t, dm.ExpectationsWereMet())
	assert.Error(t, err)

}
