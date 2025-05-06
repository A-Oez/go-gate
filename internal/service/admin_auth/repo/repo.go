package repo

import (
	"database/sql"
	"go-gate/internal/service/admin_auth/entity"
	"time"

	"github.com/google/uuid"
)

type AdminAuthRepository struct {
	DB *sql.DB
}

func NewAdminAuthRepository(db *sql.DB) *AdminAuthRepository {
	return &AdminAuthRepository{
		DB: db,
	}
}

func (ar *AdminAuthRepository) CreateSession(user entity.AdminUser) (entity.SessionCreationResp, error){
	query := `
		INSERT INTO sessions (id, user_email, created_at, expires_at)
		VALUES ($1, $2, $3, $4);
	`

	session_id := uuid.New()
	now := time.Now().UTC()
	expires_at := now.Add(time.Hour)
	
	_, err := ar.DB.Exec(query,
		session_id,
		user.Email,
		now,
		expires_at,
	)

	if err != nil {
		return entity.SessionCreationResp{}, err
	}

	return entity.SessionCreationResp{session_id, expires_at}, nil
}

func (ar *AdminAuthRepository) GetUserByMail(email string) (entity.AdminUser, error){
	var entity entity.AdminUser

	query := `
		SELECT *
		FROM admin_users
		WHERE email = $1
	`
	
	row := ar.DB.QueryRow(query, email)
	err := row.Scan(
		&entity.ID,
		&entity.Email,
		&entity.Password,
		&entity.CreatedAt,
	)

	if err != nil {
		return entity, err
	}

	return entity, nil
}
