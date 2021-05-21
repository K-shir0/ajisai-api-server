package infrastructure

import (
	"context"
	"fmt"
	"github.com/K-shir0/ajisai-api-server/config"
	"github.com/K-shir0/ajisai-api-server/domain"
	"github.com/labstack/echo"
	"go.mongodb.org/mongo-driver/bson"
	"net/http"
	"time"
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

	// Context
	ctx := context.Background()

	// ルーティングa
	r.e.GET("/weathers", func(c echo.Context) error {

		collection := db.Database("test").Collection("hoge")

		cur, err := collection.Find(ctx, bson.M{})
		if err != nil {
			return err
		}

		var weathers []bson.M
		if err = cur.All(ctx, &weathers); err != nil {
			return err
		}

		return c.JSON(http.StatusOK, weathers)
	})

	r.e.POST("/weathers", func(c echo.Context) error {
		param := new(domain.Weather)
		if err := c.Bind(param); err != nil {
			return err
		}

		// set date
		now := time.Now()
		param.UpdatedAt, param.CreatedAt = now, now

		// insert DB
		res, err := db.Database("test").Collection("hoge").InsertOne(ctx, param)
		if err != nil {
			return err
		}

		return c.JSON(http.StatusOK, res)
	})

	// echo server start
	r.e.Logger.Fatal(r.e.Start(":" + r.config.Port))
}
