package service

import (
	"time"

	"github.com/Cmdliner/streem/internal/model"
	"github.com/Cmdliner/streem/internal/repository"
	"github.com/Cmdliner/streem/internal/util"
)

type AuthService struct {
	userRepository *repository.UserRepository
	jwtSecret string
	jwtExpiration time.Duration
}

func NewAuthService(userRepository *repository.UserRepository, jwtSecret string, jwtExpiration time.Duration) *AuthService {
	return &AuthService{
		userRepository: userRepository,
		jwtSecret: jwtSecret,
		jwtExpiration: jwtExpiration,
	}
}

func (s *AuthService) Register(user *model.User) (interface{}, error) {

	hashedPwd, err := util.HashPassword(user.Password)
	if err != nil {
		return nil, err
	}
	user.Password = hashedPwd

	return s.userRepository.Create(user)

}

// func (s *AuthService) Login


func (s *AuthService) Login() {}