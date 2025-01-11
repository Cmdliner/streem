package service

import (
	"time"

	"github.com/Cmdliner/streem/internal/config"
	"github.com/Cmdliner/streem/internal/model"
	"github.com/Cmdliner/streem/internal/repository"
	"github.com/Cmdliner/streem/internal/util"
	"github.com/golang-jwt/jwt"
)

type AuthService struct {
	userRepository *repository.UserRepository
	otpRepository *repository.OtpRepository
	jwtSecret string
	jwtExpiration time.Duration
}

type UserLogin struct {
	Email string `json:"email" bson:"email"`
	Password string `json:"password" bson:"password"`
}

func NewAuthService(cfg *config.Config, userRepository *repository.UserRepository, otpRepository * repository.OtpRepository) *AuthService {
	return &AuthService{
		userRepository: userRepository,
		otpRepository: otpRepository,
		jwtSecret: cfg.JWT.Secret,
		jwtExpiration: time.Duration(cfg.JWT.ExpirationHours),
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

	// Create auth token and return
	authToken := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": user.ID,
		"exp": time.Now().Add(s.jwtExpiration).Unix(),
	})
	
	return authToken.SignedString([]byte(s.jwtSecret))
}

func (s *AuthService) ForgotPassword(email string) (string, error) {
	user, err := s.userRepository.GetByEmail(email)
	if err != nil {
		return "", err
	}

	// Create a new otp
	otpCode, err := s.otpRepository.Create(user, "pwd_reset")
	if err != nil {
		return "", err
	}

	// !todo => Send email with otp code
	return otpCode, nil
}


func (s *AuthService) UpdatePassword(email string, code string, password string) (*model.User, error) {
	// Find user with that ID
	user, err := s.userRepository.GetByEmail(email)
	if err != nil {
		return nil, err
	}

	// Find otp that matches kind, code and user
	_, err = s.otpRepository.GetOne(user, code, "pwd_reset")
	if err != nil {
		return nil, err
	}

	// !todo => Delete otp

	// Hash new password input and update in db
	user.Password, err = util.HashPassword(password)
	if err != nil {
		return nil, err
	}
	
	user, err = s.userRepository.Update(user.ID.String(), user)
	if err != nil {
		return nil, err
	}

	return user, nil
}