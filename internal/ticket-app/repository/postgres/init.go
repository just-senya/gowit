package postgres

import (
	"database/sql"
	"errors"
	"gowit-task/internal/ticket-app/model"

	_ "github.com/lib/pq"
)

var (
	errCreateTicketTable   = errors.New("fail to create table: tickets")
	errCreatePurchaseTable = errors.New("fail to create table: purchases")
)

type PostgresDb struct {
	db *sql.DB
}

func NewPostgresStore() (*PostgresDb, error) {
	connStr := "host=db user=dbusername dbname=gowit_db password=dbpassword sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}

	return &PostgresDb{
		db: db,
	}, nil
}

func (p PostgresDb) Init() error {
	_, err := p.db.Exec(createTableQuery)
	if err != nil {
		return errCreateTicketTable
	}

	_, err = p.db.Exec(createPurchaseTable)
	if err != nil {
		return errCreatePurchaseTable
	}

	return nil
}

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

func (p PostgresDb) MakePurchase(id, quantity int, userId string) error {
	tx, err := p.db.Begin()
	if err != nil {
		return err
	}

	var availableSeats int
	err = tx.QueryRow(getAvailableSeatCount, id, quantity).Scan(&availableSeats)
	if err != nil {
		tx.Rollback()
		return err
	}

	_, err = tx.Exec(makePurchaseQuery, id, quantity)
	if err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit()
}
