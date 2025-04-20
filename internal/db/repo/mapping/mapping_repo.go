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