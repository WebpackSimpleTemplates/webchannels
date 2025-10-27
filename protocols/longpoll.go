package protocols

import (
	"main/manager"
	"strings"

	"github.com/gofiber/fiber/v2"
)

func UseLongpoll(app *fiber.App, core *manager.Core) {
	app.Get("/longpoll/*", func(ctx *fiber.Ctx) error {
		channel := strings.Replace(ctx.OriginalURL(), "/longpoll", "", 1)

		dataChan := core.Add(channel, 1)

		data := <-dataChan

		close(dataChan)

		return ctx.JSON(data)
	})
}
