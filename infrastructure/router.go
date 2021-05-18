package infrastructure

import (
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
	r := Router{
		config: config.New(),
		e:      echo.New(),
	}

	// コントローラの呼び出し
	db := r.NewSqlHandler()

	// Ping the primary
	fmt.Println(db)

	// ルーティング
	r.e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})

	// echo server start
	r.e.Logger.Fatal(r.e.Start(":" + r.config.Port))
}
