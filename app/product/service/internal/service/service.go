package service

import (
	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
	pb "kay/api/product/v1"
	"kay/app/product/service/internal/biz"
	"kay/app/product/service/internal/data"
)

var ProviderSet = wire.NewSet(NewProductService)

type ProductService struct {
	pb.UnimplementedProductServer

	productUc   *biz.ProductUseCase
	itemStockUc *biz.ItemStockUseCase

	data *data.Data

	log *log.Helper
}

func NewProductService(productUc *biz.ProductUseCase, itemStockUc *biz.ItemStockUseCase, data *data.Data, logger log.Logger) pb.ProductServer {
	return &ProductService{
		productUc:   productUc,
		itemStockUc: itemStockUc,

		data: data,

		log: log.NewHelper("product/service", logger),
	}
}
