package ticketapp

import (
	"gowit-task/internal/ticket-app/base"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

type FiestaGetProcessor struct {
	extractor     base.RequestRetrieveExtractor
	dataRetriever base.RepositoryReader
	response      base.ResponseWriter
}

func NewFiestaGetProcessor(
	extractor base.RequestRetrieveExtractor,
	dataRetriever base.RepositoryReader,
	response base.ResponseWriter,
) *FiestaGetProcessor {
	if extractor == nil || dataRetriever == nil || response == nil {
		return nil
	}
	return &FiestaGetProcessor{
		extractor:     extractor,
		dataRetriever: dataRetriever,
		response:      response,
	}
}

func (f FiestaGetProcessor) Handle(ctx *fiber.Ctx) error {
	extr, err := f.extractor.Extract(ctx)
	if err != nil {
		return err
	}
	info, err := f.dataRetriever.Read(extr.GetId())
	if err != nil {
		ctx.Context().SetStatusCode(http.StatusBadRequest)
		ctx.Context().SetBody([]byte("no ticket for given id"))
		return nil
	}

	f.response.Write(ctx, http.StatusOK, &info)

	return nil
}
