package usecase

import (
	"project/internal/pkg/models"
)

type featuredUsecase struct {
	featuredRepo models.FeaturedRepositoryI
}

func NewFeaturedUsecase(featured models.FeaturedRepositoryI) models.FeaturedUsecaseI {
	return &featuredUsecase{
		featuredRepo: featured,
	}
}

func (uc *featuredUsecase) GetAll(user_id int64) ([]*models.Featured, error) {
	gp, err := uc.featuredRepo.GetAll(user_id)
	if err != nil {
		return nil, err
	}
	return gp, nil
}

func (uc *featuredUsecase) Create(featured *models.Featured) (int, error) {
	id, err := uc.featuredRepo.Create(featured)
	if err != nil {
		return 0, err
	}
	return id, nil
}

func (uc *featuredUsecase) Delete(featured *models.Featured) error {
	err := uc.featuredRepo.Delete(featured)
	if err != nil {
		return err
	}
	return nil
}
