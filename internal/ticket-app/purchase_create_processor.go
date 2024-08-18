package ticketapp

import (
	"gowit-task/internal/ticket-app/base"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

type PurchaseMakeProcessor struct {
	extractor base.PurchaseExtractor
	purchaser base.RepositoryPurchaseMaker
	responser base.PurchaseResponseWriter
}

func NewPurchaseMakeProcessor(
	extractor base.PurchaseExtractor,
	purchaser base.RepositoryPurchaseMaker,
	responser base.PurchaseResponseWriter,
) *PurchaseMakeProcessor {
	if extractor == nil || purchaser == nil || responser == nil {
		return nil
	}
	return &PurchaseMakeProcessor{
		extractor: extractor,
		purchaser: purchaser,
		responser: responser,
	}
}

func (p PurchaseMakeProcessor) Handle(ctx *fiber.Ctx) error {
	extr, err := p.extractor.Extract(ctx)
	if err != nil {
		ctx.Context().SetStatusCode(http.StatusBadRequest)
		ctx.Context().SetBody([]byte(err.Error()))
		return nil
	}

	err = p.purchaser.MakePurchase(extr.GetId(), extr.GetQuantity(), extr.GetUserId())
	if err != nil {
		p.responser.Write(ctx, http.StatusBadRequest)
		ctx.Context().SetBody([]byte("fail to make purchase"))
		return nil
	}

	p.responser.Write(ctx, http.StatusOK)
	return nil
}
