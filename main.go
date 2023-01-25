package main

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func main() {
	app := fiber.New()

	app.Use(cors.New())
	groupV1 := app.Group("/v1")

	groupV1.Get("health", func(c *fiber.Ctx) error {

		return c.SendString("OK")
	})

	groupUser := groupV1.Group("/user")

	groupUser.Use(func(c *fiber.Ctx) error {
		c.Set("My-Custom-Header", "CUSTOM_HEADER_VALUE")
		c.Next()
		return nil
	})

	groupUser.Get("/:name", func(c *fiber.Ctx) error {
		fmt.Fprintf(c, "hello %s", c.Params("name"))
		return nil
	})

	app.Listen(":3000")
}
