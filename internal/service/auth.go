package service

import (
	"errors"
	"online_wallet_humo/internal/repository"
	"online_wallet_humo/pkg/models"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

type tokenClaims struct {
	jwt.RegisteredClaims
	UserID uint
}

var secretKey = []byte("your_secret_key")

type AuthService struct {
	authRepo *repository.AuthPostgres
}

func NewAuthService(authRepo repository.AuthPostgres) *AuthService {
	return &AuthService{&authRepo}
}

func (s *AuthService) CreateUser(fullName, email, phoneNumber, password string) (uint, error) {
	var u models.User
	u.FullName = fullName
	u.Email = email
	u.PhoneNumber = phoneNumber
	u.IsAdmin = true

	pswd, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return 0, err
	}

	u.PasswordHash = string(pswd)
	return s.authRepo.CreateUser(u)
}

func (s *AuthService) GenerateToken(email, password string) (string, error) {
	u, err := s.authRepo.GetUser(email)
	if err != nil {
		return "", err
	}

	if err := bcrypt.CompareHashAndPassword([]byte(u.PasswordHash), []byte(password)); err != nil {
		return "", err
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &tokenClaims{
		jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(
				time.Duration(24 * time.Hour),
			)),
			IssuedAt: jwt.NewNumericDate(time.Now()),
		},
		u.ID,
	})

	t, err := token.SignedString(secretKey)
	return t, err
}

func (s *AuthService) ParseToken(token string) (uint, error) {
	tkn, err := jwt.ParseWithClaims(token, &tokenClaims{}, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return 0, errors.New("Invalid signing method")
		}
		return secretKey, nil
	})
	if err != nil {
		return 0, err
	}

	claims, ok := tkn.Claims.(*tokenClaims)
	if !ok {
		return 0, errors.New("Token claims are not type of `*tokenClaims`")
	}

	return claims.UserID, nil
}
