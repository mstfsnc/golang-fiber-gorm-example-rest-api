package main

import (
	"fmt"
	"github.com/gofiber/adaptor/v2"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"sample-app/internal/handlers"
	"sample-app/internal/repositories"
	"sample-app/internal/services"
	"sample-app/pkg/config"
	"sample-app/pkg/metric"
)

func main() {
	dsn := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s connect_timeout=%s",
		config.Get("DB_HOST"),
		config.Get("DB_PORT"),
		config.Get("DB_USER"),
		config.Get("DB_PASS"),
		config.Get("DB_NAME"),
		config.Get("DB_TIMEOUT"),
	)

	dbConn, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	metric := metric.NewMetric(config.Get("APP_NAME"))

	userRepository := repositories.NewUserRepository(dbConn, metric)
	userService := services.NewUserService(userRepository)
	userHandler := handlers.NewUserHandler(userService)

	app := fiber.New()
	app.Use(logger.New())

	userHandler.SetRoute(app)

	metrics := adaptor.HTTPHandler(promhttp.Handler())
	app.Get("/metrics", metrics)

	log.Fatal(app.Listen(config.Get("APP_PORT")))
}
