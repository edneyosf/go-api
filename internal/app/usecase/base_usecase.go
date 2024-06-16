package usecase

import (
	"base-api/internal/app/entity"
	"base-api/internal/app/repository"
)

type BaseUsecase interface {
	Create(name string) (*entity.Base, error)
	GetById(id int64) (*entity.Base, error)
	List() ([]*entity.Base, error)
}

type baseUsecase struct {
	baseRepo repository.BaseRepository
}

func NewUsecase(repo *repository.BaseRepository) BaseUsecase {
	return &baseUsecase{
		baseRepo: *repo,
	}
}

func (u *baseUsecase) Create(name string) (*entity.Base, error) {
	Baser := &entity.Base{Name: name}
	return u.baseRepo.Create(Baser)
}

func (u *baseUsecase) GetById(id int64) (*entity.Base, error) {
	return u.baseRepo.GetById(id)
}

func (u *baseUsecase) List() ([]*entity.Base, error) {
	return u.baseRepo.List()
}
