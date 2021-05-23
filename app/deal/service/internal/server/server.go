package server

import (
	consul "github.com/go-kratos/consul/registry"
	"github.com/go-kratos/kratos/v2/registry"
	"github.com/google/wire"
	consulAPI "github.com/hashicorp/consul/api"
)

// ProviderSet is server providers.
var ProviderSet = wire.NewSet(NewGRPCServer, NewRegistrar)

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
