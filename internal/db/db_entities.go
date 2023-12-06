package db

import (
	"context"
	"fmt"
	"github.com/axelx/go-diploma2/internal/models"
	"github.com/jmoiron/sqlx"
)

func ReadEntity(ctx context.Context, db *sqlx.DB, userID int) (models.Entity, error) {
	row := db.QueryRowContext(ctx, `SELECT * FROM entities WHERE user_id = $1`, userID)

	var ent models.Entity
	err := row.Scan(&ent.ID, &ent.UserID, &ent.Text, &ent.BankCard, &ent.CreatedAtTimestamp, &ent.CreatedAt, &ent.UpdatedAt)
	if err != nil {
		fmt.Println("Error Entity :", "about ERR", err.Error())
		return models.Entity{}, err
	}
	return ent, nil
}

func UpdateORCreateEntity(ctx context.Context, db *sqlx.DB, ent models.Entity) error {
	fmt.Println(ctx, db, ent)
	_, err := db.DB.ExecContext(ctx,
		`INSERT INTO entities (user_id, text, bankcard, created_at_time_stamp, created_at, uploaded_at) 
				VALUES ($1, $2, $3, $4, NOW(), NOW())
				ON CONFLICT (user_id) DO UPDATE SET text = $2, bankcard = $3, created_at_time_stamp = $4;`,
		ent.UserID, ent.Text, ent.BankCard, ent.CreatedAtTimestamp)

	if err != nil {
		fmt.Println("Error UpdateORCreateEntity :", "about ERR", err.Error())
		return err
	}

	return nil
}
