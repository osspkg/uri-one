package db

import (
	"database/sql"
	"io/ioutil"

	"uri-one/app/config"

	_ "github.com/mattn/go-sqlite3"
	"github.com/sirupsen/logrus"
)

type SQLite struct {
	db  *sql.DB
	sql string
}

func MustNew(cfg *config.Config) *SQLite {
	db, err := sql.Open("sqlite3", cfg.Providers.DB.Path)
	if err != nil {
		panic(err)
	}

	sq := &SQLite{db: db, sql: cfg.Providers.DB.SQL}

	return sq
}

func (db *SQLite) Start() error {
	row := db.db.QueryRow(cSelectTables)
	var count int
	if err := row.Scan(&count); err != nil {
		return err
	}

	if count == 0 {
		logrus.Infof("Creating SQLite database: %s", db.sql)

		data, err := ioutil.ReadFile(db.sql)
		if err != nil {
			return err
		}

		if _, err := db.db.Exec(string(data)); err != nil {
			return err
		}
	}

	return nil
}

func (db *SQLite) Stop() error {
	return db.db.Close()
}

func (db *SQLite) Query(query string, args ...interface{}) (r *sql.Rows, err error) {
	r, err = db.db.Query(query, args...)

	if err != nil {
		logrus.WithFields(logrus.Fields{
			"err":   err.Error(),
			"type":  "Query",
			"query": query,
		}).Warn("SQLite")
	}

	return
}

func (db *SQLite) Exec(query string, args ...interface{}) (r sql.Result, err error) {
	r, err = db.db.Exec(query, args...)

	if err != nil {
		logrus.WithFields(logrus.Fields{
			"err":   err.Error(),
			"type":  "Exec",
			"query": query,
		}).Warn("SQLite")
	}

	return
}
