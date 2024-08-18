package ticket_create

import (
	"gowit-task/internal/ticket-app/base"
	"gowit-task/internal/ticket-app/dto"
	"gowit-task/internal/ticket-app/extracttypes"
)

func convertToFrontReq(req *dto.FiestaCreateReqDTO) base.FrontReq {
	var res CreateExtractor

	extracttypes.ExtractTicketBody(req, &res.TicketBody)

	return res
}
