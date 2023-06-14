package service

import (
	"context"

	"github.com/go-kratos/kratos/v2/log"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"

	pb "github.com/adshao/ordinals-indexer/api/token/v1"
	"github.com/adshao/ordinals-indexer/internal/biz"
)

type TokenService struct {
	pb.UnimplementedTokenServer

	tokenUsecase *biz.TokenUsecase
	log          *log.Helper
}

func NewTokenService(tokenUsecase *biz.TokenUsecase, logger log.Logger) *TokenService {
	return &TokenService{
		tokenUsecase: tokenUsecase,
		log:          log.NewHelper(logger),
	}
}

func (s *TokenService) GetToken(ctx context.Context, req *pb.GetTokenRequest) (*pb.TokenReply, error) {
	if req.Tick == "" {
		tokens, err := s.tokenUsecase.FindByInscriptionID(ctx, req.InscriptionId)
		if err != nil {
			return nil, err
		}
		if len(tokens) == 0 {
			return nil, pb.ErrorTokenNotFound("token not found: %d", req.InscriptionId)
		}
		return &pb.TokenReply{
			Data: s.fromBizToken(tokens[0]),
		}, nil
	}
	if req.P == "" {
		req.P = biz.ProtocolTypeBRC721
	}
	token, err := s.tokenUsecase.FindByTickTokenID(ctx, req.P, req.Tick, req.TokenId)
	if err != nil {
		return nil, err
	}
	if token == nil {
		return nil, pb.ErrorTokenNotFound("token not found: %d", req.TokenId)
	}
	return &pb.TokenReply{
		Data: s.fromBizToken(token),
	}, nil
}

func (s *TokenService) ListTokens(ctx context.Context, req *pb.ListTokenRequest) (*pb.ListTokenReply, error) {
	opt := &biz.TokenListOption{
		Limit:  int(req.Limit),
		Offset: int(req.Offset),
		P:      req.P,
		Tick:   req.Tick,
		Order:  req.OrderBy,
	}
	tokens, err := s.tokenUsecase.ListTokens(ctx, opt)
	if err != nil {
		return nil, err
	}
	totalCount, err := s.tokenUsecase.CountTokens(ctx, opt)
	if err != nil {
		return nil, err
	}
	var data []*pb.TokenMessage
	for _, token := range tokens {
		data = append(data, s.fromBizToken(token))
	}
	paging := &pb.Paging{
		TotalCount: uint64(totalCount),
		Count:      uint64(len(data)),
	}
	return &pb.ListTokenReply{
		Data:   data,
		Paging: paging,
	}, nil
}

func (s *TokenService) fromBizToken(token *biz.Token) *pb.TokenMessage {
	return &pb.TokenMessage{
		P:              token.P,
		Tick:           token.Tick,
		TokenId:        token.TokenID,
		TxHash:         token.TxHash,
		BlockHeight:    token.BlockHeight,
		BlockTime:      timestamppb.New(token.BlockTime),
		Address:        token.Address,
		InscriptionId:  token.InscriptionID,
		InscriptionUid: token.InscriptionUID,
	}
}
