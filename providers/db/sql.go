package db

const (
	CGetUrl       = "select `data` from `urls` where `id` = ?;"
	CSetUrl       = "insert into `urls` (`data`, `datetime`) values (?, ?);"
	cSelectTables = "select count(*) from `sqlite_master`;"
)
