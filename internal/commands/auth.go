package commands

import (
	"github.com/axelx/go-diploma2/internal/models"
)

type User struct {
	user   models.User
	entity models.Entity
}

type userdo interface {
	Register(string, string) string
	Auth(string, string) string
}

// общасется по grpc с сервером
func (u User) Register() models.JWT {

	return ""
}

func (u User) Auth() models.JWT {

	return ""
}
