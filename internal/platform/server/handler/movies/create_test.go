package movies

import (
	mmdb "api-testing/internal"
	mmdbMock "api-testing/internal/internalmocks"
	"context"
	. "github.com/golang/mock/gomock"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

var userJSON = `{"id":"11","name":"Return to Houston","genre":"Western","duration":"100"}`
var userJSONNoName = `{"id":"11","name":"","genre":"Western","duration":"100"}`

func Test_CreateHandler(t *testing.T) {
	controller := NewController(t)
	req := httptest.NewRequest(http.MethodPost, "/movies", strings.NewReader(userJSON))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()

	mr := mmdbMock.NewMockMovieRepository(controller)
	e := echo.New()
	rf := CreateHandler(mr)
	c := e.NewContext(req, rec)

	movie, _ := mmdb.NewMovie("11", "Return to Houston", "Western", "100")
	mr.EXPECT().Save(context.Background(), movie).Return(nil)

	if assert.NoError(t, rf(c)) {
		assert.Equal(t, userJSON, rec.Body.String()[:len(rec.Body.String())-1])
	}
}

func Test_CreateHandler_Error_EmptyString(t *testing.T) {
	controller := NewController(t)
	req := httptest.NewRequest(http.MethodPost, "/movies", strings.NewReader(userJSONNoName))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()

	mr := mmdbMock.NewMockMovieRepository(controller)
	e := echo.New()
	rf := CreateHandler(mr)
	c := e.NewContext(req, rec)

	//movie,_ := mmdb.NewMovie("11", "Return to Houston", "Western", "100")
	//mr.EXPECT().Save(context.Background(), movie).Return(nil)

	if assert.NoError(t, rf(c)) {
		assert.True(t, strings.Contains(rec.Body.String(), "the field Name can not be empty"))
	}
}
