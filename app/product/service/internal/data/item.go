package data

import (
	"context"
	"errors"
	"github.com/go-kratos/kratos/v2/log"
	"kay/app/product/service/internal/biz"
	"kay/app/product/service/internal/common"
	"kay/app/product/service/internal/data/ent"
	"kay/app/product/service/internal/data/ent/item"
	"kay/pkg/utils/dec"
)

type itemStockRepo struct {
	data *Data

	log *log.Helper
}

func NewItemStockRepo(data *Data, logger log.Logger) *itemStockRepo {
	return &itemStockRepo{data: data, log: log.NewHelper("itemStock/data", logger)}
}

func (repo *itemStockRepo) CreateStock(ctx context.Context, stock *biz.ItemStock) error {
	repo.log.Infof("CreateStock start stock:%s", dec.JsonEncode(stock))

	_, err := repo.data.db.Item.Create().
		SetID(stock.SkuId).
		SetTotalStock(int(stock.TotalStock)).
		SetLeftStock(int(stock.TotalStock)).
		Save(ctx)

	if err != nil {
		repo.log.Errorf("Save failed err:%s", err.Error())
		return err
	}

	repo.log.Infof("CreateStock start success")
	return nil
}

func (repo *itemStockRepo) CreateStockWithTx(ctx context.Context, stock *biz.ItemStock) error {
	repo.log.Infof("CreateStock start stock:%s", dec.JsonEncode(stock))

	txValue := ctx.Value(common.TxKey)
	tx, ok := txValue.(*ent.Tx)
	if !ok {
		return errors.New("事务对象不存在")
	}

	_, err := tx.Item.Create().
		SetID(stock.SkuId).
		SetTotalStock(int(stock.TotalStock)).
		SetLeftStock(int(stock.TotalStock)).
		Save(ctx)

	if err != nil {
		repo.log.Errorf("Save failed err:%s", err.Error())
		return err
	}

	repo.log.Infof("CreateStock start success")
	return nil
}

func (repo *itemStockRepo) DeStock(ctx context.Context, skuId int64) (*biz.ItemStock, error) {
	repo.log.Infof("DeStock start skuId:%d", skuId)

	is, err := repo.data.db.Item.Get(ctx, skuId)
	if err != nil {
		repo.log.Errorf("Get skuId:%d failed err:%s", skuId, err.Error())
		return nil, err
	}

	if is == nil {
		repo.log.Errorf("Get stock skuId:%d empty", skuId)
		return nil, errors.New("empty stock")
	}
	repo.log.Infof("DeStock skuId:%d before is:%s", skuId, dec.JsonEncode(is))

	is.LeftStock--
	is.ConsumeStock++
	repo.log.Infof("DeStock skuId:%d after is:%s", skuId, dec.JsonEncode(is))

	_, err = repo.data.db.Item.Update().Where(item.IDEQ(skuId)).
		SetLeftStock(is.LeftStock).
		SetConsumeStock(is.ConsumeStock).
		Save(ctx)

	if err != nil {
		repo.log.Errorf("DeStock skuId:%d failed err:%s", skuId, err.Error())
		return nil, err
	}

	itemStock := &biz.ItemStock{}
	itemStock.FromItemStockPo(is)

	repo.log.Infof("DeStock start success")
	return itemStock, nil
}
