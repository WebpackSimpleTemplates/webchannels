package protocols

import (
	"strings"
	"webchannels/manager"

	"github.com/gofiber/contrib/v3/websocket"
	"github.com/gofiber/fiber/v3"
)

func UseWs(app *fiber.App, core *manager.Core) {
	app.Use("/ws", func(c fiber.Ctx) error {
		if websocket.IsWebSocketUpgrade(c) {
			c.Locals("allowed", true)
			return c.Next()
		}
		return fiber.ErrUpgradeRequired
	})

	app.Get("/ws/*", func(ctx fiber.Ctx) error {
		channel := strings.Replace(ctx.OriginalURL(), "/ws", "", 1)

		dataChan := core.Add(channel, 1)

		return websocket.New(func(socket *websocket.Conn) {
			defer recover()

			go func() {
				for {
					if _, _, err := socket.ReadMessage(); err != nil {
						close(dataChan)
						return
					}
				}
			}()

			for {
				data := <-dataChan

				if err := socket.WriteJSON(data); err != nil {
					return
				}
			}
		})(ctx)
	})
}
