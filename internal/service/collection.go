package service

import (
	"context"

	"github.com/go-kratos/kratos/v2/log"
	structpb "google.golang.org/protobuf/types/known/structpb"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"

	pb "github.com/adshao/ordinals-indexer/api/collection/v1"
	"github.com/adshao/ordinals-indexer/internal/biz"
	"github.com/adshao/ordinals-indexer/internal/ord/page"
)

type CollectionService struct {
	pb.UnimplementedCollectionServer

	p                 page.PageParser
	collectionUsecase *biz.CollectionUsecase
	log               *log.Helper
}

func NewCollectionService(p page.PageParser, collectionUsecase *biz.CollectionUsecase, logger log.Logger) *CollectionService {
	return &CollectionService{
		p:                 p,
		collectionUsecase: collectionUsecase,
		log:               log.NewHelper(logger),
	}
}

func (s *CollectionService) GetCollection(ctx context.Context, req *pb.GetCollectionRequest) (*pb.GetCollectionReply, error) {
	if req.P == "" {
		req.P = biz.ProtocolTypeBRC721
	}
	collection, err := s.collectionUsecase.GetCollectionByTick(ctx, req.P, req.Tick)
	if err != nil {
		return nil, err
	}
	if collection == nil {
		return nil, pb.ErrorCollectionNotFound("collection not found: %s", req.Tick)
	}
	return &pb.GetCollectionReply{
		Data: s.fromBizCollection(collection),
	}, nil
}

func (s *CollectionService) GetInscriptionCollection(ctx context.Context, req *pb.GetInscriptionCollectionRequest) (*pb.GetCollectionReply, error) {
	collections, err := s.collectionUsecase.GetCollectionByInscriptionID(ctx, req.InscriptionId)
	if err != nil {
		return nil, err
	}
	if len(collections) == 0 {
		return nil, pb.ErrorCollectionNotFound("collection not found by inscription id: %d", req.InscriptionId)
	}
	return &pb.GetCollectionReply{
		Data: s.fromBizCollection(collections[0]),
	}, nil
}

func (s *CollectionService) ListCollections(ctx context.Context, req *pb.ListCollectionRequest) (*pb.ListCollectionReply, error) {
	opt := &biz.CollectionListOption{
		Limit:  int(req.Limit),
		Offset: int(req.Offset),
		P:      req.P,
		Tick:   req.Tick,
		Order:  req.OrderBy,
	}
	collections, err := s.collectionUsecase.ListCollections(ctx, opt)
	if err != nil {
		return nil, err
	}
	totalCount, err := s.collectionUsecase.CountCollection(ctx, opt)
	if err != nil {
		return nil, err
	}
	var data []*pb.CollectionMessage
	for _, collection := range collections {
		data = append(data, s.fromBizCollection(collection))
	}
	paging := &pb.Paging{
		TotalCount: uint64(totalCount),
		Count:      uint64(len(data)),
	}
	return &pb.ListCollectionReply{
		Data:   data,
		Paging: paging,
	}, nil
}

func (s *CollectionService) fromBizCollection(collection *biz.Collection) *pb.CollectionMessage {
	m := &pb.CollectionMessage{
		P:              collection.P,
		Tick:           collection.Tick,
		Max:            collection.Max,
		Supply:         collection.Supply,
		BaseUri:        collection.BaseURI,
		Name:           collection.Name,
		Description:    collection.Description,
		Image:          collection.Image,
		TxHash:         collection.TxHash,
		BlockHeight:    collection.BlockHeight,
		BlockTime:      timestamppb.New(collection.BlockTime),
		Address:        collection.Address,
		InscriptionId:  collection.InscriptionID,
		InscriptionUid: collection.InscriptionUID,
	}
	for _, attr := range collection.Attributes {
		at, _ := structpb.NewStruct(attr)
		m.Attributes = append(m.Attributes, at)
	}
	return m
}
