package usecase

import (
	"project/internal/pkg/models"
)

type qualResultUsecase struct {
	qualResultRepo models.QualResultRepositoryI
}

func NewQualResultUsecase(qualResultRepo models.QualResultRepositoryI) models.QualResultUsecaseI {
	return &qualResultUsecase{
		qualResultRepo: qualResultRepo,
	}
}

func (uc *qualResultUsecase) GetQualResultById(id int) (*models.QualResultView, error) {
	result, err := uc.qualResultRepo.GetQualResultById(id)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (uc *qualResultUsecase) GetQualResultsOfGP(gp_id int) ([]*models.QualResultView, error) {
	qual_results, err := uc.qualResultRepo.GetQualResultsOfGP(gp_id)
	if err != nil {
		return nil, err
	}
	return qual_results, nil
}

func (uc *qualResultUsecase) Create(driver *models.QualResult) (int, error) {
	id, err := uc.qualResultRepo.Create(driver)
	if err != nil {
		return 0, err
	}
	return id, nil
}

func (uc *qualResultUsecase) Update(id int, newResult *models.QualResult) error {
	newResult.ID = id
	err := uc.qualResultRepo.Update(newResult)
	if err != nil {
		return err
	}
	return nil
}

func (uc *qualResultUsecase) Delete(id int) error {
	err := uc.qualResultRepo.Delete(id)
	if err != nil {
		return err
	}
	return nil
}
