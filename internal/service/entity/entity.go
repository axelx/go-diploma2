package entity

import (
	"context"
	"fmt"

	"github.com/jmoiron/sqlx"

	"github.com/axelx/go-diploma2/internal/db"
	"github.com/axelx/go-diploma2/internal/models"
)

type Entity struct {
	models.Entity
	DB *sqlx.DB
}

func (u Entity) Create() models.Entity {

	return u.Entity
}

func (u Entity) Read(ctx context.Context, userID int) *models.Entity {

	ent, err := db.ReadEntity(ctx, u.DB, userID)
	if err != nil {
		fmt.Println("Read Err: ", err)
	}

	return &ent
}

func (u Entity) UpdateORCreate(ctx context.Context, ent models.Entity) *models.Entity {

	err := db.UpdateORCreateEntity(ctx, u.DB, ent)
	if err != nil {
		fmt.Println("UpdateORCreate Err: ", err)
	}
	ent, err = db.ReadEntity(ctx, u.DB, ent.UserID)
	if err != nil {
		fmt.Println("UpdateORCreate read Err: ", err)
	}
	return &ent
}
