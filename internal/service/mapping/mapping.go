package mapping

import (
	entity "go-gate/internal/db/entity/mapping"
)


type MappingRepository interface {
	GetAll() ([]entity.ProxyMapping, error)
	GetRequestByClient(method string, publicPath string) (entity.ProxyMapping, error)
	GetRequestMappingByID(id int) (entity.ProxyMapping, error)
	AddRequest(entity entity.ProxyMappingAdd) (bool, error)
}

type MappingService struct {
	MappingRepository MappingRepository
}

func NewMappingService(mappingRepository MappingRepository) *MappingService {
	return &MappingService{
		MappingRepository: mappingRepository,
	}
}

func (ms *MappingService) GetAllMappings() ([]entity.ProxyMapping, error) {
	mappings, err := ms.MappingRepository.GetAll()
	if err != nil {
		return nil, err
	}

	return mappings, nil
}

func (ms *MappingService) GetRequestByClient(method string, publicPath string) (entity.ProxyMapping, error) {
	var mapping entity.ProxyMapping
	
	mapping, err := ms.MappingRepository.GetRequestByClient(method, publicPath)
	if err != nil {
		return mapping, err
	}

	return mapping, nil
}

func (ms *MappingService) GetRequestMappingByID(id int) (entity.ProxyMapping, error) {
	var mapping entity.ProxyMapping
	
	mapping, err := ms.MappingRepository.GetRequestMappingByID(id)
	if err != nil {
		return mapping, err
	}

	return mapping, nil
}

func (ms *MappingService) AddRequest(entity entity.ProxyMappingAdd) (bool, error) {
	return ms.MappingRepository.AddRequest(entity)
}