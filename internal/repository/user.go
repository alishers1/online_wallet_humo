package repository

import (
	"online_wallet_humo/pkg/models"

	"gorm.io/gorm"
)

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{db}
}

func (r *UserRepository) CreateUserByAdmin(u *models.User) (uint, error) {
	if err := r.db.Create(&u).Error; err != nil {
		return 0, err
	}

	return u.ID, nil
}

func (r *UserRepository) CreateUser(phoneNumber, password string) (uint, error) {
	var user models.User

	user.PhoneNumber = phoneNumber
	user.PasswordHash = password

	if err := r.db.Create(&user).Error; err != nil {
		return 0, err
	}

	return user.ID, nil
}

func (r *UserRepository) GetUserByID(id uint) (*models.User, error) {
	var user models.User

	if err := r.db.Where("id = ?", id).First(&user).Error; err != nil {
		return nil, err
	}

	return &user, nil
}

func (r *UserRepository) GetUserByPhoneNumber(phoneNumber string) (*models.User, error) {
	var user models.User

	if err := r.db.Where("phone_number = ?", phoneNumber).First(&user).Error; err != nil {
		return nil, err
	}

	return &user, nil
}

func (r *UserRepository) UpdateUser(user, updateUser *models.User) error {
	if err := r.db.Model(&user).Updates(updateUser).Error; err != nil {
		return err
	}

	return nil
}

func (r *UserRepository) DeleteUser(u *models.User) error {
	var user models.User
	if err := r.db.Where("id = ?", u.ID).First(&user).Error; err != nil {
		return err
	}

	user.IsActive = false

	if err := r.db.Where("id = ?", user.ID).Delete(&user).Error; err != nil {
		return err
	}

	return nil
}

func (r *UserRepository) UpdateAvatar(userID uint, avatarFilename string) error {
	if err := r.db.Model(&models.User{}).Where("id = ?", userID).Update("avatar", avatarFilename).Error; err != nil {
		return err
	}

	return nil
}
