package extracttypes

import dto "gowit-task/internal/ticket-app/dto"

type TicketBody struct {
	Name       string
	Desc       string
	Allocation uint32
}

func (t TicketBody) GetName() string {
	return t.Name
}

func (t TicketBody) GetDesc() string {
	return t.Desc
}

func (t TicketBody) GetAllocation() uint32 {
	return t.Allocation
}

func ExtractTicketBody(source *dto.FiestaCreateReqDTO, dest *TicketBody) {
	if source == nil || dest == nil {
		return
	}
	dest.Name = source.Name
	dest.Desc = source.Desc
	dest.Allocation = source.Allocation
}
