package data

import "kay/app/kv/service/internal/data/etcd"

type Data struct {
	*etcd.EtcdImpl
}

func NewData(opts ...etcd.Option) (*Data, func(), error) {
	impl, err := etcd.NewEtcdImpl(opts...)
	if err != nil {
		return nil, nil, err
	}

	data := &Data{
		impl,
	}

	return data, func() {
		impl.Close()
	}, nil
}
