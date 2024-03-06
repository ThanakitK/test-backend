package services

import (
	"test-go/core/models"
	"test-go/core/repositories"
)

type powerSrv struct {
	powerRepo repositories.PowerRepository
}

func NewPowerService(powerRepo repositories.PowerRepository) powerSrv {
	return powerSrv{
		powerRepo: powerRepo,
	}
}

func (s powerSrv) CreatePower() error {
	err := s.powerRepo.CreatePower()
	if err != nil {
		return err
	}
	return nil
}

func (s powerSrv) GetSumPower(sum models.PowerGetSumSrv) (result []models.PowerRepository, err error) {
	res, err := s.powerRepo.GetPower()
	if err != nil {
		return result, err
	}
	for _, v := range res {
		subSum := v.ActivePower + v.PowerInput
		if subSum == sum.Sum {
			result = append(result, v)
		}
	}
	return result, err
}
