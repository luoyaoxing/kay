package main

import (
	"flag"
	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/config"
	"github.com/go-kratos/kratos/v2/config/file"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/registry"
	"github.com/go-kratos/kratos/v2/transport/grpc"
	"gopkg.in/yaml.v2"
	"kay/app/product/service/internal/biz"
	"kay/app/product/service/internal/conf"
	"kay/app/product/service/internal/data"
	"kay/app/product/service/internal/server"
	"kay/app/product/service/internal/service"
	"os"
)

// go build -ldflags "-X main.Version=x.y.z"
var (
	// flagconf is the config flag.
	flagconf string
)

func init() {
	flag.StringVar(&flagconf, "conf", "D:/dev/go/kay/app/product/service/configs", "config path, eg: -conf config.yaml")
}

func newApp(logger log.Logger, gs *grpc.Server, rr registry.Registrar) *kratos.App {
	return kratos.New(
		kratos.Name("kay.product.service"),
		kratos.Version("v1"),
		kratos.Metadata(map[string]string{}),
		kratos.Logger(logger),
		kratos.Server(
			gs,
		),
		kratos.Registrar(rr),
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
	itemStockRepo := data.NewItemStockRepo(dataData, logger)
	productRepo := data.NewProductRepo(dataData, logger)
	itemStockUseCase := biz.NewItemStockUseCase(itemStockRepo, logger)
	productUseCase := biz.NewProductUseCase(productRepo, logger)
	productService := service.NewProductService(productUseCase, itemStockUseCase, dataData, logger)
	grpcServer := server.NewGrpcServer(confServer, productService)
	registrar := server.NewRegistrar()
	app := newApp(logger, grpcServer, registrar)
	return app, func() {
		cleanup()
	}, nil
}
