package etcd

import (
	"github.com/go-kratos/kratos/v2/log"
	"os"
	"time"
)

var (
	DefaultTimeout  = 5 * time.Second
	DefaultEndPoint = "106.54.179.120:2379"
	DefaultLogger   = log.NewStdLogger(os.Stdout)
)

type Options struct {
	timeout   time.Duration
	endPoints []string

	logger log.Logger
}

type Option interface {
	Apply(o *Options)
}

type funcOption struct {
	f func(o *Options)
}

func newFuncOption(f func(o *Options)) Option {
	return &funcOption{
		f: f,
	}
}

func (opt *funcOption) Apply(o *Options) {
	opt.f(o)
}

func WithTimeout(timeout time.Duration) Option {
	return newFuncOption(func(o *Options) {
		o.timeout = timeout
	})
}

func WithEndpoints(endpoints ...string) Option {
	return newFuncOption(func(o *Options) {
		o.endPoints = append(o.endPoints, endpoints...)
	})
}

func WithLogger(logger log.Logger) Option {
	return newFuncOption(func(o *Options) {
		o.logger = logger
	})
}
