package handlers

import (
	"context"
	"fmt"
	"strings"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/axelx/go-diploma2/internal/models"
	pb "github.com/axelx/go-diploma2/internal/proto"
	serviceentity "github.com/axelx/go-diploma2/internal/service/entity"
	servicejwt "github.com/axelx/go-diploma2/internal/service/jwt"
	serviceuser "github.com/axelx/go-diploma2/internal/service/user"
)

// ProtoHandler data for gRPC server
type ProtoHandler struct {
	pb.UnimplementedGRPCHandlerServer
	ServiceUsr serviceuser.User
	ServiceEnt serviceentity.Entity
	Addr       string
}

// func PBNew(db *sqlx.DB, NewDBStorage *pg.PgStorage, addr string) ProtoHandler {
func PBNew(usr serviceuser.User, ent serviceentity.Entity, addr string) ProtoHandler {
	return ProtoHandler{
		ServiceEnt: ent, // service.user
		ServiceUsr: usr, // service.user
		Addr:       addr,
	}
}

// ProtoHandler  RegisterUser
func (s *ProtoHandler) RegisterUser(ctx context.Context, in *pb.RegisterUserRequest) (*pb.RegisterUserResponse, error) {

	fmt.Println("--RegisterUser:", in)
	var response pb.RegisterUserResponse
	return &response, nil
}

// ProtoHandler  AuthrUser
func (s *ProtoHandler) AuthUser(ctx context.Context, in *pb.AuthUserRequest) (*pb.AuthUserResponse, error) {
	u := in.User
	jwt := s.ServiceUsr.FindUser(ctx, u.Login, u.Password)

	var response pb.AuthUserResponse
	response.Jwt = string(jwt)

	return &response, nil
}

func UnaryInterceptor(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
	method := strings.Split(info.FullMethod, "/")
	if method[2] == "AuthUser" {
		return handler(ctx, req)
	}

	getEntityReq, ok1 := req.(*pb.GetEntityRequest)
	if !ok1 {
		fmt.Println("unaryInterceptor: не удалось преобразовать запрос в тип GetEntityRequest")
	}
	UpdateEntityRequest, ok2 := req.(*pb.UpdateEntityRequest)
	if !ok1 {
		fmt.Println("unaryInterceptor: не удалось преобразовать запрос в тип UpdateEntityRequest")
	}

	jwt := ""
	if ok1 {
		jwt = getEntityReq.JWT
	} else if ok2 {
		jwt = UpdateEntityRequest.JWT
	}

	fmt.Println("unaryInterceptor: ok1, ok2", ok1, ok2)
	servicejwt.CheckJWT(models.JWT(jwt))
	if !servicejwt.CheckJWT(models.JWT(jwt)) {
		return nil, status.Error(codes.Unauthenticated, "invalid token")
	}
	return handler(ctx, req)
}
