package biz

import (
	v1 "kay/api/product/v1"
	"kay/app/product/service/internal/data/ent"
)

type Product struct {
	SkuId       int64
	VersionId   int64
	VersionName string
	ProductId   int64
	ProductName string
	Price       int
	Attr        string
	ProductDesc string
}

func (p *Product) FromProductPo(productPo *ent.Product) {
	p.SkuId = productPo.ID
	p.VersionId = productPo.VersionId
	p.VersionName = productPo.VersionName
	p.ProductId = productPo.ProductId
	p.ProductName = productPo.ProductName
	p.Price = productPo.Price
	p.ProductDesc = productPo.ProductDesc
	p.Attr = productPo.Attr
	return
}

func (p *Product) FromProductDto(productDto *v1.ProductDTO) {
	p.SkuId = productDto.GetSkuId()
	p.VersionId = productDto.GetVersionId()
	p.VersionName = productDto.GetVersionName()
	p.ProductId = productDto.GetProductId()
	p.ProductName = productDto.GetProductName()
	p.Price = int(productDto.GetPrice())
	p.ProductDesc = productDto.GetProductDesc()
	p.Attr = productDto.GetAttr()
	return
}

func (p *Product) ToProductDto(productDto *v1.ProductDTO) {
	productDto.SkuId = p.SkuId
	productDto.VersionId = p.VersionId
	productDto.VersionName = p.VersionName
	productDto.ProductId = p.ProductId
	productDto.ProductName = p.ProductName
	productDto.Price = int32(p.Price)
	productDto.ProductDesc = p.ProductDesc
	productDto.Attr = p.Attr
	return
}
