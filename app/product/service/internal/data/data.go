package data

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"kay/app/product/service/internal/conf"
	"kay/app/product/service/internal/data/ent"
	"kay/pkg/utils/dec"
)

type Data struct {
	db *ent.Client
}

func NewData(conf conf.Data, logger log.Logger) (*Data, func(), error) {
	log := log.NewHelper("product/data", logger)
	log.Infof("NewData start conf:%s \n", dec.JsonEncode(conf))

	client, err := ent.Open(
		conf.Database.Driver,
		conf.Database.Source,
	)

	if err != nil {
		log.Errorf("open db failed err:%s \n", err.Error())
		return nil, nil, err
	}

	err = client.Schema.Create(context.Background())
	if err != nil {
		log.Errorf("create schema failed err:%s \n", err.Error())
		return nil, nil, err
	}

	data := &Data{
		db: client,
	}

	return data, func() {
		if err = client.Close(); err != nil {
			log.Errorf("close db err:%s\n", err.Error())
		}
	}, nil
}
