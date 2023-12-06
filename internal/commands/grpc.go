package commands

import (
	"context"
	"fmt"
	"github.com/axelx/go-diploma2/internal/models"
	pb "github.com/axelx/go-diploma2/internal/proto"
	servicejwt "github.com/axelx/go-diploma2/internal/service/jwt"
	"github.com/axelx/go-diploma2/internal/utils"
	"log"
	"time"
)

// func EntityHandlerUpdate обновляем пользовательские данные
func EntityHandlerUpdate(userId int, jwt string, c pb.GRPCHandlerClient, txt, bc string) {
	fmt.Println("-EntityHandlerUpdate--: userId: ", userId, "__", c)
	entity := []*pb.Entity{{
		//ID:                 1,
		UserID:             int32(userId),
		Text:               txt,
		BankCard:           int64(utils.StrToInt(bc)),
		CreatedAtTimestamp: time.Now().Unix(),
	}}
	_, err := c.UpdateEntity(context.Background(), &pb.UpdateEntityRequest{
		Entity: entity[0],
		JWT:    jwt,
	})
	if err != nil {
		log.Fatal(err)
	}
}

// func EntityHandler получаем пользовательские данные
func EntityHandler(userId int, jwt string, c pb.GRPCHandlerClient) models.Entity {
	resp, err := c.GetEntity(context.Background(), &pb.GetEntityRequest{
		JWT: jwt,
	})
	if err != nil {
		log.Fatal(err)
	}
	return utils.ProtoToEnt(resp.Entity)
}

// func AuthUser получаем токен пользователя
func AuthUser(c pb.GRPCHandlerClient, log, psw string) (int, string) {
	users := []*pb.User{{Login: log, Password: psw}}
	var resUserID int
	resp, err := c.AuthUser(context.Background(), &pb.AuthUserRequest{
		User: users[0],
	})
	if err != nil {
		fmt.Println("-----000", err)
	}
	if resp.Error != "" {
		fmt.Println(resp.Error)
	}
	resUserID = servicejwt.UserIDFromJwt(resp.Jwt)
	return resUserID, resp.Jwt
}
