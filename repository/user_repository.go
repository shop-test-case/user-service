package repository

import (
	"user-service/model"

	"gorm.io/gorm"
)

type IUserRepo interface {
	Create(user *model.User) error
	FindByIdentifier(id string) (*model.User, error)
}

type UserRepo struct {
	DB *gorm.DB
}

func (r *UserRepo) Create(user *model.User) error {
	return r.DB.Create(user).Error
}

func (r *UserRepo) FindByIdentifier(id string) (*model.User, error) {
	var user model.User

	err := r.DB.Where("email = ? OR phone = ?", id, id).First(&user).Error

	return &user, err
}
