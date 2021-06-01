package service

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"kay/app/kv/service/internal/biz"

	pb "kay/api/kv/v1"
)

type KvstoreService struct {
	pb.UnimplementedKvstoreServer

	etcdUseCase *biz.EtcdKVUseCase
	log         *log.Helper
}

func NewKvstoreService(etcdUseCase *biz.EtcdKVUseCase, logger log.Logger) pb.KvstoreServer {
	return &KvstoreService{etcdUseCase: etcdUseCase, log: log.NewHelper("etcd/service", logger)}
}

func (s *KvstoreService) PutKvstore(ctx context.Context, req *pb.PutKvstoreRequest) (*pb.PutKvstoreReply, error) {
	return &pb.PutKvstoreReply{}, nil
}
func (s *KvstoreService) GetKvstore(ctx context.Context, req *pb.GetKvstoreRequest) (*pb.GetKvstoreReply, error) {
	return &pb.GetKvstoreReply{}, nil
}
