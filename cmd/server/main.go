package main

import (
	config "github.com/axelx/go-diploma2/internal/config/server"
	pg "github.com/axelx/go-diploma2/internal/db"
	"github.com/axelx/go-diploma2/internal/handlers"
	"github.com/axelx/go-diploma2/internal/proto"
	"github.com/axelx/go-diploma2/internal/service/entity"
	"github.com/axelx/go-diploma2/internal/service/user"
	"google.golang.org/grpc"
	"log"
	"net"
)

func main() {
	conf := config.NewConfigServer()

	db, err := pg.InitDB(conf.DatabaseDSN)
	if err != nil {
		log.Println("Error not connect to pg", "about ERR", err.Error())
	}
	log.Println("hello")
	usr := user.User{DB: db}
	ent := entity.Entity{DB: db}

	gRPCsrv := handlers.PBNew(usr, ent, conf.RunAddr)
	lis, err := net.Listen("tcp", conf.RunAddr)

	if err != nil {
		log.Println(err)
	}
	s := grpc.NewServer(grpc.UnaryInterceptor(handlers.UnaryInterceptor))
	go_diploma2.RegisterGRPCHandlerServer(s, &gRPCsrv)

	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Println(err)
	}
}
