package main

import (
	"webchannels/manager"
	"webchannels/protocols"

	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/middleware/cors"
)

func main() {
	core := manager.NewCore()

	app := fiber.New(fiber.Config{})

	app.Use(cors.New())

	protocols.UseLongpoll(app, core)
	protocols.UseWs(app, core)
	protocols.UseSSE(app, core)

	app.Post("/*", func(c fiber.Ctx) error {
		data := new(interface{})

		if err := c.Bind().Body(data); err != nil {
			return err
		}

		core.Send(c.OriginalURL(), data)

		return c.SendStatus(204)
	})

	app.Listen(":3240")
}
