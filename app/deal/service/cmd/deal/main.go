package main

import (
	"flag"
	"github.com/go-kratos/kratos/v2/registry"
	"kay/app/deal/service/internal/biz"
	"kay/app/deal/service/internal/conf"
	"kay/app/deal/service/internal/data"
	"kay/app/deal/service/internal/server"
	"kay/app/deal/service/internal/service"
	"os"

	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/config"
	"github.com/go-kratos/kratos/v2/config/file"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/transport/grpc"
	"gopkg.in/yaml.v2"
)

// go build -ldflags "-X main.Version=x.y.z"
var (
	// flagconf is the config flag.
	flagconf string
)

func init() {
	flag.StringVar(&flagconf, "conf", "D:/dev/go/kay/app/deal/service/configs", "config path, eg: -conf config.yaml")
}

func newApp(logger log.Logger, gs *grpc.Server, rr registry.Registrar) *kratos.App {
	return kratos.New(
		kratos.Name("kay.deal.service"),
		kratos.Version("v1"),
		kratos.Metadata(map[string]string{}),
		kratos.Logger(logger),
		kratos.Registrar(rr),
		kratos.Server(
			gs,
		),
	)
}

func main() {
	flag.Parse()
	logger := log.NewStdLogger(os.Stdout)

	c := config.New(
		config.WithSource(
			file.NewSource(flagconf),
		),
		config.WithDecoder(func(kv *config.KeyValue, v map[string]interface{}) error {
			return yaml.Unmarshal(kv.Value, v)
		}),
	)
	if err := c.Load(); err != nil {
		panic(err)
	}

	var bc conf.Bootstrap
	if err := c.Scan(&bc); err != nil {
		panic(err)
	}

	app, cleanup, err := initApp(bc.Server, bc.Data, logger)
	if err != nil {
		panic(err)
	}
	defer cleanup()

	// start and wait for stop signal
	if err := app.Run(); err != nil {
		panic(err)
	}
}

func initApp(confServer *conf.Server, confData *conf.Data, logger log.Logger) (*kratos.App, func(), error) {
	dataData, cleanup, err := data.NewData(confData, logger)
	if err != nil {
		return nil, nil, err
	}
	dealRepo := data.NewDealRepo(dataData, logger)
	dealUseCase := biz.NewDealUseCase(dealRepo, logger)
	productService := service.NewDealService(dealUseCase, logger)
	grpcServer := server.NewGRPCServer(confServer, productService, logger)
	registrar := server.NewRegistrar()
	app := newApp(logger, grpcServer, registrar)
	return app, func() {
		cleanup()
	}, nil
}
