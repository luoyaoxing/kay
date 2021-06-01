package data

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
)

type kvStoreRepo struct {
	data *Data

	log *log.Helper
}

func NewKVStoreRepo(data *Data, logger log.Logger) *kvStoreRepo {
	return &kvStoreRepo{data: data, log: log.NewHelper("etcd/data", logger)}
}

func (repo *kvStoreRepo) PutKV(ctx context.Context, k, v string) error {
	return repo.data.Put(ctx, k, v)
}

func (repo *kvStoreRepo) GetV(ctx context.Context, k string) ([]byte, error) {
	return repo.data.Get(ctx, k)
}
