package infrastructure

import (
	"fmt"
	"github.com/K-shir0/ajisai-api-server/config"
	"github.com/K-shir0/ajisai-api-server/domain"
	"github.com/K-shir0/ajisai-api-server/interfaces/database"
	"github.com/gorilla/websocket"
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

	// コントローラの呼び出し
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

		message = param
		flag = true

		return c.JSON(http.StatusOK, res)
	})

	r.e.GET("/ws", hello)

	// echo server start
	r.e.Logger.Fatal(r.e.Start(":" + r.config.Port))
}

var message *domain.Weather
var flag = false

var (
	upgrader = websocket.Upgrader{
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
	}
)

func hello(c echo.Context) error {
	ws, err := upgrader.Upgrade(c.Response(), c.Request(), nil)
	if err != nil {
		fmt.Println(err)
		return err
	}
	defer ws.Close()

	for {
		// Write
		err := ws.WriteMessage(websocket.TextMessage, []byte("Hello, Client!"))
		if err != nil {
			c.Logger().Error(err)
		}

		// Read
		_, msg, err := ws.ReadMessage()
		if err != nil {
			c.Logger().Error(err)
		}
		fmt.Printf("%s\n", msg)
	}

	//websocket.Handler(func(ws *websocket.Conn) {
	//	defer func(ws *websocket.Conn) {
	//		err := ws.Close()
	//		if err != nil {
	//			c.Logger().Error(err)
	//		}
	//	}(ws)
	//	for {
	//		// Write
	//		if flag {
	//			flag = false
	//
	//			err := websocket.JSON.Send(ws, message)
	//			if err != nil {
	//				c.Logger().Error(err)
	//				break
	//			}
	//		}
	//	}
	//}).ServeHTTP(c.Response(), c.Request())
	//return nil
}
