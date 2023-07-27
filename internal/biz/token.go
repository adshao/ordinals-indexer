package biz

import (
	"context"
	"time"

	"github.com/adshao/go-brc721/sig"
	"github.com/go-kratos/kratos/v2/log"
)

const (
	ProtocolTypeBRC721 = "brc-721"
)

// Token is a Token model.
type Token struct {
	ID             int         `json:"id"`
	P              string      `json:"p"`
	Tick           string      `json:"tick"`
	TokenID        uint64      `json:"token_id"`
	TxHash         string      `json:"tx_hash"`
	BlockHeight    uint64      `json:"block_height"`
	BlockTime      time.Time   `json:"block_time"`
	Address        string      `json:"address"`
	InscriptionID  int64       `json:"inscription_id"`
	InscriptionUID string      `json:"inscription_uid"`
	CollectionID   int         `json:"collection_id"`
	Sig            sig.MintSig `json:"sig,omitempty"`
}

type TokenListOption struct {
	Limit  int
	Offset int
	Tick   string
	P      string
	Order  string
}

// TokenRepo is a Greater repo.
type TokenRepo interface {
	Create(context.Context, *Token) (*Token, error)
	Update(context.Context, *Token) (*Token, error)
	FindByTickTokenID(context.Context, string, string, uint64) (*Token, error)
	FindByInscriptionID(context.Context, int64) ([]*Token, error)
	FindByTickSigUID(context.Context, string, string, string) (*Token, error)
	List(context.Context, ...TokenListOption) ([]*Token, error)
	Delete(context.Context, int) error
	Count(context.Context, ...TokenListOption) (int, error)
}

// TokenUsecase is a Token usecase.
type TokenUsecase struct {
	repo TokenRepo
	log  *log.Helper
}

// NewTokenUsecase new a Token usecase.
func NewTokenUsecase(repo TokenRepo, logger log.Logger) *TokenUsecase {
	return &TokenUsecase{repo: repo, log: log.NewHelper(logger)}
}

// CreateToken creates a Token, and returns the new Token.
func (uc *TokenUsecase) CreateToken(ctx context.Context, g *Token) (*Token, error) {
	uc.log.WithContext(ctx).Debugf("CreateToken for inscription %d", g.InscriptionID)
	return uc.repo.Create(ctx, g)
}

// UpdateToken updates a Token, and returns the new Token.
func (uc *TokenUsecase) UpdateToken(ctx context.Context, g *Token) (*Token, error) {
	uc.log.WithContext(ctx).Debugf("UpdateToken for inscription %d", g.InscriptionID)
	return uc.repo.Update(ctx, g)
}

// FindByTickTokenID finds the Token by Tick and TokenID.
func (uc *TokenUsecase) FindByTickTokenID(ctx context.Context, p, tick string, tokenID uint64) (*Token, error) {
	uc.log.WithContext(ctx).Debugf("FindByTickTokenID for %s %s %s", p, tick, tokenID)
	return uc.repo.FindByTickTokenID(ctx, p, tick, tokenID)
}

// FindByInscriptionID finds the Token by InscriptionID.
func (uc *TokenUsecase) FindByInscriptionID(ctx context.Context, inscriptionID int64) ([]*Token, error) {
	uc.log.WithContext(ctx).Debugf("FindByInscriptionID for %d", inscriptionID)
	return uc.repo.FindByInscriptionID(ctx, inscriptionID)
}

// FindByTickSigUID finds the Token by Tick and SigUID.
func (uc *TokenUsecase) FindByTickSigUID(ctx context.Context, p, tick, sigUID string) (*Token, error) {
	uc.log.WithContext(ctx).Debugf("FindByTickSigUID for %s %s %s", p, tick, sigUID)
	return uc.repo.FindByTickSigUID(ctx, p, tick, sigUID)
}

// ListTokens lists Tokens.
func (uc *TokenUsecase) ListTokens(ctx context.Context, opt *TokenListOption) ([]*Token, error) {
	uc.log.WithContext(ctx).Debugf("ListTokens for %v", opt)
	return uc.repo.List(ctx, *opt)
}

// DeleteToken deletes a Token.
func (uc *TokenUsecase) DeleteToken(ctx context.Context, id int) error {
	uc.log.WithContext(ctx).Debugf("DeleteToken for %d", id)
	return uc.repo.Delete(ctx, id)
}

// CountTokens counts Tokens.
func (uc *TokenUsecase) CountTokens(ctx context.Context, opt *TokenListOption) (int, error) {
	uc.log.WithContext(ctx).Debugf("CountTokens for %v", opt)
	return uc.repo.Count(ctx, *opt)
}
