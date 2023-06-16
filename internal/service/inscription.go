package service

import (
	"context"

	"github.com/go-kratos/kratos/v2/log"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"

	pb "github.com/adshao/ordinals-indexer/api/inscription/v1"
	"github.com/adshao/ordinals-indexer/internal/biz"
	"github.com/adshao/ordinals-indexer/internal/ord/page"
)

type InscriptionService struct {
	pb.UnimplementedInscriptionServer

	p           page.PageParser
	inscription *biz.InscriptionUsecase
	log         *log.Helper
}

func NewInscriptionService(p page.PageParser, inscription *biz.InscriptionUsecase, logger log.Logger) *InscriptionService {
	return &InscriptionService{
		p:           p,
		inscription: inscription,
		log:         log.NewHelper(logger),
	}
}

func (s *InscriptionService) GetInscription(ctx context.Context, req *pb.GetInscriptionRequest) (*pb.GetInscriptionReply, error) {
	inscriptionPage := page.NewInscriptionPage(req.InscriptionUid)
	res, err := s.p.Parse(inscriptionPage)
	if err != nil {
		return nil, err
	}
	inscription := res.(*page.Inscription)
	return &pb.GetInscriptionReply{
		Data: &pb.InscriptionMessage{
			Id:            inscription.ID,
			InscriptionId: inscription.ID,
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
		},
	}, nil
}

func (s *InscriptionService) ListInscription(ctx context.Context, req *pb.ListInscriptionRequest) (*pb.ListInscriptionReply, error) {
	var inscriptionsPage *page.InscriptionsPage
	if req.InscriptionId != nil {
		inscriptionsPage = page.NewInscriptionsPage(*req.InscriptionId)
	} else {
		inscriptionsPage = page.NewInscriptionsPage()
	}
	res, err := s.p.Parse(inscriptionsPage)
	if err != nil {
		return nil, err
	}
	inscriptions := res.(*page.Inscriptions)
	var data []*pb.InscriptionMessage
	for _, uid := range inscriptions.UIDs {
		inscriptionPage := page.NewInscriptionPage(uid)
		res, err := s.p.Parse(inscriptionPage)
		if err != nil {
			return nil, err
		}
		inscription := res.(*page.Inscription)
		data = append(data, &pb.InscriptionMessage{
			Id:            inscription.ID,
			InscriptionId: inscription.ID,
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
		})
	}
	return &pb.ListInscriptionReply{
		Data: data,
		Paging: &pb.Paging{
			PrevId: inscriptions.PrevID,
			NextId: inscriptions.NextID,
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
