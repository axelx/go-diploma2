package db

import (
	"context"
	"fmt"

	"github.com/jmoiron/sqlx"

	"github.com/axelx/go-diploma2/internal/models"
)

func CreateNewUser(ctx context.Context, db *sqlx.DB, login, password string) error {
	_, err := db.ExecContext(ctx, `INSERT INTO users (login, password) VALUES ($1, $2)`, login, password)
	if err != nil {
		fmt.Println("Error CreateNewUser :", "about ERR", err.Error())

		return err
	}
	return nil
}

func FindUser(ctx context.Context, db *sqlx.DB, login, password string) models.User {
	row := db.QueryRowContext(ctx,
		` SELECT * FROM users WHERE login = $1 AND password = $2 `, login, password)

	var usr models.User
	err := row.Scan(&usr.ID, &usr.Login, &usr.Password)
	if err != nil {
		fmt.Println("Error FindUser :", "about ERR", err.Error())

		return models.User{}
	}
	return usr
}
