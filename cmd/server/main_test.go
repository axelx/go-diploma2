package main

import (
	"context"
	"log"
	"net"
	"strings"
	"testing"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/test/bufconn"

	"github.com/axelx/go-diploma2/internal/commands"
	pg "github.com/axelx/go-diploma2/internal/db"
	"github.com/axelx/go-diploma2/internal/handlers"
	pb "github.com/axelx/go-diploma2/internal/proto"
	"github.com/axelx/go-diploma2/internal/service/entity"
	"github.com/axelx/go-diploma2/internal/service/user"
)

func server(ctx context.Context) (pb.GRPCHandlerClient, func()) {
	buffer := 101024 * 1024
	lis := bufconn.Listen(buffer)

	baseServer := grpc.NewServer()

	db, err := pg.InitDB("postgres://user:password@localhost:5464/go-ya-gophkeeper")
	if err != nil {
		log.Println("Error not connect to pg", "about ERR", err.Error())
	}

	usr := user.User{DB: db}
	ent := entity.Entity{DB: db}
	gRPCsrv := handlers.PBNew(usr, ent, ":50051")
	pb.RegisterGRPCHandlerServer(baseServer, &gRPCsrv)
	go func() {
		if err := baseServer.Serve(lis); err != nil {
			log.Printf("error serving server: %v", err)
		}
	}()

	conn, err := grpc.DialContext(ctx, "",
		grpc.WithContextDialer(func(context.Context, string) (net.Conn, error) {
			return lis.Dial()
		}), grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Printf("error connecting to server: %v", err)
	}

	closer := func() {
		err := lis.Close()
		if err != nil {
			log.Printf("error closing listener: %v", err)
		}
		baseServer.Stop()
	}

	client := pb.NewGRPCHandlerClient(conn)

	return client, closer
}

func TestAuth(t *testing.T) {
	ctx := context.Background()

	client, closer := server(ctx)
	defer closer()

	type expectation struct {
		out *pb.AuthUserResponse
		err error
	}

	users := []*pb.User{{Login: "usr1", Password: "psw1"}}

	tests := map[string]struct {
		in       *pb.AuthUserRequest
		expected expectation
	}{
		"Must_Success": {
			in: &pb.AuthUserRequest{
				User: users[0],
			},
			expected: expectation{
				out: &pb.AuthUserResponse{
					Jwt: "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9",
				},
				err: nil,
			},
		},
	}

	for scenario, tt := range tests {
		t.Run(scenario, func(t *testing.T) {
			out, err := client.AuthUser(ctx, tt.in)
			if err != nil {
				if tt.expected.err.Error() != err.Error() {
					t.Errorf("Err -> \nWant: %q\nGot: %q\n", tt.expected.err, err)
				}
			} else {
				if tt.expected.out.Jwt != strings.Split(out.Jwt, ".")[0] {
					t.Errorf("Out -> \nWant: %q\nGot : %q", tt.expected.out.Jwt, strings.Split(out.Jwt, ".")[0])
				}
			}

		})
	}
}

func TestUpdateEntity(t *testing.T) {
	ctx := context.Background()

	client, closer := server(ctx)
	defer closer()

	userId, _ := commands.AuthUser(client, "usr", "psw")

	type expectation struct {
		out *pb.UpdateEntityResponse
		err error
	}

	entity := []*pb.Entity{{
		UserID:             int32(userId),
		Text:               "Test",
		BankCard:           int64(1111222233334444),
		CreatedAtTimestamp: time.Now().Unix(),
	}}

	tests := map[string]struct {
		in       *pb.UpdateEntityRequest
		expected expectation
	}{
		"Must_Success": {
			in: &pb.UpdateEntityRequest{
				Entity: entity[0],
			},
			expected: expectation{
				out: &pb.UpdateEntityResponse{
					Entity: entity[0],
				},
				err: nil,
			},
		},
	}

	for scenario, tt := range tests {
		t.Run(scenario, func(t *testing.T) {
			out, err := client.UpdateEntity(ctx, tt.in)
			if err != nil {
				if tt.expected.err.Error() != err.Error() {
					t.Errorf("Err -> \nWant: %q\nGot: %q\n", tt.expected.err, err)
				}
			} else {
				if tt.expected.out.Entity.Text != out.Entity.Text {
					t.Errorf("Out -> \nWant: %q\nGot : %q", tt.expected.out.Entity, out.Entity)
				}
			}

		})
	}
}

func TestGetEntity(t *testing.T) {
	ctx := context.Background()

	client, closer := server(ctx)
	defer closer()

	userId, jwt := commands.AuthUser(client, "usr", "psw")

	type expectation struct {
		out *pb.GetEntityResponse
		err error
	}

	entity := []*pb.Entity{{
		UserID:             int32(userId),
		Text:               "Test text",
		BankCard:           int64(1111222233334444),
		CreatedAtTimestamp: time.Now().Unix(),
	}}
	client.UpdateEntity(context.Background(), &pb.UpdateEntityRequest{
		Entity: entity[0],
		JWT:    jwt,
	})

	tests := map[string]struct {
		in       *pb.GetEntityRequest
		expected expectation
	}{
		"Must_Success": {
			in: &pb.GetEntityRequest{
				UserID: int32(userId),
				JWT:    jwt,
			},
			expected: expectation{
				out: &pb.GetEntityResponse{
					Entity: entity[0],
				},
				err: nil,
			},
		},
	}

	for scenario, tt := range tests {
		t.Run(scenario, func(t *testing.T) {
			out, err := client.GetEntity(ctx, tt.in)
			if err != nil {
				if tt.expected.err.Error() != err.Error() {
					t.Errorf("Err -> \nWant: %q\nGot: %q\n", tt.expected.err, err)
				}
			} else {
				if tt.expected.out.Entity.Text != out.Entity.Text {
					t.Errorf("Out -> \nWant: %q\nGot : %q", tt.expected.out.Entity, out.Entity)
				}
			}

		})
	}
}
