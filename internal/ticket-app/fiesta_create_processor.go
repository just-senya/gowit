package ticketapp

import (
	"gowit-task/internal/ticket-app/base"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

type fiestaCreateProcessor struct {
	extractor     base.RequestExtractor
	fiestaCreator base.FiestaCreator
	response      base.ResponseWriter
	dbSender      base.RepositoryWriter
}

func NewFiestaCreateProcessor(
	extractor base.RequestExtractor,
	fiestaCreator base.FiestaCreator,
	response base.ResponseWriter,
	dbSender base.RepositoryWriter,
) *fiestaCreateProcessor {
	if extractor == nil || fiestaCreator == nil || response == nil || dbSender == nil {
		return nil
	}

	return &fiestaCreateProcessor{
		extractor:     extractor,
		fiestaCreator: fiestaCreator,
		response:      response,
		dbSender:      dbSender,
	}
}

func (f fiestaCreateProcessor) Handle(ctx *fiber.Ctx) error {
	tickReq, err := f.extractor.Extract(ctx)
	if err != nil {
		ctx.Context().SetStatusCode(http.StatusBadRequest)
		ctx.Context().SetBody([]byte(err.Error()))
		return nil
	}

	fiestaInfo, err := f.fiestaCreator.CreateFiesta(tickReq)
	if err != nil {
		f.response.Write(ctx, http.StatusInternalServerError, nil)
		return err
	}

	fiestaInfo.Id, err = f.dbSender.Write(fiestaInfo)
	if err != nil {
		f.response.Write(ctx, http.StatusInternalServerError, nil)
		return err
	}

	f.response.Write(ctx, http.StatusOK, &fiestaInfo)
	return nil
}
