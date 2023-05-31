package service

import (
	"context"

	"github.com/go-kratos/kratos/v2/log"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"

	pb "github.com/adshao/ordinals-indexer/api/inscription/v1"
	"github.com/adshao/ordinals-indexer/internal/biz"
)

type InscriptionService struct {
	pb.UnimplementedInscriptionServer

	inscription *biz.InscriptionUsecase
	log         *log.Helper
}

func NewInscriptionService(inscription *biz.InscriptionUsecase, logger log.Logger) *InscriptionService {
	return &InscriptionService{
		inscription: inscription,
		log:         log.NewHelper(logger),
	}
}

func (s *InscriptionService) GetInscription(ctx context.Context, req *pb.GetInscriptionRequest) (*pb.GetInscriptionReply, error) {
	res, err := s.inscription.FindByInscriptionID(ctx, req.InscriptionId)
	if err != nil {
		return nil, err
	}
	if res == nil {
		return nil, pb.ErrorInscriptionNotFound("inscription not found: %d", req.InscriptionId)
	}
	return &pb.GetInscriptionReply{
		Data: s.fromBizInscription(res),
	}, nil
}
func (s *InscriptionService) ListInscription(ctx context.Context, req *pb.ListInscriptionRequest) (*pb.ListInscriptionReply, error) {
	opt := &biz.InscriptionListOption{
		Limit:  int(req.Limit),
		Offset: int(req.Offset),
		Order:  req.OrderBy,
	}
	inscriptions, err := s.inscription.ListInscriptions(ctx, opt)
	if err != nil {
		return nil, err
	}
	totalCount, err := s.inscription.CountInscriptions(ctx, opt)
	if err != nil {
		return nil, err
	}
	var data []*pb.InscriptionMessage
	for _, inscription := range inscriptions {
		data = append(data, s.fromBizInscription(inscription))
	}
	return &pb.ListInscriptionReply{
		Data: data,
		Paging: &pb.Paging{
			TotalCount: uint64(totalCount),
			Count:      uint64(len(data)),
		},
	}, nil
}

func (s *InscriptionService) fromBizInscription(inscription *biz.Inscription) *pb.InscriptionMessage {
	return &pb.InscriptionMessage{
		Id:            int64(inscription.ID),
		InscriptionId: inscription.InscriptionID,
		Uid:           inscription.UID,
		Address:       inscription.Address,
		OutputValue:   inscription.OutputValue,
		ContentLength: inscription.ContentLength,
		ContentType:   inscription.ContentType,
		Timestamp:     timestamppb.New(inscription.Timestamp),
		GenesisHeight: inscription.GenesisHeight,
		GenesisFee:    inscription.GenesisFee,
		GenesisTx:     inscription.GenesisTx,
		Location:      inscription.Location,
		Output:        inscription.Output,
		Offset:        inscription.Offset,
	}
}
