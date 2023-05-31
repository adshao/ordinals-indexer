package server

import (
	collectionv1 "github.com/adshao/ordinals-indexer/api/collection/v1"
	inscriptionv1 "github.com/adshao/ordinals-indexer/api/inscription/v1"
	tokenv1 "github.com/adshao/ordinals-indexer/api/token/v1"
	"github.com/adshao/ordinals-indexer/internal/conf"
	"github.com/adshao/ordinals-indexer/internal/service"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware/recovery"
	"github.com/go-kratos/kratos/v2/transport/grpc"
)

// NewGRPCServer new a gRPC server.
func NewGRPCServer(c *conf.Server, collection *service.CollectionService, token *service.TokenService, inscription *service.InscriptionService, logger log.Logger) *grpc.Server {
	var opts = []grpc.ServerOption{
		grpc.Middleware(
			recovery.Recovery(),
		),
	}
	if c.Grpc.Network != "" {
		opts = append(opts, grpc.Network(c.Grpc.Network))
	}
	if c.Grpc.Addr != "" {
		opts = append(opts, grpc.Address(c.Grpc.Addr))
	}
	if c.Grpc.Timeout != nil {
		opts = append(opts, grpc.Timeout(c.Grpc.Timeout.AsDuration()))
	}
	srv := grpc.NewServer(opts...)
	tokenv1.RegisterTokenServer(srv, token)
	collectionv1.RegisterCollectionServer(srv, collection)
	inscriptionv1.RegisterInscriptionServer(srv, inscription)
	return srv
}
