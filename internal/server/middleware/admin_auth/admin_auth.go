package adminauth

import (
	"database/sql"
	adminauth "go-gate/internal/service/admin_auth"
	"go-gate/internal/service/admin_auth/repo"
	"go-gate/pkg/httperror"
	"net/http"
)

func AuthAdmin(db *sql.DB, next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		session_id := ""
		for _, cookie := range r.Cookies() {
			if(cookie.Name == "session_id"){
				session_id = cookie.Value
			}
		}
		
		service := adminauth.NewAdminAuthService(repo.NewAdminAuthRepository(db))

		_, err := service.GetSession(session_id)
		if err != nil {
			httperror.DefaultError{
				Status: http.StatusUnauthorized,
				Msg: err.Error(),
			}.WriteError(w)
			return
		}

		next.ServeHTTP(w, r)
	})
}