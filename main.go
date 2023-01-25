package main

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	groupV1 := app.Group("/v1")
	groupUser := groupV1.Group("/user")

	groupUser.Get("/:name", func(c *fiber.Ctx) error {
		fmt.Fprintf(c, "hello %s", c.Params("name"))
		return nil
	})

	app.Listen(":3000")
}
