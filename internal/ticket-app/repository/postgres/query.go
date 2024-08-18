package postgres

// ticket
const (
	createTableQuery = `CREATE TABLE IF NOT EXISTS tickets (
		id serial primary key,
		name VARCHAR(30),
		description TEXT,
		allocation INTEGER NOT NULL,
		event_time TIMESTAMP DEFAULT CURRENT_TIMESTAMP);`

	getTicketQuery = `SELECT name, description, allocation
		FROM tickets
		WHERE id=$1;`

	createTicketQuery = `INSERT INTO tickets 
		(name, description, allocation) 
		VALUES ($1, $2, $3)
		RETURNING id;`
)

// purchase
const (
	createPurchaseTable = `CREATE TABLE IF NOT EXISTS purchases (
		id serial primary key,
		ticket_id INTEGER REFERENCES tickets(id),
		purchased_seats INTEGER NOT NULL,
		event_time TIMESTAMP DEFAULT CURRENT_TIMESTAMP);`

	makePurchaseQuery = `INSERT INTO purchases
		(ticket_id, purchased_seats)
		VALUES ($1, $2)`

	getAvailableSeatCount = `SELECT *
		FROM (
			SELECT allocation - COALESCE(SUM(purchased_seats), 0) AS available_seats
			FROM tickets
			LEFT JOIN purchases ON tickets.id = purchases.ticket_id
			WHERE tickets.id = $1
			GROUP BY tickets.allocation
		) as subquery
		WHERE available_seats >= $2`
)
