/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"log"

	"github.com/spf13/cobra"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	"github.com/axelx/go-diploma2/internal/commands"
	config "github.com/axelx/go-diploma2/internal/config/client"
	pb "github.com/axelx/go-diploma2/internal/proto"
	"github.com/axelx/go-diploma2/internal/ui"
)

// startCmd represents the start command
var startCmd = &cobra.Command{
	Use:   "start",
	Short: "CLI interface for GophKeeper",
	Long:  "CLI interface for GophKeeper",
	Run: func(cmd *cobra.Command, args []string) {
		createNewDialog()
	},
}

func init() {
	rootCmd.AddCommand(startCmd)
}

func createNewDialog() {
	userPromptContent := ui.PromptContent{
		"Имя пользователя должно быть больше 2х букв",
		"Введите имя пользователя",
	}
	usr := ui.PromptGetInput(userPromptContent, "Имя")

	pswPromptContent := ui.PromptContent{
		"Пароль пользователя должен быть больше 2х букв",
		"Введите пароль",
	}
	psw := ui.PromptGetInputPsw(pswPromptContent)

	conf := config.NewConfigClient()
	conn, err := grpc.Dial(conf.RunAddr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()
	c := pb.NewGRPCHandlerClient(conn)
	userId, jwt := commands.AuthUser(c, usr, psw)

	for {
		en := commands.EntityHandler(userId, jwt, c)
		fmt.Printf("\033[0;36m-------Ваши данныые-------\033[0m\n")
		fmt.Printf("Text:\033[1;36m%s\033[0m, BankCard:\u001B[1;36m%s\u001B[0m\n", en.Text, en.BankCard)

		s := ui.PromptChose3()
		txt := ""
		bc := ""
		if s == "Изменить" {

			text := ui.PromptContent{
				"Текст должно быть больше 2х букв",
				"Введите текст для сохранения",
			}
			txt = ui.PromptGetInput(text, "Текст")

			bca := ui.PromptContent{
				"Банковская карта должна быть 16 цифр",
				"Введите 16 цифр для сохранения ",
			}
			bc = ui.PromptGetInput(bca, "Банковская карта")

			conf := ui.PromptConfirm()
			if conf == "y" {
				commands.EntityHandlerUpdate(userId, jwt, c, txt, bc)
			}
		} else if s == "Выход" {
			break
		}
	}
}
