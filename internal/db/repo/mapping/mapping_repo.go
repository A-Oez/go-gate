package mapping

import (
	"database/sql"
	entity "go-gate/internal/db/entity/mapping"
)

type MappingRepository struct {
	DB *sql.DB
}

func NewMappingRepository(db *sql.DB) *MappingRepository{
	return &MappingRepository{
		DB: db,
	}
}

func (r *MappingRepository) GetAll() ([]entity.ProxyMapping, error) {
	rows, err := r.DB.Query("SELECT * FROM mappings")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var mappings []entity.ProxyMapping
	for rows.Next() {
		var mapping entity.ProxyMapping
		err := rows.Scan(&mapping.ID, &mapping.Method, &mapping.PublicPath, &mapping.ServiceScheme, &mapping.ServiceHost, &mapping.ServicePath)
		if err != nil {
			return nil, err
		}
		mappings = append(mappings, mapping)
	}

	return mappings, nil
}

func (r *MappingRepository) GetRequestByClient(method string, publicPath string) (entity.ProxyMapping, error) {
	var mapping entity.ProxyMapping

	query := `
		SELECT *
		FROM mappings
		WHERE method = $1 AND public_path = $2
	`
	
	row := r.DB.QueryRow(query, method, publicPath)
	err := row.Scan(
		&mapping.ID,
		&mapping.Method,
		&mapping.PublicPath,
		&mapping.ServiceScheme,
		&mapping.ServiceHost,
		&mapping.ServicePath,
	)

	if err != nil {
		return mapping, err
	}

	return mapping, nil
}

func (r *MappingRepository) AddRequest(entity entity.ProxyMappingAdd) (bool, error) {
	query := `
		INSERT INTO mappings (method, public_path, service_scheme, service_host, service_path)
		VALUES ($1, $2, $3, $4, $5);
	`
	
	_, err := r.DB.Exec(query,
		entity.Method,
		entity.PublicPath,
		entity.ServiceScheme,
		entity.ServiceHost,
		entity.ServicePath,
	)

	if err != nil {
		return false, err
	}

	return true, nil
}

func (r *MappingRepository) GetRequestMappingByID(id int) (entity.ProxyMapping, error) {
	var mapping entity.ProxyMapping

	query := `
		SELECT *
		FROM mappings
		WHERE id = $1
	`
	
	row := r.DB.QueryRow(query, id)
	err := row.Scan(
		&mapping.ID,
		&mapping.Method,
		&mapping.PublicPath,
		&mapping.ServiceScheme,
		&mapping.ServiceHost,
		&mapping.ServicePath,
	)

	if err != nil {
		return mapping, err
	}

	return mapping, nil
}