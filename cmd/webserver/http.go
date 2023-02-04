package webserver

import (
	"fmt"

	"github.com/gofiber/fiber/v2"

	"github.com/gofiber/swagger"
	"github.com/mrayone/keychain/api/configuration"
	"github.com/mrayone/keychain/configs"
	_ "github.com/mrayone/keychain/docs/swagger"
)

func Serve() error {
	app := fiber.New(fiber.Config{
		ReadTimeout:  configs.GetDuration("HTTP_SERVER_READ_TIMEOUT_MILLIS"),
		WriteTimeout: configs.GetDuration("HTTP_SERVER_WRITE_TIMEOUT_MILLIS"),
	})

	// health
	app.Get("/", func(c *fiber.Ctx) error {

		return c.JSON(struct {
			Status string `json:"status"`
		}{
			Status: "ok",
		})
	})

	app.Get("/swagger/*", swagger.HandlerDefault) // default

	// configuration
	app.Post("/v1/configuration", configuration.Post)

	return app.Listen(fmt.Sprintf(":%s", configs.GetString("HTTP_PORT")))
}
