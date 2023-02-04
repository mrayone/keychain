package configuration

import "github.com/gofiber/fiber/v2"

// @Summary Handles configuration keys
// @Description This endpoint perform to create a configration bundle based on redis key
// @Tags
// @Accept json
// @Produce json
// @Router /v1/configuration [post]
func Post(c *fiber.Ctx) error {

	return c.SendString("hello World")
}
