package ratelimit

import (
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware"
	"go.uber.org/ratelimit"
	"time"
)

func Funnel() middleware.Middleware {
	rl := ratelimit.New(1, ratelimit.Per(1*time.Second))
	return func(handler middleware.Handler) middleware.Handler {
		logger := log.DefaultLogger
		now := rl.Take() // 基于漏斗的限流
		logger.Log("funnel take time now:%d", now.Unix())

		return handler
	}
}
