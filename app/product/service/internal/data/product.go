package data

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"kay/app/product/service/internal/biz"
	"kay/app/product/service/internal/data/ent/product"
	"kay/pkg/utils/dec"
)

type ProductRepo struct {
	data *Data
	log  *log.Helper
}

func NewProductRepo(data *Data, logger log.Logger) *ProductRepo {
	return &ProductRepo{data: data, log: log.NewHelper("product/data", logger)}
}

func (pr *ProductRepo) CreateProduct(ctx context.Context, p *biz.Product) (int64, error) {
	pr.log.Infof("create product start p:%s", p)

	product, err := pr.data.db.Product.Create().
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
		pr.log.Errorf("save failed err:%s", err.Error())
		return 0, err
	}

	pr.log.Debugf("save product:%s", dec.JsonEncode(product))
	skuId := product.ID
	pr.log.Infof("create product success skuId:%d", skuId)
	return skuId, err
}

func (pr *ProductRepo) GetProductById(ctx context.Context, skuId int64) (*biz.Product, error) {
	pr.log.Infof("GetProductById start skuId:%s", skuId)

	p, err := pr.data.db.Product.Get(ctx, skuId)
	if err != nil {
		pr.log.Errorf("get failed skuId:%d err:%s", skuId, err.Error())
		return nil, err
	}

	productDo := &biz.Product{}
	productDo.FromProductPo(p)

	pr.log.Info("GetProductById success")
	return productDo, err
}

func (pr *ProductRepo) ListProductByProductId(ctx context.Context, productId int64) ([]*biz.Product, error) {
	pr.log.Infof("listProductByProductId start productId:%s", productId)

	pl, err := pr.data.db.Product.Query().
		Where(product.ProductId(productId)).
		All(ctx)
	if err != nil {
		pr.log.Errorf("listProductByProductId falied err:%s", err.Error())
		return nil, err
	}

	pdl := make([]*biz.Product, len(pl))
	for _, p := range pl {
		productPo := &biz.Product{}
		productPo.FromProductPo(p)
		pdl = append(pdl, productPo)
	}

	pr.log.Info("listProductByProductId success")
	return pdl, nil
}
