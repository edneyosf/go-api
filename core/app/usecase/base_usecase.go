package usecase

import (
	"base-api/core/app/data/request"
	"base-api/core/app/entity"
	"base-api/core/app/repository"
)

func (useCase *baseUsecase) CreateBase(data request.BaseData) error {
	base := &entity.Base{Name: data.Name}

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
	CreateBase(data request.BaseData) error
	GetBaseById(id int64) (*entity.Base, error)
	ListBases() ([]*entity.Base, error)
}
