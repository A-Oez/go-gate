package handler

import (
	"go-gate/internal/service/routes"
	"go-gate/internal/service/routes/entity"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

type mockRoutesService struct{}

func (m *mockRoutesService) GetAll() ([]entity.Route, error) {
	return nil, routes.ErrDBQueryFailed
}

func (m *mockRoutesService) GetRouteByID(id int) (entity.Route, error) {
	return entity.Route{}, nil
}

func (m *mockRoutesService) AddRoute(e entity.AddRoute) (bool, error) {
	return true, nil
}

func (m *mockRoutesService) DeleteRouteByID(id int) (bool, error) {
	return true, nil
}

func TestGetAll_DBError(t *testing.T) {
	rh := &RoutesHandler{service: &mockRoutesService{}}

	r := httptest.NewRequest(http.MethodGet, "/api/routes", nil)
	rc := httptest.NewRecorder()
	rh.GetAll().ServeHTTP(rc, r)

	res := rc.Result()
	defer res.Body.Close()

	expected := `{
		"status": 500,
		"error_message": "database query failed"
	}`
	assert.Equal(t, http.StatusInternalServerError, res.StatusCode)
	assert.JSONEq(t, expected, rc.Body.String())
}
