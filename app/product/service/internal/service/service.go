package service

import (
	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
	pb "kay/api/product/v1"
	"kay/app/product/service/internal/biz"
)

var ProviderSet = wire.NewSet(NewProductService)

type ProductService struct {
	pb.UnimplementedProductServer

	productUc   *biz.ProductUseCase
	itemStockUc *biz.ItemStockUseCase

	log *log.Helper
}

func NewProductService(productUc *biz.ProductUseCase, itemStockUc *biz.ItemStockUseCase, logger log.Logger) pb.ProductServer {
	return &ProductService{
		productUc:   productUc,
		itemStockUc: itemStockUc,
		log:         log.NewHelper("product/service", logger),
	}
}
