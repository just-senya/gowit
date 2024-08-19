package postgres

import "gowit-task/internal/ticket-app/model"

func (p PostgresDb) Write(fiestaInfo model.FiestaInfo) (int, error) {
	var id int
	err := p.db.QueryRow(
		createTicketQuery,
		fiestaInfo.Name,
		fiestaInfo.Desc,
		fiestaInfo.Allocation,
	).Scan(&id)

	if err != nil {
		return 0, err
	}

	return id, err
}

func (p PostgresDb) Read(id int) (model.FiestaInfo, error) {
	res := p.db.QueryRow(getTicketQuery, id)

	var info model.FiestaInfo
	info.Id = id
	err := res.Scan(
		&info.Name,
		&info.Desc,
		&info.Allocation,
	)

	return info, err
}
