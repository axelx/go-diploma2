package db

import (
	"github.com/axelx/go-diploma2/internal/models"
	"github.com/jmoiron/sqlx"
)

func CreateEntity(db *sqlx.DB, models.Entity, jwt string) (models.Entity, error) {
	return entity, nil
}

func ReadEntity(db *sqlx.DB, jwt string) (models.Entity, error) {
	return entity, nil
}

func UpdateEntity(db *sqlx.DB, models.Entity, jwt string) (models.Entity, error) {
	return nil
}

func DeleteEntity(db *sqlx.DB, models.Entity, jwt string) error {
}
