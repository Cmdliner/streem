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

type UserLogin struct {
	Email string `json:"email" bson:"email"`
	Password string `json:"password" bson:"password"`
}

func NewAuthService(userRepository *repository.UserRepository, jwtSecret string, jwtExpiration time.Duration) *AuthService {
	return &AuthService{
		userRepository: userRepository,
		jwtSecret: jwtSecret,
		jwtExpiration: jwtExpiration,
	}
}

func (s *AuthService) Register(user *model.User) (*model.User, error) {

	hashedPwd, err := util.HashPassword(user.Password)
	if err != nil {
		return nil, err
	}
	user.Password = hashedPwd

	return s.userRepository.Create(user)
}


func (s *AuthService) Login(loginData *UserLogin) (string, error) { 
	user, err := s.userRepository.GetByEmail(loginData.Email)
	if err != nil {
		return "", err
	}

	pwdMatch := util.CheckPasswordHash(loginData.Password, user.Password)
	if !pwdMatch {
		return "", err
	}
	// !todo => Create auth token and return
	
	

	return "nlsnfvjlnerlejrleklkrjelkjr", nil
}