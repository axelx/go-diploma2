package user

import (
	"context"
	"fmt"

	"github.com/jmoiron/sqlx"

	"github.com/axelx/go-diploma2/internal/db"
	"github.com/axelx/go-diploma2/internal/models"
	servicejwt "github.com/axelx/go-diploma2/internal/service/jwt"
)

type User struct {
	models.User
	DB *sqlx.DB
}

func (u User) FindUser(ctx context.Context, login, psw string) models.JWT {
	fmt.Println("-- service user: FindUser..:", login, psw)
	us := db.FindUser(ctx, u.DB, login, psw)
	if us.ID == 0 {
		fmt.Println("-- service user: FindUser..:", "need create new user")
		db.CreateNewUser(ctx, u.DB, login, psw)
		us = db.FindUser(ctx, u.DB, login, psw)
	}

	jwt := servicejwt.CreateJWT(us)

	return jwt
}
