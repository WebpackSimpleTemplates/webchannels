package protocols

import (
	"bufio"
	"encoding/json"
	"fmt"
	"main/manager"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/valyala/fasthttp"
)

func UseSSE(app *fiber.App, core *manager.Core) {
	app.Get("/sse/*", func(ctx *fiber.Ctx) error {
		channel := strings.Replace(ctx.OriginalURL(), "/sse", "", 1)

		dataChan := core.Add(channel, 1)

		ctx.Set("Content-Type", "text/event-stream")
		ctx.Set("Cache-Control", "no-cache")
		ctx.Set("Connection", "keep-alive")
		ctx.Set("Transfer-Encoding", "chunked")

		ctx.Status(fiber.StatusOK).Context().SetBodyStreamWriter(fasthttp.StreamWriter(func(w *bufio.Writer) {
			for {
				strData, _ := json.Marshal(<-dataChan)
				fmt.Fprintf(w, "data: %s\n\n", strData)

				err := w.Flush()
				if err != nil {
					close(dataChan)
					break
				}
			}
		}))

		return nil
	})
}
