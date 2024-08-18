package ticket_create

import (
	"encoding/json"
	"gowit-task/internal/ticket-app/dto"
	"gowit-task/internal/ticket-app/model"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

type respWriter struct{}

func NewRespWriter() *respWriter {
	return &respWriter{}
}

func (resp *respWriter) Write(ctx *fiber.Ctx, statusCode int, info *model.FiestaInfo) {
	ctx.Response().SetStatusCode(http.StatusOK)

	if info == nil {
		return
	}
	b := convertToResp(info)

	ctx.Response().SetBody(b)
}

func convertToResp(info *model.FiestaInfo) []byte {
	resp := dto.FrontRespDTO{
		Id:         info.Id,
		Name:       info.Name,
		Desc:       info.Desc,
		Allocation: info.Allocation,
	}

	b, _ := json.Marshal(resp)
	return b
}
