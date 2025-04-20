package mapping

import (
	"database/sql"
	"encoding/json"
	repo "go-gate/internal/db/repo/mapping"
	service "go-gate/internal/service/mapping"
	"net/http"
)

func GetRequestMappings(db *sql.DB) http.Handler {
	service := service.NewMappingService(repo.NewMappingRepository(db))
	mappings, err := service.GetAllMappings()

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}		

		json.NewEncoder(w).Encode(&mappings)
	})
}