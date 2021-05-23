package biz

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"kay/pkg/utils/dec"
)

type ProductRepo interface {
	CreateProduct(ctx context.Context, p *Product) (int64, error)
	GetProductById(ctx context.Context, skuId int64) (*Product, error)
	ListProductByProductId(ctx context.Context, productId int64) ([]*Product, error)
}

type ProductUseCase struct {
	repo ProductRepo
	log  *log.Helper
}

func NewProductUseCase(repo ProductRepo, logger log.Logger) *ProductUseCase {
	return &ProductUseCase{repo: repo, log: log.NewHelper("product/biz", logger)}
}

func (uc *ProductUseCase) CreateProduct(ctx context.Context, p *Product) (int64, error) {
	uc.log.Infof("create start p:%s ", dec.JsonEncode(p))

	skuId, err := uc.repo.CreateProduct(ctx, p)
	if err != nil {
		uc.log.Errorf("create product failed err:%s", err.Error())
		return 0, err
	}

	uc.log.Info("create success skuId:%s", skuId)
	return skuId, nil
}

func (uc *ProductUseCase) GetProductInfo(ctx context.Context, skuId int64) (*Product, error) {
	uc.log.Infof("getProductInfo start skuId:%d", skuId)

	p, err := uc.repo.GetProductById(ctx, skuId)
	if err != nil {
		uc.log.Errorf("get product failed err:%s", err.Error())
		return nil, err
	}

	uc.log.Info("getProductInfo success")
	return p, nil
}

func (uc *ProductUseCase) ListProductByProductId(ctx context.Context, productId int64) ([]*Product, error) {
	uc.log.Infof("listProductByProductId start productId:%d", productId)

	pl, err := uc.repo.ListProductByProductId(ctx, productId)
	if err != nil {
		uc.log.Errorf("list product failed err:%s", err.Error())
		return nil, err
	}

	uc.log.Info("listProductByProductId success")
	return pl, nil
}
