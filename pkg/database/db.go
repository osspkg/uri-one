package database

import (
	"context"
	"database/sql"
	"time"

	"github.com/deweppro/go-logger"
	"github.com/deweppro/go-orm"
	"github.com/deweppro/go-orm/schema"
	"github.com/deweppro/go-orm/schema/sqlite"
)

type Database struct {
	conn schema.Connector
	pool orm.StmtInterface
	log  logger.Logger
}

func New(log logger.Logger, conf *sqlite.Config) (*Database, error) {
	conn, err := sqlite.New(conf)
	if err != nil {
		return nil, err
	}
	pool := orm.NewDB(conn, orm.Plugins{Logger: log})
	db := &Database{
		conn: conn,
		pool: pool.Pool("base"),
		log:  log,
	}
	return db, nil
}

func (v *Database) Up() error {
	if err := v.pool.Ping(); err != nil {
		return err
	}
	return v.pool.Call("check_tables", func(conn *sql.Conn, ctx context.Context) error {
		row := conn.QueryRowContext(ctx, checkTables)
		var count int
		if err := row.Scan(&count); err != nil {
			return err
		}
		if err := row.Err(); err != nil {
			return err
		}
		if count == 0 {
			v.log.Infof("Creating SQLite database")
			if _, err := conn.ExecContext(ctx, sqlDefaul); err != nil {
				return err
			}
		}
		return nil
	})
}

func (v *Database) Down() error {
	return v.conn.Close()
}

func (v *Database) GetUrl(id int) (string, error) {
	var data string
	err := v.pool.Call("get_url", func(conn *sql.Conn, ctx context.Context) error {
		row := conn.QueryRowContext(ctx, getUri, id)
		if err := row.Scan(&data); err != nil {
			return err
		}
		return row.Err()
	})
	if err != nil {
		return "", err
	}
	if len(data) == 0 {
		return "", ErrRecordNotFound
	}
	return data, nil
}

func (v *Database) SetUrl(data string) (int, error) {
	var id int
	err := v.pool.Call("set_url", func(conn *sql.Conn, ctx context.Context) error {
		result, err := conn.ExecContext(ctx, setUri, data, time.Now().Unix())
		if err != nil {
			return err
		}
		i, err := result.LastInsertId()
		if err != nil {
			return err
		}
		id = int(i)
		return nil
	})
	if err != nil {
		return 0, err
	}
	return id, nil
}
