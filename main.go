package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/khairulharu/bookingapi/internal/api"
	"github.com/khairulharu/bookingapi/internal/component"
	"github.com/khairulharu/bookingapi/internal/config"
	"github.com/khairulharu/bookingapi/internal/repository"
	"github.com/khairulharu/bookingapi/internal/service"
)

func main() {
	cnf := config.Get()
	dbConnection := component.GetDatabase(cnf)
	userRepository := repository.NewUser(dbConnection)
	chairRepository := repository.NewChair(dbConnection)

	bookingService := service.NewBooking(chairRepository, userRepository)

	app := fiber.New()
	app.Use(logger.New())

	api.NewBooking(app, bookingService)
	app.Listen(cnf.SRV.Host + ":" + cnf.SRV.Port)
}
