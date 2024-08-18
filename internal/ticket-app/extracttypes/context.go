package extracttypes

import (
	"github.com/gofiber/fiber/v2"
)

type TicketContext struct {
	Id int
}

func (t *TicketContext) ExtractCtx(ctx *fiber.Ctx) {
	t.Id, _ = ctx.ParamsInt("id", 0)
}

func (t TicketContext) GetId() int {
	return t.Id
}
