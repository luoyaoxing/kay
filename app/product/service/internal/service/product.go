package service

import (
	"context"
	"kay/app/product/service/internal/biz"
	"kay/app/product/service/internal/common"
	"kay/app/product/service/internal/data/ent"
	"kay/pkg/utils/dec"

	pb "kay/api/product/v1"
)

func (s *ProductService) CreateProduct(ctx context.Context, req *pb.CreateProductRequest) (*pb.CreateProductReply, error) {
	s.log.Infof("create product start req:%s", dec.JsonEncode(req))

	reply := &pb.CreateProductReply{}

	replyHeader := &pb.ResponseHeader{}
	replyHeader.Errno = 0
	replyHeader.Errmsg = "success"
	reply.Header = replyHeader

	tx, err := s.data.GetDb().Tx(ctx)
	tx.OnCommit(func(committer ent.Committer) ent.Committer {
		s.log.Infof("tx on commit")
		return committer
	})
	tx.OnRollback(func(rollbacker ent.Rollbacker) ent.Rollbacker {
		s.log.Infof("tx on rollback")
		return rollbacker
	})

	if err != nil {
		replyHeader.Errno = 10003
		replyHeader.Errmsg = "开启事务失败"
		s.log.Errorf("Tx failed err:%s", err.Error())
		return reply, nil
	}

	ctx = context.WithValue(ctx, common.TxKey, tx)

	product := &biz.Product{}
	product.FromProductDto(req.GetBody().GetProduct())

	skuId, err := s.productUc.CreateProductWithTx(ctx, product)
	if err != nil {
		replyHeader.Errno = 10001
		replyHeader.Errmsg = "创建产品失败"
		s.log.Errorf("create product failed err:%s", err.Error())
		err = tx.Rollback()
		if err != nil {
			s.log.Errorf("Rollback failed err:%s", err.Error())
		}
		return reply, nil
	}

	stock := req.Body.GetStock()

	itemStock := &biz.ItemStock{
		SkuId:      skuId,
		TotalStock: stock,
	}
	err = s.itemStockUc.CreateStockWithTx(ctx, itemStock)
	if err != nil {
		replyHeader.Errno = 10001
		replyHeader.Errmsg = "创建产品库存失败"
		s.log.Errorf("CreateStock failed err:%s", err.Error())
		err = tx.Rollback()
		if err != nil {
			s.log.Errorf("Rollback failed err:%s", err.Error())
		}
		return reply, nil
	}

	err = tx.Commit()
	if err != nil {
		replyHeader.Errno = 10004
		replyHeader.Errmsg = "事务提交失败"
		s.log.Errorf("Commit failed err:%s", err.Error())
		return reply, nil
	}

	replyBody := &pb.CreateProductReply_Body{}
	replyBody.SkuId = skuId

	reply.Body = replyBody
	return reply, nil
}

func (s *ProductService) GetProductInfo(ctx context.Context, req *pb.GetProductInfoRequest) (*pb.GetProductInfoReply, error) {
	s.log.Infof("get productInfo req:%s", dec.JsonEncode(req))

	reply := &pb.GetProductInfoReply{}

	replyHeader := &pb.ResponseHeader{}
	replyHeader.Errno = 0
	replyHeader.Errmsg = "success"

	reply.Header = replyHeader
	product, err := s.productUc.GetProductInfo(ctx, req.Body.SkuId)
	if err != nil {
		replyHeader.Errno = 10001
		replyHeader.Errmsg = "查询产品失败"
		s.log.Errorf("getProductInfo falied err:%s", err.Error())
		return reply, nil
	}

	productDto := &pb.ProductDTO{}
	product.ToProductDto(productDto)

	replyBody := &pb.GetProductInfoReply_Body{}
	replyBody.Product = productDto

	reply.Body = replyBody
	return reply, nil
}

func (s *ProductService) ListProductByProductId(ctx context.Context, req *pb.ListProductByProductIdRequest) (*pb.ListProductByProductIdReply, error) {
	s.log.Infof("listProductByProductId start req:%s", dec.JsonEncode(req))

	reply := &pb.ListProductByProductIdReply{}

	replyHeader := &pb.ResponseHeader{}
	replyHeader.Errno = 0
	replyHeader.Errmsg = "success"

	reply.Header = replyHeader
	productList, err := s.productUc.ListProductByProductId(ctx, req.Body.ProductId)
	if err != nil {
		replyHeader.Errno = 10001
		replyHeader.Errmsg = "查询产品失败"
		s.log.Errorf("getProductInfo falied err:%s", err.Error())
		return reply, nil
	}

	var productDtoList []*pb.ProductDTO
	for _, product := range productList {
		productDto := &pb.ProductDTO{}
		product.ToProductDto(productDto)
		productDtoList = append(productDtoList, productDto)
	}

	replyBody := &pb.ListProductByProductIdReply_Body{}
	replyBody.ProductList = productDtoList

	reply.Body = replyBody
	return reply, nil
}

func (s *ProductService) DeStock(ctx context.Context, req *pb.DeStockRequest) (*pb.DeStockReply, error) {
	s.log.Infof("DeStock start req:%s", dec.JsonEncode(req))

	reply := &pb.DeStockReply{}

	replyHeader := &pb.ResponseHeader{}
	replyHeader.Errno = 0
	replyHeader.Errmsg = "success"
	reply.Header = replyHeader

	skuId := req.Body.GetSkuId()
	p, err := s.productUc.GetProductInfo(ctx, skuId)
	if err != nil {
		replyHeader.Errno = 10001
		replyHeader.Errmsg = "查询产品失败"
		s.log.Errorf("getProductInfo falied err:%s", err.Error())
		return reply, nil
	}

	_, err = s.itemStockUc.DoDeStock(ctx, p)
	if err != nil {
		replyHeader.Errno = 10002
		replyHeader.Errmsg = "扣减库存失败"
		s.log.Errorf("getProductInfo falied err:%s", err.Error())
		return reply, nil
	}
	return reply, nil
}
