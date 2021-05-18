package infrastructure

import (
	"context"
	"fmt"
	"github.com/K-shir0/ajisai-api-server/config"
	"github.com/labstack/echo"
	"net/http"
)

type Router struct {
	config *config.Config
	e      *echo.Echo
}

func Init() {
	var ctx = context.Background()

	r := Router{
		config: config.New(),
		e:      echo.New(),
	}

	// コントローラの呼び出し
	rdb := r.NewRedisHandler()

	err := rdb.Set(ctx, "key", "value", 0).Err()
	if err != nil {
		panic(err)
	}

	val, err := rdb.Get(ctx, "key").Result()
	if err != nil {
		panic(err)
	}
	fmt.Println("key", val)

	//val2, err := rdb.Get(ctx, "key2").Result()
	//if err == redis.Nil {
	//	fmt.Println("key2 does not exist")
	//} else if err != nil {
	//	panic(err)
	//} else {
	//	fmt.Println("key2", val2)
	//}

	// ルーティング
	r.e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})

	// echo server start
	r.e.Logger.Fatal(r.e.Start(":" + r.config.Port))
}
