package infrastructure

import (
	"github.com/K-shir0/ajisai-api-server/config"
	"github.com/K-shir0/ajisai-api-server/domain"
	"github.com/K-shir0/ajisai-api-server/interfaces/database"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
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

	r.e.Use(middleware.CORS())

	// sqlコントローラの呼び出し
	db := r.NewSqlHandler()

	wr := database.WeathersRepository{Collection: db.Collection("weather")}

	// ルーティング
	r.e.GET("/weathers", func(c echo.Context) error {

		weathers, err := wr.FindAll()
		if err != nil {
			return c.String(http.StatusBadRequest, "")
		}

		return c.JSON(http.StatusOK, weathers)
	})

	r.e.POST("/weathers", func(c echo.Context) error {
		param := new(domain.Weather)
		if err := c.Bind(param); err != nil {
			return err
		}

		// insert DB
		res, err := wr.Store(param)
		if err != nil {
			return err
		}

		return c.JSON(http.StatusOK, res)
	})

	// echo server start
	r.e.Logger.Fatal(r.e.Start(":" + r.config.Port))
}
