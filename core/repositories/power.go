package repositories

import "test-go/core/models"

type PowerRepository interface {
	CreatePower() error

	GetPower() (result []models.PowerRepository, err error)
}
