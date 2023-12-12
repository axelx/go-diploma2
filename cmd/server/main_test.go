package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"strings"
	"testing"
	"time"

	"github.com/jmoiron/sqlx"
	sqlmock "github.com/zhashkevych/go-sqlxmock"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/test/bufconn"

	"github.com/axelx/go-diploma2/internal/commands"
	"github.com/axelx/go-diploma2/internal/handlers"
	pb "github.com/axelx/go-diploma2/internal/proto"
	"github.com/axelx/go-diploma2/internal/service/entity"
	"github.com/axelx/go-diploma2/internal/service/user"
)

func server(ctx context.Context, db *sqlx.DB) (pb.GRPCHandlerClient, func()) {

	buffer := 101024 * 1024
	lis := bufconn.Listen(buffer)

	baseServer := grpc.NewServer()

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

	//

	db, mock, err := sqlmock.Newx()
	if err != nil {
		fmt.Println("ERROR_mock_init")
		log.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	columns := []string{"id", "login", "password"}

	mock.ExpectQuery("SELECT (.+) FROM users").
		WithArgs("usr1", "psw1").
		WillReturnRows(sqlmock.NewRows(columns).AddRow(0, "0", "0"))

	mock.ExpectExec("INSERT INTO users").
		WithArgs("usr1", "psw1").
		WillReturnResult(sqlmock.NewResult(1, 1))

	mock.ExpectQuery("SELECT (.+) FROM users").
		WithArgs("usr1", "psw1").
		WillReturnRows(sqlmock.NewRows(columns).AddRow(25, "usr116", "psw1"))

	//

	client, closer := server(ctx, db)
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

	//

	db, mock, err := sqlmock.Newx()
	if err != nil {
		fmt.Println("ERROR_mock_init")
		log.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	columns_usr := []string{"id", "login", "password"}
	columns_ent := []string{"id", "user_id", "text", "bankcard", "created_at_time_stamp", "created_at", "uploaded_at"}
	timeNowUnix := time.Now().Unix()
	tNow := time.Now()

	mock.ExpectQuery("SELECT (.+) FROM users").
		WithArgs("usr1", "psw1").
		WillReturnRows(sqlmock.NewRows(columns_usr).AddRow(25, "usr116", "psw1"))

	mock.ExpectExec("INSERT INTO entities").
		WithArgs(25, "Test_txt", "1111222233334444", timeNowUnix).
		WillReturnResult(sqlmock.NewResult(1, 1))

	mock.ExpectQuery("SELECT (.+) FROM entities").
		WithArgs(25).
		WillReturnRows(sqlmock.NewRows(columns_ent).AddRow(1, 25, "Test_txt", "1111222233334444", timeNowUnix, tNow, tNow))

	//

	client, closer := server(ctx, db)
	defer closer()

	userId, _ := commands.AuthUser(client, "usr1", "psw1")

	fmt.Println("userId--: ", userId)
	type expectation struct {
		out *pb.UpdateEntityResponse
		err error
	}

	entity := []*pb.Entity{{
		UserID:             int32(userId),
		Text:               "Test_txt",
		BankCard:           int64(1111222233334444),
		CreatedAtTimestamp: timeNowUnix,
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

	//

	db, mock, err := sqlmock.Newx()
	if err != nil {
		fmt.Println("ERROR_mock_init")
		log.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	columns_usr := []string{"id", "login", "password"}
	columns_ent := []string{"id", "user_id", "text", "bankcard", "created_at_time_stamp", "created_at", "uploaded_at"}
	timeNowUnix := time.Now().Unix()
	tNow := time.Now()

	mock.ExpectQuery("SELECT (.+) FROM users").
		WithArgs("usr1", "psw1").
		WillReturnRows(sqlmock.NewRows(columns_usr).AddRow(25, "usr116", "psw1"))

	mock.ExpectExec("INSERT INTO entities").
		WithArgs(25, "Test_txt", "1111222233334444", timeNowUnix).
		WillReturnResult(sqlmock.NewResult(1, 1))

		})
	}
}

	mock.ExpectQuery("SELECT (.+) FROM entities").
		WithArgs(25).
		WillReturnRows(sqlmock.NewRows(columns_ent).AddRow(1, 25, "Test_txt", "1111222233334444", timeNowUnix, tNow, tNow))

	client, closer := server(ctx, db)
	defer closer()

	userId, jwt := commands.AuthUser(client, "usr1", "psw1")

	type expectation struct {
		out *pb.GetEntityResponse
		err error
	}

	entity := []*pb.Entity{{
		UserID:             int32(userId),
		Text:               "Test_txt",
		BankCard:           int64(1111222233334444),
		CreatedAtTimestamp: timeNowUnix,
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
