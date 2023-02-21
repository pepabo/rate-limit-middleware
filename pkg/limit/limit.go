package limit

import (
	"context"
	"log"
	"time"

	"github.com/go-redis/redis_rate/v10"
	"github.com/labstack/echo/v4"
	"github.com/redis/go-redis/v9"
)

const redisListName = "allList"

func RateLimitMiddleware(redisClient *redis.Client, limitPerSec int) echo.MiddlewareFunc {
	limiter := redis_rate.NewLimiter(redisClient)
	chanList := []chan any{}
	go func() {
		for {
			if len(chanList) != 0 {
				result, err := limiter.Allow(context.Background(), redisListName, redis_rate.PerSecond(limitPerSec))
				if err != nil {
					// chanListの中身にすべてシグナルを送る
					log.Println(err)
					for _, ch := range chanList {
						ch <- true
					}
					continue
				}
				if result.Allowed == 1 {
					ch := chanList[0]
					chanList = chanList[1:]
					// タイムアウトが起きてもcloseしないのでpanicが起こらない
					ch <- true
				}
			}
			time.Sleep(10 * time.Millisecond)
		}
	}()
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		ch := make(chan any)
		chanList = append(chanList, ch)
		return func(c echo.Context) error {
			// チャネルの応答を待つ
			_ = <-ch
			return next(c)
		}
	}
}
