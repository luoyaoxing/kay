package biz

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"kay/pkg/utils/dec"
)

type ItemStockRepo interface {
	CreateStock(ctx context.Context, stock *ItemStock) error
	CreateStockWithTx(ctx context.Context, stock *ItemStock) error
	DeStock(ctx context.Context, skuId int64) (*ItemStock, error)
}

type ItemStockUseCase struct {
	repo ItemStockRepo

	log *log.Helper
}

func NewItemStockUseCase(repo ItemStockRepo, logger log.Logger) *ItemStockUseCase {
	return &ItemStockUseCase{repo: repo, log: log.NewHelper("itemStock/biz", logger)}
}

func (uc *ItemStockUseCase) CreateStockWithTx(ctx context.Context, stock *ItemStock) error {
	uc.log.Infof("CreateStock start stock:%s", dec.JsonEncode(stock))

	err := uc.repo.CreateStockWithTx(ctx, stock)
	if err != nil {
		uc.log.Errorf("CreateStock failed err:%s", err.Error())
		return err
	}

	uc.log.Infof("CreateStock success")
	return nil
}

func (uc *ItemStockUseCase) CreateStock(ctx context.Context, stock *ItemStock) error {
	uc.log.Infof("CreateStock start stock:%s", dec.JsonEncode(stock))

	err := uc.repo.CreateStock(ctx, stock)
	if err != nil {
		uc.log.Errorf("CreateStock failed err:%s", err.Error())
		return err
	}

	uc.log.Infof("CreateStock success")
	return nil
}

func (uc *ItemStockUseCase) DoDeStock(ctx context.Context, product *Product) (*ItemStock, error) {
	uc.log.Infof("DeStock start product:%d", dec.JsonEncode(product))

	// do something

	skuId := product.SkuId
	is, err := uc.repo.DeStock(ctx, skuId)
	if err != nil {
		uc.log.Errorf("DeStock failed err:%s", err.Error())
		return nil, err
	}

	uc.log.Infof("DeStock success is:%s", dec.JsonEncode(is))
	return is, nil
}
