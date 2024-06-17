package usecase

import (
	"base-api/core/app/entity"
	"base-api/core/app/repository"
)

func (useCase *baseUsecase) CreateBase(name string) error {
	base := &entity.Base{Name: name}

	return useCase.baseRepo.CreateBase(base)
}

func (useCase *baseUsecase) GetBaseById(id int64) (*entity.Base, error) {
	return useCase.baseRepo.GetBaseById(id)
}

func (useCase *baseUsecase) ListBases() ([]*entity.Base, error) {
	return useCase.baseRepo.ListBases()
}

func NewBaseUsecase(baseRepo repository.BaseRepository) BaseUsecase {
	return &baseUsecase{ baseRepo }
}

type baseUsecase struct { baseRepo repository.BaseRepository }

type BaseUsecase interface {
	CreateBase(name string) error
	GetBaseById(id int64) (*entity.Base, error)
	ListBases() ([]*entity.Base, error)
}
