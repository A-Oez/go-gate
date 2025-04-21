package mapping

import (
	"database/sql"
	"encoding/json"
	entity "go-gate/internal/db/entity/mapping"
	repo "go-gate/internal/db/repo/mapping"
	service "go-gate/internal/service/mapping"
	"io"
	"net/http"
	"strconv"
)

func GetRequestMappings(db *sql.DB) http.Handler {
	service := service.NewMappingService(repo.NewMappingRepository(db))

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		mappings, err := service.GetAllMappings()

		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}		

		json.NewEncoder(w).Encode(&mappings)
	})
}

func GetRequestMappingByID(db *sql.DB) http.Handler {
	service := service.NewMappingService(repo.NewMappingRepository(db))

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		idStr := r.PathValue("id")
		id, err := strconv.Atoi(idStr)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		
		mappings, err := service.GetRequestMappingByID(id)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}		

		json.NewEncoder(w).Encode(&mappings)
	})
}

func AddRequest(db *sql.DB) http.Handler {
	var mapping entity.ProxyMappingAdd
	service := service.NewMappingService(repo.NewMappingRepository(db))

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
		
		_, err = service.AddRequest(mapping)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
	})
}