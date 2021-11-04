package database

import "github.com/pkg/errors"

const (
	checkTables = "select count(*) from `sqlite_master`;"
	setUri      = "insert into `urls` (`data`, `updated_at`) values (?, ?);"
	getUri      = "select `data` from `urls` where `id` = ?;"
)

const sqlDefaul = `
		create table urls (
			id			INTEGER PRIMARY KEY AUTOINCREMENT UNIQUE,
			data		TEXT,
			updated_at	NUMERIC
		);
`

var (
	ErrRecordNotFound = errors.New("record is not found")
)
