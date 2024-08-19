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
	res, err := p.db.Query(getTicketQuery, id)
	if err != nil {
		return model.FiestaInfo{}, err
	}
	var info model.FiestaInfo
	for res.Next() {
		info.Id = id
		res.Scan(
			&info.Name,
			&info.Desc,
			&info.Allocation,
		)
	}

	return info, nil
}
