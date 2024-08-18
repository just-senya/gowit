package purchase

import "github.com/gofiber/fiber/v2"

type respWriter struct{}

func NewRespWriter() *respWriter {
	return &respWriter{}
}

func (resp *respWriter) Write(ctx *fiber.Ctx, statusCode int) {
	ctx.Context().SetStatusCode(statusCode)
}
