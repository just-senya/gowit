package ticket_create

import (
	"encoding/json"
	"errors"
	"gowit-task/internal/ticket-app/base"
	dto "gowit-task/internal/ticket-app/dto"
	"gowit-task/internal/ticket-app/extracttypes"

	"github.com/gofiber/fiber/v2"
)

var (
	errEmptyCtx     = errors.New("got empty fiber-context")
	errEmptyReqBody = errors.New("got empty request-body")
)

type CreateExtractor struct {
	extracttypes.TicketBody
	extracttypes.TicketContext
}

func NewCreateExtractor() *CreateExtractor {
	return &CreateExtractor{}
}

func (extr CreateExtractor) Extract(ctx *fiber.Ctx) (base.FrontReq, error) {
	if ctx == nil {
		return nil, errEmptyCtx
	}

	b := ctx.Request().Body()
	if b == nil {
		return nil, errEmptyReqBody
	}
	var req dto.FiestaCreateReqDTO
	err := json.Unmarshal(b, &req)
	if err != nil {
		return nil, err
	}

	return convertToFrontReq(&req), nil
}
