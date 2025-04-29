package adminauth

import (
	"go-gate/internal/service/admin_auth/entity"
	"time"
)

type AdminAuthRepository interface {
	CreateSession(user entity.AdminUser) (time.Time, error)
}

type AdminAuthService struct {
	repository AdminAuthRepository
}

func NewAdminAuthServiceService(repo AdminAuthRepository) *AdminAuthService {
	return &AdminAuthService{
		repository: repo,
	}
}

func (as AdminAuthService) Login(credentials entity.AdminCredentials) (time.Time, error) {
	//TODO: check credentials with adminuser then create session and return expiresAt time
	return time.Time{}, nil
}