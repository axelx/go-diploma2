package entity

import (
	"github.com/axelx/go-diploma2/internal/models"
	"github.com/jmoiron/sqlx"
)

type Entity struct {
	models.Entity
	DB *sqlx.DB
}

func (u Entity) Create() models.Entity {

	return entity
}

func (u Entity) Read() models.Entity {

	return entity
}

func (u Entity) Update() models.Entity {

	return entity
}

func (u Entity) Delete() {

}
