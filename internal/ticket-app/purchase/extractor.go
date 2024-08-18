package purchase

import (
	"encoding/json"
	"errors"
	"gowit-task/internal/ticket-app/base"
	"gowit-task/internal/ticket-app/dto"
	"gowit-task/internal/ticket-app/extracttypes"

	"github.com/gofiber/fiber/v2"
)

var (
	errEmptyCtx     = errors.New("got empty fiber-context")
	errEmptyReqBody = errors.New("got empty request-body")
)

type PurchaseMakeExtractor struct {
	extracttypes.PurchaseBody
}

func NewPurchaseMakeExtractor() *PurchaseMakeExtractor {
	return &PurchaseMakeExtractor{}
}

func (extr PurchaseMakeExtractor) Extract(ctx *fiber.Ctx) (base.PurchaseReq, error) {
	if ctx == nil {
		return nil, errEmptyCtx
	}

	b := ctx.Request().Body()
	if b == nil {
		return nil, errEmptyReqBody
	}

	req := dto.MakePurchaseReqDTO{}
	err := json.Unmarshal(b, &req)
	if err != nil {
		return nil, err
	}
	req.Id, _ = ctx.ParamsInt("id", 0)

	return convertToPurchaseReq(&req), nil
}
