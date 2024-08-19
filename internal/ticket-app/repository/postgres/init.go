package postgres

import (
	"database/sql"
	"errors"
	"fmt"

	_ "github.com/lib/pq"
)

var (
	errCreateTicketTable   = errors.New("fail to create table: tickets")
	errCreatePurchaseTable = errors.New("fail to create table: purchases")
)

type PostgresDb struct {
	db *sql.DB
}

func NewPostgresStore(host, user, dbname, password string) (*PostgresDb, error) {
	connStr := fmt.Sprintf("host=%s user=%s dbname=%s password=%s sslmode=disable",
		host,
		user,
		dbname,
		password,
	)

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
