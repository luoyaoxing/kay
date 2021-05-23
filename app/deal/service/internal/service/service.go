package service

import (
	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
	pb "kay/api/deal/v1"
	"kay/app/deal/service/internal/biz"
)

// ProviderSet is service providers.
var ProviderSet = wire.NewSet(NewDealService)

type DealService struct {
	pb.UnimplementedDealServer

	uc *biz.DealUseCase

	log *log.Helper
}

func NewDealService(uc *biz.DealUseCase, logger log.Logger) pb.DealServer {
	return &DealService{uc: uc, log: log.NewHelper("server/service", logger)}
}
