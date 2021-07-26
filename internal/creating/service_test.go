package creating

import (
	mmdbMock "api-testing/internal/internalmocks"
	"context"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"reflect"
	"testing"
)

func Test_Returns_Expected_Struct(t *testing.T) {
	cnt := gomock.NewController(t)
	repoMock := mmdbMock.NewMockMovieRepository(cnt)
	s := NewMovieService(repoMock)
	assert.Equal(t, reflect.TypeOf(s).String(), "creating.MovieService")

}

func Test_Movie_DataValidation_NoError(t *testing.T) {
	cnt := gomock.NewController(t)
	repoMock := mmdbMock.NewMockMovieRepository(cnt)
	s := NewMovieService(repoMock)
	repoMock.EXPECT().Save(gomock.Any(), gomock.Any()).Return(nil)
	assert.NoError(t, s.CreateMovie(context.Background(), "1", "Gone with the Wind", "SciFi", "100"))
}

func Test_Movie_DataValidation_Errors(t *testing.T) {
	cnt := gomock.NewController(t)
	repoMock := mmdbMock.NewMockMovieRepository(cnt)
	s := NewMovieService(repoMock)
	assert.Error(t, s.CreateMovie(context.Background(), "", "Popi", "SciFi", "100"), "the field Id can not be empty")
	assert.Error(t, s.CreateMovie(context.Background(), "20", "", "SciFi", "100"), "the field Name can not be empty")
	assert.Error(t, s.CreateMovie(context.Background(), "20", "Xufi", "", "100"), "the field Genre can not be empty")
	assert.Error(t, s.CreateMovie(context.Background(), "20", "Xufi", "Tropex", ""), "the field Duration can not be empty")
}
