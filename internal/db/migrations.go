package db

import (
	"context"
	"github.com/jmoiron/sqlx"
)

func createTable(db *sqlx.DB) error {
	_, err := db.ExecContext(context.Background(),
		` CREATE TABLE IF NOT EXISTS users (
					 id serial PRIMARY KEY,
					 login varchar(450) NOT NULL UNIQUE,
					 password varchar(450) NOT NULL
				);

				CREATE TABLE IF NOT EXISTS entities (
				);
			`)

	return err
}
