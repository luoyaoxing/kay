package biz

import (
	pb "kay/api/deal/v1"
	"kay/app/deal/service/internal/data/ent"
)

type Deal struct {
	DealId      int64
	SkuId       int64
	ProductId   int64
	ProductName string
	Uid         int64
	CustomName  string
	Phone       string
	Price       int32
}

func (d *Deal) FromDealPO(dealPo *ent.Deal) {
	d.DealId = dealPo.ID
	d.SkuId = dealPo.SkuId
	d.ProductId = dealPo.ProductId
	d.ProductName = dealPo.ProductName
	d.Uid = dealPo.UID
	d.CustomName = dealPo.CustomName
	d.Phone = dealPo.Phone
	d.Price = int32(dealPo.Price)
	return
}

func (d *Deal) FromDealDTO(dealDTO *pb.DealDTO) {
	d.DealId = dealDTO.DealId
	d.SkuId = dealDTO.SkuId
	d.ProductId = dealDTO.ProductId
	d.ProductName = dealDTO.ProductName
	d.Uid = dealDTO.Uid
	d.CustomName = dealDTO.CustomName
	d.Phone = dealDTO.Phone
	d.Price = dealDTO.Price
	return
}

func (d *Deal) ToDealDTO(dealDTO *pb.DealDTO) {
	dealDTO.DealId = d.DealId
	dealDTO.SkuId = d.SkuId
	dealDTO.ProductId = d.ProductId
	dealDTO.ProductName = d.ProductName
	dealDTO.Uid = d.Uid
	dealDTO.CustomName = d.CustomName
	dealDTO.Phone = d.Phone
	dealDTO.Price = d.Price
	return
}
