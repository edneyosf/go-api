package repository

import (
	"base-api/internal/app/entity"
	"gorm.io/gorm"
)

type BaseRepository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *BaseRepository {
	return &BaseRepository{db: db}
}

func (r *BaseRepository) Create(user *entity.Base) (*entity.Base, error) {
	if err := r.db.Create(user).Error; err != nil {
		return nil, err
	}
	return user, nil
}

func (r *BaseRepository) GetById(id int64) (*entity.Base, error) {
	var user entity.Base
	if err := r.db.First(&user, id).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *BaseRepository) List() ([]*entity.Base, error) {
	var users []*entity.Base
	if err := r.db.Find(&users).Error; err != nil {
		return nil, err
	}
	return users, nil
}
