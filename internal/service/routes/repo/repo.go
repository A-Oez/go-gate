package repo

import (
	"database/sql"
	"go-gate/internal/service/routes/entity"
)

type RouteRepository struct {
	DB *sql.DB
}

func NewRouteRepository(db *sql.DB) *RouteRepository{
	return &RouteRepository{
		DB: db,
	}
}

func (r *RouteRepository) GetAll() ([]entity.Route, error) {
	rows, err := r.DB.Query("SELECT * FROM routes")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var routes []entity.Route
	for rows.Next() {
		var route entity.Route
		err := rows.Scan(&route.ID, &route.Method, &route.PublicPath, &route.ServiceScheme, &route.ServiceHost, &route.ServicePath)
		if err != nil {
			return nil, err
		}
		routes = append(routes, route)
	}

	return routes, nil
}

func (r *RouteRepository) GetRouteByClient(method string, publicPath string) (entity.Route, error) {
	var entity entity.Route

	query := `
		SELECT *
		FROM routes
		WHERE method = $1 AND public_path = $2
	`
	
	row := r.DB.QueryRow(query, method, publicPath)
	err := row.Scan(
		&entity.ID,
		&entity.Method,
		&entity.PublicPath,
		&entity.ServiceScheme,
		&entity.ServiceHost,
		&entity.ServicePath,
	)

	if err != nil {
		return entity, err
	}

	return entity, nil
}

func (r *RouteRepository) AddRoute(entity entity.AddRoute) (bool, error) {
	query := `
		INSERT INTO routes (method, public_path, service_scheme, service_host, service_path)
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

func (r *RouteRepository) GetRouteByID(id int) (entity.Route, error) {
	var entity entity.Route

	query := `
		SELECT *
		FROM routes
		WHERE id = $1
	`
	
	row := r.DB.QueryRow(query, id)
	err := row.Scan(
		&entity.ID,
		&entity.Method,
		&entity.PublicPath,
		&entity.ServiceScheme,
		&entity.ServiceHost,
		&entity.ServicePath,
	)

	if err != nil {
		return entity, err
	}

	return entity, nil
}

func (r *RouteRepository) DeleteRouteByID(id int) (bool, error) {
	query := `
		DELETE FROM routes WHERE id = $1
	`

	_, err := r.DB.Exec(query,
		id,
	)

	if err != nil {
		return false, err
	}

	return true, nil
}

func (r *RouteRepository) UpdateRoute(entity entity.Route) (bool, error) {
	query := `
		UPDATE routes 
		SET method = $1, 
		public_path = $2,
		service_scheme = $3, 
		service_host = $4, 
		service_path = $5
		WHERE id = $6
	`

	_, err := r.DB.Exec(query,
		entity.Method,
		entity.PublicPath,
		entity.ServiceScheme,
		entity.ServiceHost,
		entity.ServicePath,
		entity.ID,
	)

	if err != nil {
		return false, err
	}

	return true, nil
}