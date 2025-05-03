package adminauth

import (
	"go-gate/internal/service/admin_auth/entity"
	"time"

	"golang.org/x/crypto/bcrypt"
)

type AdminAuthRepository interface {
	CreateSession(user entity.AdminUser) (time.Time, error)
	GetUserByMail(email string) (entity.AdminUser, error)
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
	user, err := as.repository.GetUserByMail(credentials.Email)
	if err != nil {
		return time.Time{}, err
	} 

	ok, err := AuthorizeUser(user, credentials)
	if !ok{
		return time.Time{}, err
	}

	//TODO: create session, return expriesAt time
	return time.Time{}, nil
}

func AuthorizeUser(user entity.AdminUser, credentials entity.AdminCredentials) (bool, error) {
	if credentials.Email != user.Email{
		return false, ErrInvalidUser
	}

	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(credentials.Password))
	if err != nil {
		return false, ErrPasswordDoesNotMatch
	}

	return true, nil
}