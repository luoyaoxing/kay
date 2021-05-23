package biz

import "kay/app/product/service/internal/data/ent"

type ItemStock struct {
	SkuId        int64
	TotalStock   int64
	ConsumeStock int64
	LeftStock    int64
}

func (s *ItemStock) FromItemStockPo(item *ent.Item) {
	s.SkuId = item.ID
	s.TotalStock = int64(item.TotalStock)
	s.LeftStock = int64(item.LeftStock)
	s.ConsumeStock = int64(item.ConsumeStock)
	return
}
