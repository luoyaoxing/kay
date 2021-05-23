package server

import (
	"github.com/go-kratos/kratos/v2/registry"
	"github.com/google/wire"

	consul "github.com/go-kratos/consul/registry"
	consulAPI "github.com/hashicorp/consul/api"
)

var ProviderSet = wire.NewSet(NewGrpcServer, NewRegistrar)

func NewRegistrar() registry.Registrar {
	config := consulAPI.DefaultConfig()
	config.Address = "106.54.179.120:8500"

	cli, err := consulAPI.NewClient(config)
	if err != nil {
		panic(err)
	}
	r := consul.New(cli)
	return r
}
