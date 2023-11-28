package user

import (
	"github.com/axelx/go-diploma2/internal/models"
	"github.com/jmoiron/sqlx"
)

type User struct {
	models.User
	DB *sqlx.DB
}

func (u User) Auth(login, psw string) models.JWT {

}

func (u User) Register(login, psw string) models.JWT {

}
