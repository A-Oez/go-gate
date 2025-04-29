package repo

import (
	"database/sql"
	"go-gate/internal/service/admin_auth/entity"
	"time"
)

type AdminAuthRepository struct {
	DB *sql.DB
}

func NewAdminAuthRepository(db *sql.DB) *AdminAuthRepository {
	return &AdminAuthRepository{
		DB: db,
	}
}

func (ar *AdminAuthRepository) CreateSession(user entity.AdminUser) (time.Time, error){
	query := `
		INSERT INTO sessions (user_email, created_at, expires_at)
		VALUES ($1, $2, $3);
	`

	now := time.Now().UTC()
	expires_at := now.Add(time.Hour)
	
	_, err := ar.DB.Exec(query,
		user.Email,
		now,
		expires_at,
	)

	if err != nil {
		return time.Time{}, err
	}

	return expires_at, nil
}
