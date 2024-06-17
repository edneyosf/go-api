package repository

import (
	"base-api/core/app/entity"
	"gorm.io/gorm"
)

func (repository *baseRepository) CreateBase(base *entity.Base) error {
	result := repository.db.Create(base)

	return result.Error
}

func (repository *baseRepository) GetBaseById(id int64) (*entity.Base, error) {
	var base *entity.Base
	result := repository.db.First(&base, id)

	return base, result.Error
}

func (repository *baseRepository) ListBases() ([]*entity.Base, error) {
	var bases []*entity.Base
	result := repository.db.Find(&bases)

	return bases, result.Error
}

func NewBaseRepository(db *gorm.DB) BaseRepository { return &baseRepository{ db } }

type baseRepository struct { db *gorm.DB }

type BaseRepository interface {
	CreateBase(base *entity.Base) error
	GetBaseById(id int64) (*entity.Base, error)
	ListBases() ([]*entity.Base, error)
}


