package service

import (
	"io/ioutil"
	"online_wallet_humo/internal/repository"
	"online_wallet_humo/pkg/models"
	"os"
	"path/filepath"
)

type UserService struct {
	userRepo *repository.UserRepository
}

func NewUserService(s *repository.UserRepository) *UserService {
	return &UserService{s}
}

func (s *UserService) CreateUser(u *models.User) (uint, error) {
	return s.userRepo.CreateUser(u)
}

func (s *UserService) GetUserByID(id uint) (*models.User, error) {
	return s.userRepo.GetUserByID(id)
}

func (s *UserService) UpdateUser(u *models.User) error {
	return s.userRepo.UpdateUser(u)
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
	if err := ioutil.WriteFile(avaPath, fileData, 0666); err != nil {
		return err
	}

	if err := s.userRepo.UpdateAvatar(userID, filename); err != nil {
		os.Remove(avaPath)
		return err
	}

	return nil
}
