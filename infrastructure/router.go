package infrastructure

import (
	"ajisai-api-server/config"
	"github.com/labstack/echo"
	"net/http"
)

func Init() {
	conf := config.New()

	e := echo.New()

	// コントローラの呼び出し

	// ルーティング
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})

	// echo server start
	e.Logger.Fatal(e.Start(":" + conf.Port))
}
