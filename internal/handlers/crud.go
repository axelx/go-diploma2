package handlers

import (
	"context"
	"github.com/axelx/go-diploma2/internal/models"
	"github.com/jmoiron/sqlx"
)

// ProtoHandler data for gRPC server
type ProtoHandler struct {
	pb.UnimplementedMetricsServer
	DB         *sqlx.DB
	DBPostgres *pg.PgStorage
	Addr       string
}

func PBNew(db *sqlx.DB, NewDBStorage *pg.PgStorage, addr string) ProtoHandler {
	return ProtoHandler{
		DB:         db,
		DBPostgres: NewDBStorage,
		Addr:       addr,
	}
}

func serverInterceptor(jwt models.JWT) {

}

// ProtoHandler  CreateEntity
func (s *ProtoHandler) CreateEntity(ctx context.Context, in *pb.UpdateMetricRequest) (*pb.UpdateMetricResponse, error) {
}

// ProtoHandler  GetEntity
func (s *ProtoHandler) ReadtEntity(ctx context.Context, in *pb.GetMetricRequest) (*pb.GetMetricResponse, error) {
}

// ProtoHandler  UpdateEntity
func (s *ProtoHandler) UpdateEntity(ctx context.Context, in *pb.UpdateMetricRequest) (*pb.UpdateMetricResponse, error) {
}

// ProtoHandler  DeleteEntity
func (s *ProtoHandler) DeleteEntity(ctx context.Context, in *pb.UpdateMetricsRequest) (*pb.UpdateMetricsResponse, error) {
}
