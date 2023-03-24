package main

import (
	"context"
	"net/url"
	"os"
	"time"

	"github.com/pepabo/rate-limit-middleware/pkg/limit"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/redis/go-redis/v9"
)

func main() {
	e := echo.New()

	e.Use(middleware.TimeoutWithConfig(middleware.TimeoutConfig{
		ErrorMessage: "Time out",
		Timeout:      30 * time.Second,
	}))

	redisClient := redis.NewClient(&redis.Options{
		// ex) "localhost:8080"
		Addr: os.Getenv("REDIS_ADDR"),
	})
	if _, err := redisClient.Ping(context.Background()).Result(); err != nil {
		e.Logger.Fatal(err)
	}

	rateLimitTime := 2
	e.Use(limit.RateLimitMiddleware(redisClient, rateLimitTime))

	proxyTargetUrl, err := url.Parse(os.Getenv("PROXY_TARGET_URL"))
	if err != nil {
		e.Logger.Fatal(err)
	}
	proxyTarget := []*middleware.ProxyTarget{
		{
			URL: proxyTargetUrl,
		},
	}
	e.Use(middleware.Proxy(middleware.NewRoundRobinBalancer(proxyTarget)))
	e.Logger.Fatal(e.Start("0.0.0.0:8088"))
}
