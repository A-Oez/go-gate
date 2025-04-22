package routes

import (
	"go-gate/internal/service/routes/entity"
)


type RoutesRepository interface {
	GetAll() ([]entity.Route, error)
	GetRouteByClient(method string, publicPath string) (entity.Route, error)
	GetRouteByID(id int) (entity.Route, error)
	AddRoute(entity entity.AddRoute) (bool, error)
}

type RoutesService struct {
	repository RoutesRepository
}

func NewRoutesService(repo RoutesRepository) *RoutesService {
	return &RoutesService{
		repository: repo,
	}
}

func (rs *RoutesService) GetAll() ([]entity.Route, error) {
	routes, err := rs.repository.GetAll()
	if err != nil {
		return nil, err
	}

	return routes, nil
}

func (rs *RoutesService) GetRouteByClient(method string, publicPath string) (entity.Route, error) {	
	route, err := rs.repository.GetRouteByClient(method, publicPath)
	if err != nil {
		return route, err
	}

	return route, nil
}

func (rs *RoutesService) GetRouteByID(id int) (entity.Route, error) {	
	route, err := rs.repository.GetRouteByID(id)
	if err != nil {
		return route, err
	}

	return route, nil
}

func (rs *RoutesService) AddRoute(entity entity.AddRoute) (bool, error) {
	return rs.repository.AddRoute(entity)
}