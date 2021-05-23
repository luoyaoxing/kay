package service

import (
	"context"
	"kay/app/deal/service/internal/biz"
	"kay/pkg/utils/dec"

	pb "kay/api/deal/v1"
)

func (s *DealService) CreateDeal(ctx context.Context, req *pb.CreateDealRequest) (*pb.CreateDealReply, error) {
	s.log.Infof("CreateDeal start req:%s", dec.JsonEncode(req))

	reply := &pb.CreateDealReply{}
	replyHeader := &pb.ResponseHeader{}
	replyHeader.Errno = 0
	replyHeader.Errmsg = "success"
	reply.Header = replyHeader

	deal := &biz.Deal{}
	deal.FromDealDTO(req.Body.Deal)

	err := s.uc.CreateDeal(ctx, deal)
	if err != nil {
		replyHeader.Errno = 10001
		replyHeader.Errmsg = "创建订单失败"
		s.log.Errorf("CreateDeal failed err:%s", err.Error())
		return reply, nil
	}

	s.log.Infof("CreateDeal success")
	return reply, nil
}

func (s *DealService) GetDeal(ctx context.Context, req *pb.GetDealRequest) (*pb.GetDealReply, error) {
	s.log.Infof("GetDeal start req:%s", dec.JsonEncode(req))

	reply := &pb.GetDealReply{}
	replyHeader := &pb.ResponseHeader{}
	replyHeader.Errno = 0
	replyHeader.Errmsg = "success"
	reply.Header = replyHeader

	deal, err := s.uc.GetDealInfo(ctx, req.Body.DealId)
	if err != nil {
		replyHeader.Errno = 10001
		replyHeader.Errmsg = "创建订单失败"
		s.log.Errorf("CreateDeal failed err:%s", err.Error())
		return reply, nil
	}

	dealDTO := &pb.DealDTO{}
	deal.ToDealDTO(dealDTO)

	body := &pb.GetDealReply_Body{}
	body.Deal = dealDTO
	reply.Body = body

	s.log.Infof("GetDeal success")
	return reply, nil
}
