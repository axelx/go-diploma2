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
				    id serial PRIMARY KEY,
				    user_id int NOT NULL UNIQUE,
				    text text,
				    BankCard varchar(100) NOT NULL,
					created_at_time_stamp INT NOT NULL, --DEFAULT CURRENT_TIME,
					created_at TIMESTAMP NOT NULL, --DEFAULT CURRENT_TIME,
					uploaded_at TIMESTAMP NOT NULL --DEFAULT CURRENT_TIME,
				);
			`)

	return err
}
