package etcd

import (
	"context"
	"github.com/coreos/etcd/clientv3"
	"github.com/go-kratos/kratos/v2/log"
	"sync"
)

type EtcdImpl struct {
	opts *Options

	m     sync.RWMutex
	close bool

	client *clientv3.Client
	log    *log.Helper
}

func NewEtcdImpl(opts ...Option) (*EtcdImpl, error) {
	defaultOpts := &Options{
		timeout:   DefaultTimeout,
		endPoints: []string{DefaultEndPoint},
		logger:    DefaultLogger,
	}

	for _, opt := range opts {
		opt.Apply(defaultOpts)
	}

	cli, err := clientv3.New(clientv3.Config{
		Endpoints:   defaultOpts.endPoints,
		DialTimeout: defaultOpts.timeout,
	})

	if err != nil {
		return nil, err
	}

	impl := &EtcdImpl{
		opts: defaultOpts,

		client: cli,
		log:    log.NewHelper("data/etcd", defaultOpts.logger),
	}

	return impl, nil
}

func (impl *EtcdImpl) Put(ctx context.Context, k, v string) error {
	impl.m.RLock()
	defer impl.m.RUnlock()

	if impl.close {
		return ERR_ETCD_CLOSE
	}

	if k == "" {
		return ERR_ETCD_EMPTY
	}

	_, err := impl.client.Put(ctx, k, v)
	if err != nil {
		impl.log.Errorf("put failed [k]:%s [v]:%s err:%s", k, v, err.Error())
		return ERR_ETCD_PUT
	}
	return nil
}

func (impl *EtcdImpl) Get(ctx context.Context, k string) ([]byte, error) {
	impl.m.RLock()
	defer impl.m.RUnlock()

	if impl.close {
		return nil, ERR_ETCD_CLOSE
	}

	if k == "" {
		return nil, ERR_ETCD_EMPTY
	}

	resp, err := impl.client.Get(ctx, k)
	if err != nil {
		impl.log.Errorf("get failed [k]:%s err:%s", k, err.Error())
		return nil, ERR_ETCD_GET
	}

	return resp.Kvs[resp.Count-1].Value, nil
}

func (impl *EtcdImpl) Close() error {
	impl.m.Lock()
	defer impl.m.Unlock()

	impl.close = true
	err := impl.client.Close()
	if err != nil {
		impl.log.Errorf("close failed err:%s", err.Error())
		return err
	}
	return nil
}
