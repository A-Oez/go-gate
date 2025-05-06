package adminauth

import (
	"go-gate/internal/service/admin_auth/entity"

	"golang.org/x/crypto/bcrypt"
)

type AdminAuthRepository interface {
	CreateSession(user entity.AdminUser) (entity.SessionCreationResp, error)
	GetUserByMail(email string) (entity.AdminUser, error)
	GetSession(id string) (entity.Session, error)
}

type AdminAuthService struct {
	repository AdminAuthRepository
}

func NewAdminAuthServiceService(repo AdminAuthRepository) *AdminAuthService {
	return &AdminAuthService{
		repository: repo,
	}
}

func (as AdminAuthService) Login(credentials entity.AdminCredentials) (entity.SessionCreationResp, error) {
	user, err := as.repository.GetUserByMail(credentials.Email)
	if err != nil {
		return entity.SessionCreationResp{}, err
	} 

	ok, err := AuthorizeUser(user, credentials)
	if !ok {
		return entity.SessionCreationResp{}, err
	}

	return as.repository.CreateSession(user)
}

func (as *AdminAuthService) GetSession(id string) (entity.Session, error) {
	return as.repository.GetSession(id)
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