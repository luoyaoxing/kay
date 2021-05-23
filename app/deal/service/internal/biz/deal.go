package biz

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"kay/pkg/utils/dec"
)

type DealRepo interface {
	CreateDeal(ctx context.Context, deal *Deal) error
	GetDeal(ctx context.Context, dealId int64) (*Deal, error)
}

type DealUseCase struct {
	repo DealRepo

	log *log.Helper
}

func NewDealUseCase(repo DealRepo, logger log.Logger) *DealUseCase {
	return &DealUseCase{repo: repo, log: log.NewHelper("deal/biz", logger)}
}

func (uc *DealUseCase) CreateDeal(ctx context.Context, deal *Deal) error {
	uc.log.Infof("CreateDeal start deal:%s", dec.JsonEncode(deal))

	err := uc.repo.CreateDeal(ctx, deal)
	if err != nil {
		uc.log.Errorf("CreateDeal failed err:%s", err.Error())
		return err
	}

	uc.log.Infof("CreateDeal success")
	return nil
}

func (uc *DealUseCase) GetDealInfo(ctx context.Context, dealId int64) (*Deal, error) {
	uc.log.Infof("GetDealInfo start dealId:%d", dealId)

	deal, err := uc.repo.GetDeal(ctx, dealId)
	if err != nil {
		uc.log.Errorf("GetDeal failed err:%s", err.Error())
		return nil, err
	}

	uc.log.Infof("GetDealInfo success")
	return deal, nil
}
