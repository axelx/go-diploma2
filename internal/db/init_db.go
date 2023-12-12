package db

import (
	_ "github.com/jackc/pgx/stdlib"
	"github.com/jmoiron/sqlx"
)

// подключение к базе
// миграции

func InitDB(url string) (*sqlx.DB, error) {

	db, err := sqlx.Connect("pgx", url)
	if err != nil {
		return db, err
	}

	db.SetMaxOpenConns(10)

	// миграции
	err = createTable(db)
	if err != nil {
		return db, err
	}

	return db, nil
}
