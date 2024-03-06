package services

import "test-go/core/models"

type PowerService interface {
	CreatePower() error

	GetSumPower(sum models.PowerGetSumSrv) (result []models.PowerRepository, err error)
}
