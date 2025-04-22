package routes

import (
	"go-gate/internal/service/routes/entity"
)


type MappingRepository interface {
	GetAll() ([]entity.Route, error)
	GetRouteByClient(method string, publicPath string) (entity.Route, error)
	GetRouteByID(id int) (entity.Route, error)
	AddRoute(entity entity.AddRoute) (bool, error)
}

type MappingService struct {
	MappingRepository MappingRepository
}

func NewMappingService(mappingRepository MappingRepository) *MappingService {
	return &MappingService{
		MappingRepository: mappingRepository,
	}
}

func (ms *MappingService) GetAll() ([]entity.Route, error) {
	mappings, err := ms.MappingRepository.GetAll()
	if err != nil {
		return nil, err
	}

	return mappings, nil
}

func (ms *MappingService) GetRouteByClient(method string, publicPath string) (entity.Route, error) {
	var mapping entity.Route
	
	mapping, err := ms.MappingRepository.GetRouteByClient(method, publicPath)
	if err != nil {
		return mapping, err
	}

	return mapping, nil
}

func (ms *MappingService) GetRouteByID(id int) (entity.Route, error) {
	var mapping entity.Route
	
	mapping, err := ms.MappingRepository.GetRouteByID(id)
	if err != nil {
		return mapping, err
	}

	return mapping, nil
}

func (ms *MappingService) AddRoute(entity entity.AddRoute) (bool, error) {
	return ms.MappingRepository.AddRoute(entity)
}