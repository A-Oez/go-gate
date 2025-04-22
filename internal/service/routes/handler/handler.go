package handler

import (
	"database/sql"
	"encoding/json"
	"go-gate/internal/service/routes"
	"go-gate/internal/service/routes/entity"
	"go-gate/internal/service/routes/repo"
	"go-gate/pkg/httperror"
	"io"
	"net/http"
	"strconv"
)

type RoutesHandler struct{
	service RoutesService
}

type RoutesService interface {
	GetAll() ([]entity.Route, error)
	GetRouteByID(id int) (entity.Route, error)
	AddRoute(entity entity.AddRoute) (bool, error)
}

func NewRoutesHandler(db *sql.DB) *RoutesHandler {
	return &RoutesHandler{
		service: routes.NewRoutesService(repo.NewRouteRepository(db)),
	}
}

func (rh *RoutesHandler) GetAll() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		entity, err := rh.service.GetAll()

		if err != nil {
			httperror.DefaultError{
				Status: http.StatusInternalServerError,
				Err: routes.ErrDBQueryFailed,
			}.WriteError(w)
			return
		}		

		json.NewEncoder(w).Encode(&entity)
	})
}

func (rh *RoutesHandler) GetRouteByID() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		idStr := r.PathValue("id")
		id, err := strconv.Atoi(idStr)
		if err != nil {
			httperror.DefaultError{
				Status: http.StatusInternalServerError,
				Err: routes.ErrInvalidID,
			}.WriteError(w)
			return
		}
		
		entity, err := rh.service.GetRouteByID(id)
		if err != nil {
			httperror.DefaultError{
				Status: http.StatusInternalServerError,
				Err: routes.ErrDBQueryFailed,
			}.WriteError(w)
			return
		}		

		json.NewEncoder(w).Encode(&entity)
	})
}

func (rh *RoutesHandler) AddRoute() http.Handler {
	var entity entity.AddRoute

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		body, err := io.ReadAll(r.Body)
		if err != nil {
			httperror.DefaultError{
				Status: http.StatusInternalServerError,
				Err: routes.ErrReadBody,
			}.WriteError(w)
			return
		}
				
		err = json.Unmarshal(body, &entity)
		if err != nil {
			httperror.DefaultError{
				Status: http.StatusInternalServerError,
				Err: routes.ErrUnmarshalJSON,
			}.WriteError(w)
			return
		}
		
		_, err = rh.service.AddRoute(entity)
		if err != nil {
			httperror.DefaultError{
				Status: http.StatusInternalServerError,
				Err: routes.ErrDBQueryFailed,
			}.WriteError(w)
			return
		}
	})
}