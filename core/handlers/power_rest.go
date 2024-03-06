package handlers

import (
	"test-go/core/models"
	"test-go/core/services"

	"github.com/gofiber/fiber/v2"
)

type PowerHandler struct {
	powerSrv services.PowerService
}

func NewPowerHandler(powerSrv services.PowerService) PowerHandler {
	return PowerHandler{
		powerSrv: powerSrv,
	}
}

func (h PowerHandler) CreatePower(c *fiber.Ctx) error {
	err := h.powerSrv.CreatePower()
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"code":    fiber.StatusBadRequest,
			"status":  false,
			"message": "Failed to create",
		})
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"code":    200,
		"status":  true,
		"message": "create success",
	})
}

func (h PowerHandler) GetSumPower(c *fiber.Ctx) error {
	body := &models.PowerGetSumHandler{}
	if err := c.BodyParser(&body); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"code":    fiber.StatusBadRequest,
			"status":  false,
			"message": "Failed to parse body create",
			"data":    "",
		})
	}

	payload := models.PowerGetSumSrv{
		Sum: body.Sum,
	}

	result, err := h.powerSrv.GetSumPower(payload)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"code":    fiber.StatusBadRequest,
			"status":  false,
			"message": "Failed to get sum",
		})
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"code":    200,
		"status":  true,
		"message": "get sum success",
		"data":    result,
	})
}
