package configure

import (
	"errors"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/monitor"
)

func NewAPI() *fiber.App {
	api := fiber.New(fiber.Config{
		StrictRouting:         false,
		ErrorHandler:          errorHandler,
		DisableStartupMessage: true,
		AppName:               "emgo",
	})

	api.Use(logger.New(logger.Config{
		Format:     "${time} ${status} - ${method} ${path}\n",
		TimeFormat: "02-Jan-2006",
		TimeZone:   "America/New_York",
	}))

	api.Get("/metrics", monitor.New(monitor.Config{Title: "MyService Metrics Page"}))

	return api
}

func errorHandler(ctx *fiber.Ctx, err error) error {
	code := fiber.StatusInternalServerError

	var e *fiber.Error
	if errors.As(err, &e) {
		code = e.Code
	}

	err = ctx.Status(code).JSON(map[string]interface{}{"error": e.Message})
	if err != nil {
		return ctx.SendStatus(fiber.StatusInternalServerError)
	}

	return nil
}
