package repository

import (
	"authentication/internal/models"
	"gorm.io/gorm"
)

type UserRepository struct {
	DB *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{DB: db}
}

func (r UserRepository) RegisterUser(user **models.Users) error {
	err := r.DB.Create(user).Error
	if err != nil {
		return err
	}
	return nil
}

func (r UserRepository) GetUserByEmail(email string) (*models.Users, error) {
	var user models.Users
	err := r.DB.Where("email = ?", email).First(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (r UserRepository) GetUserByID(id uint) (*models.Users, error) {
	var user models.Users
	err := r.DB.Where("user_id = ?", id).First(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (r UserRepository) UpdateUser(user **models.Users) error {
	err := r.DB.Save(user).Error
	if err != nil {
		return err
	}
	return nil
}

func (r UserRepository) DeleteUser(user **models.Users) error {
	err := r.DB.Delete(user).Error
	if err != nil {
		return err
	}
	return nil
}

func (r UserRepository) GetAllUsers() (*[]models.Users, error) {
	var users []models.Users
	err := r.DB.Find(&users).Error
	if err != nil {
		return nil, err
	}
	return &users, nil
}

func (r UserRepository) GetUsers() (*[]models.Users, error) {
	var users []models.Users
	err := r.DB.Where("deleted_at IS NOT NULL").Find(&users).Error
	if err != nil {
		return nil, err
	}
	return &users, nil
}

func (r UserRepository) GetUserByRole(role uint) (*[]models.Users, error) {
	var users []models.Users
	err := r.DB.Where("role_id = ?", role).Find(&users).Error
	if err != nil {
		return nil, err
	}
	return &users, nil
}

func (r UserRepository) GetUserByClientID(clientID string) (*models.Users, error) {
	var users models.Users
	err := r.DB.Where("client_id = ?", clientID).Find(&users).Error
	if err != nil {
		return nil, err
	}
	return &users, nil
}

func (r UserRepository) GetUserByClientAndRole(clientID, roleID uint) (*[]models.Users, error) {
	var users []models.Users
	err := r.DB.Where("client_id = ? AND role_id = ?", clientID, roleID).Find(&users).Error
	if err != nil {
		return nil, err
	}
	return &users, nil
}

func (r UserRepository) DeleteUserByID(id uint) error {
	err := r.DB.Where("user_id = ?", id).Delete(&models.Users{}).Error
	if err != nil {
		return err
	}
	return nil
}