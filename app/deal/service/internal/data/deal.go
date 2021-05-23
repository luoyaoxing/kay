package data

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"kay/app/deal/service/internal/biz"
	"kay/pkg/utils/dec"
)

type dealRepo struct {
	data *Data

	log *log.Helper
}

func NewDealRepo(data *Data, logger log.Logger) *dealRepo {
	return &dealRepo{data, log.NewHelper("deal/data", logger)}
}

func (repo *dealRepo) CreateDeal(ctx context.Context, deal *biz.Deal) error {
	repo.log.Infof("CreateDeal start deal:%s", dec.JsonEncode(deal))

	_, err := repo.data.db.Deal.Create().
		SetID(deal.DealId).
		SetSkuId(deal.SkuId).
		SetProductId(deal.ProductId).
		SetProductName(deal.ProductName).
		SetUID(deal.Uid).
		SetCustomName(deal.CustomName).
		SetPhone(deal.Phone).
		SetPrice(int(deal.Price)).
		Save(ctx)

	if err != nil {
		repo.log.Errorf("Save failed err:%s", err.Error())
		return err
	}

	repo.log.Infof("CreateDeal success")
	return nil
}

func (repo *dealRepo) GetDeal(ctx context.Context, dealId int64) (*biz.Deal, error) {
	repo.log.Infof("GetDeal start dealId:%d", dealId)

	d, err := repo.data.db.Deal.Get(ctx, dealId)
	if err != nil {
		repo.log.Errorf("Get failed err:%s", err.Error())
		return nil, err
	}

	deal := &biz.Deal{}
	deal.FromDealPO(d)

	repo.log.Infof("GetDeal success")
	return deal, nil
}
