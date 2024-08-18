package ticket_create

import (
	"gowit-task/internal/ticket-app/base"
	"gowit-task/internal/ticket-app/model"
)

type fiestaEventCreator struct{}

func NewFiestaEventCreator() fiestaEventCreator {
	return fiestaEventCreator{}
}

func (f fiestaEventCreator) CreateFiesta(req base.FrontReq) (model.FiestaInfo, error) {
	fiesta := model.FiestaInfo{
		Name:       req.GetName(),
		Desc:       req.GetDesc(),
		Allocation: req.GetAllocation(),
	}
	return fiesta, nil
}
