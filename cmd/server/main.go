package main

import (
	"github.com/axelx/go-diploma2/internal/handlers"
	"log"
	"net"
)

func main() {

	db, err := pg.InitDB(conf.FlagDatabaseDSN, lg)

	gRPCsrv := handlers.PBNew(db, NewDBStorage, ":50051")
	lis, err := net.Listen("tcp", gsrv.Addr)
	if err != nil {
		log.Println(err)
	}
	s := grpc.NewServer()
	go_diploma2.RegisterCrudServer(s, &gRPCsrv)
	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Println(err)
	}
}
