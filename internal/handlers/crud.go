package handlers

import (
	"context"
	pb "github.com/axelx/go-diploma2/internal/proto"
	servicejwt "github.com/axelx/go-diploma2/internal/service/jwt"
	"github.com/axelx/go-diploma2/internal/utils"
)

// ProtoHandler  GetEntity
func (s *ProtoHandler) GetEntity(ctx context.Context, in *pb.GetEntityRequest) (*pb.GetEntityResponse, error) {

	userID := servicejwt.UserIDFromJwt(in.JWT)
	ent := s.ServiceEnt.Read(ctx, userID)

	var response pb.GetEntityResponse
	response.Entity = utils.EntToProto(ent)
	return &response, nil
}

// ProtoHandler  UpdateEntity
func (s *ProtoHandler) UpdateEntity(ctx context.Context, in *pb.UpdateEntityRequest) (*pb.UpdateEntityResponse, error) {
	e := utils.ProtoToEnt(in.Entity)
	ent := s.ServiceEnt.UpdateORCreate(ctx, e)

	var response pb.UpdateEntityResponse
	response.Entity = utils.EntToProto(ent)
	return &response, nil
}
