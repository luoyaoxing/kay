package biz

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
)

type KvStoreRepo interface {
	PutKV(ctx context.Context, k, v string) error
	GetV(ctx context.Context, k string) ([]byte, error)
}

type EtcdKVUseCase struct {
	repo KvStoreRepo
	log  *log.Helper
}

func NewEtcdKVUseCase(repo KvStoreRepo, logger log.Logger) *EtcdKVUseCase {
	return &EtcdKVUseCase{repo: repo, log: log.NewHelper("etcd/biz", logger)}
}

func (repo *EtcdKVUseCase) SetKV(ctx context.Context, k, v string) error {
	return repo.repo.PutKV(ctx, k, v)
}

func (repo *EtcdKVUseCase) GetV(ctx context.Context, k string) ([]byte, error) {
	return repo.repo.GetV(ctx, k)
}
