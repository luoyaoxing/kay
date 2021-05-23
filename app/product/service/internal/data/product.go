package data

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"kay/app/product/service/internal/biz"
	"kay/app/product/service/internal/data/ent/product"
	"kay/pkg/utils/dec"
)

type productRepo struct {
	data *Data
	log  *log.Helper
}

func NewProductRepo(data *Data, logger log.Logger) *productRepo {
	return &productRepo{data: data, log: log.NewHelper("product/data", logger)}
}

func (repo *productRepo) CreateProduct(ctx context.Context, p *biz.Product) (int64, error) {
	repo.log.Infof("create product start p:%s", dec.JsonEncode(p))

	product, err := repo.data.db.Product.Create().
		SetID(p.SkuId).
		SetVersionId(p.VersionId).
		SetVersionName(p.VersionName).
		SetProductId(p.ProductId).
		SetProductName(p.ProductName).
		SetPrice(p.Price).
		SetProductDesc(p.ProductDesc).
		SetAttr(p.Attr).
		Save(ctx)

	if err != nil {
		repo.log.Errorf("save failed err:%s", err.Error())
		return 0, err
	}

	repo.log.Debugf("save product:%s", dec.JsonEncode(product))
	skuId := product.ID
	repo.log.Infof("create product success skuId:%d", skuId)
	return skuId, err
}

func (repo *productRepo) GetProductById(ctx context.Context, skuId int64) (*biz.Product, error) {
	repo.log.Infof("GetProductById start skuId:%s", skuId)

	p, err := repo.data.db.Product.Get(ctx, skuId)
	if err != nil {
		repo.log.Errorf("get failed skuId:%d err:%s", skuId, err.Error())
		return nil, err
	}

	productDo := &biz.Product{}
	productDo.FromProductPo(p)

	repo.log.Info("GetProductById success")
	return productDo, err
}

func (repo *productRepo) ListProductByProductId(ctx context.Context, productId int64) ([]*biz.Product, error) {
	repo.log.Infof("listProductByProductId start productId:%s", productId)

	pl, err := repo.data.db.Product.Query().
		Where(product.ProductId(productId)).
		All(ctx)
	if err != nil {
		repo.log.Errorf("listProductByProductId falied err:%s", err.Error())
		return nil, err
	}

	pdl := make([]*biz.Product, len(pl))
	for _, p := range pl {
		productPo := &biz.Product{}
		productPo.FromProductPo(p)
		pdl = append(pdl, productPo)
	}

	repo.log.Info("listProductByProductId success")
	return pdl, nil
}
