package service

import (
	"errors"
	"online_wallet_humo/internal/repository"
	"online_wallet_humo/pkg/models"
	"os"
	"path/filepath"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

type UserService struct {
	userRepo *repository.UserRepository
}

func NewUserService(s *repository.UserRepository) *UserService {
	return &UserService{s}
}

func (s *UserService) CreateUserByAdmin(u *models.User) (uint, error) {
	return s.userRepo.CreateUserByAdmin(u)
}

func (s *UserService) CreateUser(phoneNumber, password string) (uint, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return 0, err
	}

	return s.userRepo.CreateUser(phoneNumber, string(hashedPassword))
}

func (s *UserService) GenerateTokenForUser(phoneNumber, password string) (string, error) {
	user, err := s.userRepo.GetUserByPhoneNumber(phoneNumber)
	if err != nil {
		return "", err
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(password)); err != nil {
		return "", err
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &tokenClaims{
		jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(
				time.Duration(24 * time.Hour),
			)),
			IssuedAt: jwt.NewNumericDate(time.Now()),
		},
		user.ID,
	})

	t, err := token.SignedString(secretKey)
	return t, err
}

func (s *UserService) ParseTokenForUser(token string) (uint, error) {
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

func (s *UserService) GetUserByID(id uint) (*models.User, error) {
	return s.userRepo.GetUserByID(id)
}

func (s *UserService) UpdateUser(user, updateUser *models.User) error {
	return s.userRepo.UpdateUser(user, updateUser)
}

func (s *UserService) DeleteUser(u *models.User) error {
	return s.userRepo.DeleteUser(u)
}

func (s *UserService) UploadAvatar(userID uint, fileData []byte, filename string, avatarPath string) error {
	_, err := s.userRepo.GetUserByID(userID)
	if err != nil {
		return err
	}

	avaPath := filepath.Join(avatarPath, filename)
	if err := os.WriteFile(avaPath, fileData, 0666); err != nil {
		return err
	}

	if err := s.userRepo.UpdateAvatar(userID, filename); err != nil {
		os.Remove(avaPath)
		return err
	}

	return nil
}
