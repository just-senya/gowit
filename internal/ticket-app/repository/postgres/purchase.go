package postgres

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
