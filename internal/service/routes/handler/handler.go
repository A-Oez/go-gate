package handler

import (
	"database/sql"
	"encoding/json"
	"go-gate/internal/service/routes"
	"go-gate/internal/service/routes/entity"
	"go-gate/internal/service/routes/repo"
	"io"
	"net/http"
	"strconv"
)

func GetRequestMappings(db *sql.DB) http.Handler {
	service := routes.NewMappingService(repo.NewRouteRepository(db))

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		mappings, err := service.GetAll()

		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}		

		json.NewEncoder(w).Encode(&mappings)
	})
}

func GetRequestMappingByID(db *sql.DB) http.Handler {
	service := routes.NewMappingService(repo.NewRouteRepository(db))

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		idStr := r.PathValue("id")
		id, err := strconv.Atoi(idStr)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		
		mappings, err := service.GetRouteByID(id)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}		

		json.NewEncoder(w).Encode(&mappings)
	})
}

func AddRequest(db *sql.DB) http.Handler {
	var mapping entity.AddRoute
	service := routes.NewMappingService(repo.NewRouteRepository(db))

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		body, err := io.ReadAll(r.Body)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
				
		err = json.Unmarshal(body, &mapping)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		
		_, err = service.AddRoute(mapping)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
	})
}