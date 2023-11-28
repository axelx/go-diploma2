package commands

import (
	"github.com/axelx/go-diploma2/internal/models"
)

type ProcessEntity interface {
	Create(int, string, float64, chan string) error
	Read(string) (models.Entity, error)
	Update(string) bool
	Delete(int) (models.Entity, error)
}

//  общасется по grpc с сервером

func (u User) Create() models.Entity {

	return entity
}

func (u User) Read() models.Entity {

	return entity
}

func (u User) Update() models.Entity {

	return entity
}

func (u User) Delete() {

}
