package handler

import (
	"database/sql"
	"encoding/json"
	"fmt"
	adminauth "go-gate/internal/service/admin_auth"
	"go-gate/internal/service/admin_auth/entity"
	"go-gate/internal/service/admin_auth/repo"
	"go-gate/pkg/httperror"
	"io"
	"net/http"
	"time"
)

type AdminAuthHandler struct {
	service AdminAuthService
}

type AdminAuthService interface {
	Login(credentials entity.AdminCredentials) (time.Time, error)
}

func NewAdminAuthHandler(db *sql.DB) *AdminAuthHandler {
	return &AdminAuthHandler{
		service: adminauth.NewAdminAuthServiceService(repo.NewAdminAuthRepository(db)),
	}
}

func (ah *AdminAuthHandler) Login() http.Handler {
	var entity entity.AdminCredentials
	
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		body, err := io.ReadAll(r.Body)
		if err != nil {
			httperror.DefaultError{
				Status: http.StatusInternalServerError,
				Msg: adminauth.ErrReadBody.Error(),
			}.WriteError(w)
			return
		}

		err = json.Unmarshal(body, &entity)
		if err != nil {
			httperror.DefaultError{
				Status: http.StatusInternalServerError,
				Msg: adminauth.ErrUnmarshalJSON.Error(),
			}.WriteError(w)
			return
		}
		
		time, err := ah.service.Login(entity)
		if err != nil {
			httperror.DefaultError{
				Status: http.StatusInternalServerError,
				Msg:    err.Error(),
			}.WriteError(w)
			return
		}

		fmt.Fprintf(w, "%s", time.String())
	})
}