package db

import (
	"github.com/jmoiron/sqlx"
)

func CreateNewUser(db *sqlx.DB, login, password string) (jwt, error) {
}

func AuthUser(db *sqlx.DB, login, password string) (jwt, error) {
}
