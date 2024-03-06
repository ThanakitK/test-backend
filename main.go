package main

import (
	"test-go/config"
	"test-go/core/handlers"
	"test-go/core/repositories"
	"test-go/core/services"

	"github.com/gofiber/fiber/v2"
	"github.com/spf13/viper"
)

func init() {
	config.NewAppInitEnvironment()

}

func main() {
	db := config.NewAppDatabase()

	// Repo
	powerRepo := repositories.NewPowerRepository(db, "power")

	// Service
	powerSrv := services.NewPowerService(powerRepo)

	// Handler
	powerHandler := handlers.NewPowerHandler(powerSrv)

	port := viper.GetString("app.port")
	app := fiber.New()

	app.Post("/power/create", powerHandler.CreatePower)

	app.Get("/power/sum", powerHandler.GetSumPower)

	app.Listen(":" + port)
}
