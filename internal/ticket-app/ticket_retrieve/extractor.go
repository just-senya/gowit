package ticket_retrieve

import (
	"gowit-task/internal/ticket-app/base"
	"gowit-task/internal/ticket-app/extracttypes"

	"github.com/gofiber/fiber/v2"
)

type TicketRetrieveExtractor struct {
	extracttypes.TicketContext
}

func NewTicketRetrieveExtractor() *TicketRetrieveExtractor {
	return &TicketRetrieveExtractor{}
}

func (extr TicketRetrieveExtractor) Extract(ctx *fiber.Ctx) (base.FiestaRetrieveReq, error) {
	extr.TicketContext.ExtractCtx(ctx)
	return extr, nil
}
